/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tabletmanager

// This file handles the agent state changes.

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"vitess.io/vitess/go/event"
	"vitess.io/vitess/go/trace"
	"vitess.io/vitess/go/vt/key"
	"vitess.io/vitess/go/vt/log"
	"vitess.io/vitess/go/vt/logutil"
	"vitess.io/vitess/go/vt/mysqlctl"
	querypb "vitess.io/vitess/go/vt/proto/query"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	"vitess.io/vitess/go/vt/topo"
	"vitess.io/vitess/go/vt/topo/topoproto"
	"vitess.io/vitess/go/vt/topotools"
	"vitess.io/vitess/go/vt/vttablet/tabletmanager/events"
	"vitess.io/vitess/go/vt/vttablet/tabletmanager/vreplication"
	"vitess.io/vitess/go/vt/vttablet/tabletserver/rules"
)

var (
	// constants for this module
	historyLength = 16

	// gracePeriod is the amount of time we pause after broadcasting to vtgate
	// that we're going to stop serving a particular target type (e.g. when going
	// spare, or when being promoted to master). During this period, we expect
	// vtgate to gracefully redirect traffic elsewhere, before we begin actually
	// rejecting queries for that target type.
	gracePeriod = flag.Duration("serving_state_grace_period", 0, "how long to pause after broadcasting health to vtgate, before enforcing a new serving state")

	publishRetryInterval = flag.Duration("publish_retry_interval", 30*time.Second, "how long vttablet waits to retry publishing the tablet record")
)

// Query rules from blacklist
const blacklistQueryRules string = "BlacklistQueryRules"

// loadBlacklistRules loads and builds the blacklist query rules
func (agent *ActionAgent) loadBlacklistRules(ctx context.Context, tablet *topodatapb.Tablet, blacklistedTables []string) (err error) {
	blacklistRules := rules.New()
	if len(blacklistedTables) > 0 {
		// tables, first resolve wildcards
		tables, err := mysqlctl.ResolveTables(ctx, agent.MysqlDaemon, topoproto.TabletDbName(tablet), blacklistedTables)
		if err != nil {
			return err
		}

		// Verify that at least one table matches the wildcards, so
		// that we don't add a rule to blacklist all tables
		if len(tables) > 0 {
			log.Infof("Blacklisting tables %v", strings.Join(tables, ", "))
			qr := rules.NewQueryRule("enforce blacklisted tables", "blacklisted_table", rules.QRFailRetry)
			for _, t := range tables {
				qr.AddTableCond(t)
			}
			blacklistRules.Add(qr)
		}
	}

	loadRuleErr := agent.QueryServiceControl.SetQueryRules(blacklistQueryRules, blacklistRules)
	if loadRuleErr != nil {
		log.Warningf("Fail to load query rule set %s: %s", blacklistQueryRules, loadRuleErr)
	}
	return nil
}

// lameduck changes the QueryServiceControl state to lameduck,
// brodcasts the new health, then sleep for grace period, to give time
// to clients to get the new status.
func (agent *ActionAgent) lameduck(reason string) {
	log.Infof("Agent is entering lameduck, reason: %v", reason)
	agent.QueryServiceControl.EnterLameduck()
	agent.broadcastHealth()
	time.Sleep(*gracePeriod)
	log.Infof("Agent is leaving lameduck")
}

func (agent *ActionAgent) broadcastHealth() {
	// get the replication delays
	agent.mutex.Lock()
	replicationDelay := agent._replicationDelay
	healthError := agent._healthy
	healthyTime := agent._healthyTime
	agent.mutex.Unlock()

	// send it to our observers
	// FIXME(alainjobart,liguo) add CpuUsage
	stats := &querypb.RealtimeStats{
		SecondsBehindMaster: uint32(replicationDelay.Seconds()),
	}
	stats.SecondsBehindMasterFilteredReplication, stats.BinlogPlayersCount = vreplication.StatusSummary()
	stats.Qps = agent.QueryServiceControl.Stats().QPSRates.TotalRate()
	if healthError != nil {
		stats.HealthError = healthError.Error()
	} else {
		timeSinceLastCheck := time.Since(healthyTime)
		if timeSinceLastCheck > *healthCheckInterval*3 {
			stats.HealthError = fmt.Sprintf("last health check is too old: %s > %s", timeSinceLastCheck, *healthCheckInterval*3)
		}
	}
	var ts int64
	terTime := agent.masterTermStartTime()
	if !terTime.IsZero() {
		ts = terTime.Unix()
	}
	go agent.QueryServiceControl.BroadcastHealth(ts, stats, *healthCheckInterval*3)
}

// refreshTablet needs to be run after an action may have changed the current
// state of the tablet.
func (agent *ActionAgent) refreshTablet(ctx context.Context, reason string) error {
	agent.checkLock()
	log.Infof("Executing post-action state refresh: %v", reason)

	span, ctx := trace.NewSpan(ctx, "ActionAgent.refreshTablet")
	span.Annotate("reason", reason)
	defer span.Finish()

	// TODO(sougou): change this to specifically look for global topo changes.
	tablet := agent.Tablet()
	agent.changeCallback(ctx, tablet, tablet)
	log.Infof("Done with post-action state refresh")
	return nil
}

// updateState will use the provided tablet record as the new tablet state,
// the current tablet as a base, run changeCallback, and dispatch the event.
func (agent *ActionAgent) updateState(ctx context.Context, newTablet *topodatapb.Tablet, reason string) {
	oldTablet := agent.Tablet()
	if oldTablet == nil {
		oldTablet = &topodatapb.Tablet{}
	}
	log.Infof("Running tablet callback because: %v", reason)
	agent.changeCallback(ctx, oldTablet, newTablet)
	agent.setTablet(newTablet)
	agent.publishState(ctx)
	event.Dispatch(&events.StateChange{
		OldTablet: *oldTablet,
		NewTablet: *newTablet,
		Reason:    reason,
	})
}

// changeCallback is run after every action that might
// have changed something in the tablet record or in the topology.
func (agent *ActionAgent) changeCallback(ctx context.Context, oldTablet, newTablet *topodatapb.Tablet) {
	agent.checkLock()

	span, ctx := trace.NewSpan(ctx, "ActionAgent.changeCallback")
	defer span.Finish()

	allowQuery := topo.IsRunningQueryService(newTablet.Type)

	// Read the shard to get SourceShards / TabletControlMap if
	// we're going to use it.
	var shardInfo *topo.ShardInfo
	var err error
	// this is just for logging
	var disallowQueryReason string
	// this is actually used to set state
	var disallowQueryService string
	var blacklistedTables []string
	updateBlacklistedTables := true
	if allowQuery {
		shardInfo, err = agent.TopoServer.GetShard(ctx, newTablet.Keyspace, newTablet.Shard)
		if err != nil {
			log.Errorf("Cannot read shard for this tablet %v, might have inaccurate SourceShards and TabletControls: %v", newTablet.Alias, err)
			updateBlacklistedTables = false
		} else {
			if oldTablet.Type == topodatapb.TabletType_RESTORE {
				// always start as NON-SERVING after a restore because
				// healthcheck has not been initialized yet
				allowQuery = false
				// setting disallowQueryService permanently turns off query service
				// since we want it to be temporary (until tablet is healthy) we don't set it
				// disallowQueryReason is only used for logging
				disallowQueryReason = "after restore from backup"
			} else {
				if newTablet.Type == topodatapb.TabletType_MASTER {
					if len(shardInfo.SourceShards) > 0 {
						allowQuery = false
						disallowQueryReason = "master tablet with filtered replication on"
						disallowQueryService = disallowQueryReason
					}
				} else {
					replicationDelay, healthErr := agent.HealthReporter.Report(true, true)
					if healthErr != nil {
						allowQuery = false
						disallowQueryReason = "unable to get health"
					} else {
						agent.mutex.Lock()
						agent._replicationDelay = replicationDelay
						agent.mutex.Unlock()
						if agent._replicationDelay > *unhealthyThreshold {
							allowQuery = false
							disallowQueryReason = "replica tablet with unhealthy replication lag"
						}
					}
				}
			}
			srvKeyspace, err := agent.TopoServer.GetSrvKeyspace(ctx, newTablet.Alias.Cell, newTablet.Keyspace)
			if err != nil {
				log.Errorf("failed to get SrvKeyspace %v with: %v", newTablet.Keyspace, err)
			} else {

				for _, partition := range srvKeyspace.GetPartitions() {
					if partition.GetServedType() != newTablet.Type {
						continue
					}

					for _, tabletControl := range partition.GetShardTabletControls() {
						if key.KeyRangeEqual(tabletControl.GetKeyRange(), newTablet.GetKeyRange()) {
							if tabletControl.QueryServiceDisabled {
								allowQuery = false
								disallowQueryReason = "TabletControl.DisableQueryService set"
								disallowQueryService = disallowQueryReason
							}
							break
						}
					}
				}
			}
			if tc := shardInfo.GetTabletControl(newTablet.Type); tc != nil {
				if topo.InCellList(newTablet.Alias.Cell, tc.Cells) {

					blacklistedTables = tc.BlacklistedTables
				}
			}
		}
	} else {
		disallowQueryReason = fmt.Sprintf("not a serving tablet type(%v)", newTablet.Type)
		disallowQueryService = disallowQueryReason
	}
	agent.setServicesDesiredState(disallowQueryService)
	if updateBlacklistedTables {
		if err := agent.loadBlacklistRules(ctx, newTablet, blacklistedTables); err != nil {
			// FIXME(alainjobart) how to handle this error?
			log.Errorf("Cannot update blacklisted tables rule: %v", err)
		} else {
			agent.setBlacklistedTables(blacklistedTables)
		}
	}

	broadcastHealth := false
	if allowQuery {
		// Query service should be running.
		if oldTablet.Type == topodatapb.TabletType_REPLICA &&
			newTablet.Type == topodatapb.TabletType_MASTER {
			// When promoting from replica to master, allow both master and replica
			// queries to be served during gracePeriod.
			if _, err := agent.QueryServiceControl.SetServingType(newTablet.Type,
				true, []topodatapb.TabletType{oldTablet.Type}); err == nil {
				// If successful, broadcast to vtgate and then wait.
				agent.broadcastHealth()
				time.Sleep(*gracePeriod)
			} else {
				log.Errorf("Can't start query service for MASTER+REPLICA mode: %v", err)
			}
		}

		if stateChanged, err := agent.QueryServiceControl.SetServingType(newTablet.Type, true, nil); err == nil {
			// If the state changed, broadcast to vtgate.
			// (e.g. this happens when the tablet was already master, but it just
			// changed from NOT_SERVING to SERVING due to
			// "vtctl MigrateServedFrom ... master".)
			if stateChanged {
				broadcastHealth = true
			}
		} else {
			log.Errorf("Cannot start query service: %v", err)
		}
	} else {
		// Query service should be stopped.
		if topo.IsSubjectToLameduck(oldTablet.Type) &&
			newTablet.Type == topodatapb.TabletType_SPARE &&
			*gracePeriod > 0 {
			// When a non-MASTER serving type is going SPARE,
			// put query service in lameduck during gracePeriod.
			agent.lameduck(disallowQueryReason)
		}

		log.Infof("Disabling query service on type change, reason: %v", disallowQueryReason)
		if stateChanged, err := agent.QueryServiceControl.SetServingType(newTablet.Type, false, nil); err == nil {
			// If the state changed, broadcast to vtgate.
			// (e.g. this happens when the tablet was already master, but it just
			// changed from SERVING to NOT_SERVING because filtered replication was
			// enabled.)
			if stateChanged {
				broadcastHealth = true
			}
		} else {
			log.Errorf("SetServingType(serving=false) failed: %v", err)
		}
	}

	// UpdateStream needs to be started or stopped too.
	if topo.IsRunningUpdateStream(newTablet.Type) {
		agent.UpdateStream.Enable()
	} else {
		agent.UpdateStream.Disable()
	}

	// Update the stats to our current type.
	s := topoproto.TabletTypeLString(newTablet.Type)
	statsTabletType.Set(s)
	statsTabletTypeCount.Add(s, 1)

	// See if we need to start or stop vreplication.
	if newTablet.Type == topodatapb.TabletType_MASTER {
		if err := agent.VREngine.Open(agent.batchCtx); err != nil {
			log.Errorf("Could not start VReplication engine: %v. Will keep retrying at health check intervals.", err)
		} else {
			log.Info("VReplication engine started")
		}
	} else {
		agent.VREngine.Close()
	}

	// Broadcast health changes to vtgate immediately.
	if broadcastHealth {
		agent.broadcastHealth()
	}
}

func (agent *ActionAgent) publishState(ctx context.Context) {
	agent.pubMu.Lock()
	defer agent.pubMu.Unlock()
	log.Infof("Publishing state: %v", agent.tablet)
	// If retry is in progress, there's nothing to do.
	if agent.isPublishing {
		return
	}
	// Common code path: publish immediately.
	ctx, cancel := context.WithTimeout(ctx, *topo.RemoteOperationTimeout)
	defer cancel()
	_, err := agent.TopoServer.UpdateTabletFields(ctx, agent.TabletAlias, func(tablet *topodatapb.Tablet) error {
		if err := topotools.CheckOwnership(tablet, agent.tablet); err != nil {
			log.Error(err)
			return topo.NewError(topo.NoUpdateNeeded, "")
		}
		*tablet = *proto.Clone(agent.tablet).(*topodatapb.Tablet)
		return nil
	})
	if err != nil {
		log.Errorf("Unable to publish state to topo, will keep retrying: %v", err)
		agent.isPublishing = true
		// Keep retrying until success.
		go agent.retryPublish()
	}
}

func (agent *ActionAgent) retryPublish() {
	agent.pubMu.Lock()
	defer func() {
		agent.isPublishing = false
		agent.pubMu.Unlock()
	}()

	for {
		// Retry immediately the first time because the previous failure might have been
		// due to an expired context.
		ctx, cancel := context.WithTimeout(agent.batchCtx, *topo.RemoteOperationTimeout)
		_, err := agent.TopoServer.UpdateTabletFields(ctx, agent.TabletAlias, func(tablet *topodatapb.Tablet) error {
			if err := topotools.CheckOwnership(tablet, agent.tablet); err != nil {
				log.Error(err)
				return topo.NewError(topo.NoUpdateNeeded, "")
			}
			*tablet = *proto.Clone(agent.tablet).(*topodatapb.Tablet)
			return nil
		})
		cancel()
		if err != nil {
			log.Errorf("Unable to publish state to topo, will keep retrying: %v", err)
			agent.pubMu.Unlock()
			time.Sleep(*publishRetryInterval)
			agent.pubMu.Lock()
			continue
		}
		log.Infof("Published state: %v", agent.tablet)
		return
	}
}

func (agent *ActionAgent) rebuildKeyspace(keyspace string, retryInterval time.Duration) {
	var srvKeyspace *topodatapb.SrvKeyspace
	defer func() {
		log.Infof("Keyspace rebuilt: %v", keyspace)
		agent.mutex.Lock()
		defer agent.mutex.Unlock()
		agent._srvKeyspace = srvKeyspace
	}()

	// RebuildKeyspace will fail until at least one tablet is up for every shard.
	firstTime := true
	var err error
	for {
		if !firstTime {
			// If keyspace was rebuilt by someone else, we can just exit.
			srvKeyspace, err = agent.TopoServer.GetSrvKeyspace(agent.batchCtx, agent.TabletAlias.Cell, keyspace)
			if err == nil {
				return
			}
		}
		err = topotools.RebuildKeyspace(agent.batchCtx, logutil.NewConsoleLogger(), agent.TopoServer, keyspace, []string{agent.TabletAlias.Cell})
		if err == nil {
			srvKeyspace, err = agent.TopoServer.GetSrvKeyspace(agent.batchCtx, agent.TabletAlias.Cell, keyspace)
			if err == nil {
				return
			}
		}
		if firstTime {
			log.Warningf("rebuildKeyspace failed, will retry every %v: %v", retryInterval, err)
		}
		firstTime = false
		time.Sleep(retryInterval)
	}
}

func (agent *ActionAgent) findMysqlPort(retryInterval time.Duration) {
	for {
		time.Sleep(retryInterval)
		mport, err := agent.MysqlDaemon.GetMysqlPort()
		if err != nil {
			continue
		}
		log.Infof("Identified mysql port: %v", mport)
		agent.pubMu.Lock()
		agent.tablet.MysqlPort = mport
		agent.pubMu.Unlock()
		agent.publishState(agent.batchCtx)
		return
	}
}
