/*
Copyright 2020 The Vitess Authors.

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

package engine

import (
	"encoding/json"

	"vitess.io/vitess/go/sqltypes"
	"vitess.io/vitess/go/vt/proto/query"
)

var _ Primitive = (*UpdateTarget)(nil)

// UpdateTarget is an operator to update target string.
type UpdateTarget struct {
	// Target string to be updated
	Target string

	noInputs

	noTxNeeded
}

// MarshalJSON serializes the UpdateTarget into a JSON representation.
// It's used for testing and diagnostics.
func (updTarget *UpdateTarget) MarshalJSON() ([]byte, error) {
	marshalUpdateTarget := struct {
		Opcode string
		Target string
	}{
		Opcode: "UpdateTarget",
		Target: updTarget.Target,
	}
	return json.Marshal(marshalUpdateTarget)
}

// RouteType implements the Primitive interface
func (updTarget UpdateTarget) RouteType() string {
	return "UpdateTarget"
}

// GetKeyspaceName implements the Primitive interface
func (updTarget UpdateTarget) GetKeyspaceName() string {
	return updTarget.Target
}

// GetTableName implements the Primitive interface
func (updTarget UpdateTarget) GetTableName() string {
	return ""
}

// Execute implements the Primitive interface
func (updTarget UpdateTarget) Execute(vcursor VCursor, bindVars map[string]*query.BindVariable, wantfields bool) (*sqltypes.Result, error) {
	err := vcursor.SetTarget(updTarget.Target)
	if err != nil {
		return nil, err
	}
	return &sqltypes.Result{}, nil
}

// StreamExecute implements the Primitive interface
func (updTarget UpdateTarget) StreamExecute(vcursor VCursor, bindVars map[string]*query.BindVariable, wantfields bool, callback func(*sqltypes.Result) error) error {
	panic("implement me")
}

// GetFields implements the Primitive interface
func (updTarget UpdateTarget) GetFields(vcursor VCursor, bindVars map[string]*query.BindVariable) (*sqltypes.Result, error) {
	panic("implement me")
}
