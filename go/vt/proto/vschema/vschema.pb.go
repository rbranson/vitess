// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vschema.proto

package vschema // import "vitess.io/vitess/go/vt/proto/vschema"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "vitess.io/vitess/go/vt/proto/query"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RoutingRules specify the high level routing rules for the VSchema.
type RoutingRules struct {
	// rules should ideally be a map. However protos dont't allow
	// repeated fields as elements of a map. So, we use a list
	// instead.
	Rules                []*RoutingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoutingRules) Reset()         { *m = RoutingRules{} }
func (m *RoutingRules) String() string { return proto.CompactTextString(m) }
func (*RoutingRules) ProtoMessage()    {}
func (*RoutingRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{0}
}
func (m *RoutingRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingRules.Unmarshal(m, b)
}
func (m *RoutingRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingRules.Marshal(b, m, deterministic)
}
func (dst *RoutingRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingRules.Merge(dst, src)
}
func (m *RoutingRules) XXX_Size() int {
	return xxx_messageInfo_RoutingRules.Size(m)
}
func (m *RoutingRules) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingRules.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingRules proto.InternalMessageInfo

func (m *RoutingRules) GetRules() []*RoutingRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// RoutingRule specifies a routing rule.
type RoutingRule struct {
	FromTable            string   `protobuf:"bytes,1,opt,name=from_table,json=fromTable,proto3" json:"from_table,omitempty"`
	ToTables             []string `protobuf:"bytes,2,rep,name=to_tables,json=toTables,proto3" json:"to_tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutingRule) Reset()         { *m = RoutingRule{} }
func (m *RoutingRule) String() string { return proto.CompactTextString(m) }
func (*RoutingRule) ProtoMessage()    {}
func (*RoutingRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{1}
}
func (m *RoutingRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingRule.Unmarshal(m, b)
}
func (m *RoutingRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingRule.Marshal(b, m, deterministic)
}
func (dst *RoutingRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingRule.Merge(dst, src)
}
func (m *RoutingRule) XXX_Size() int {
	return xxx_messageInfo_RoutingRule.Size(m)
}
func (m *RoutingRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingRule.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingRule proto.InternalMessageInfo

func (m *RoutingRule) GetFromTable() string {
	if m != nil {
		return m.FromTable
	}
	return ""
}

func (m *RoutingRule) GetToTables() []string {
	if m != nil {
		return m.ToTables
	}
	return nil
}

// Keyspace is the vschema for a keyspace.
type Keyspace struct {
	// If sharded is false, vindexes and tables are ignored.
	Sharded              bool               `protobuf:"varint,1,opt,name=sharded,proto3" json:"sharded,omitempty"`
	Vindexes             map[string]*Vindex `protobuf:"bytes,2,rep,name=vindexes,proto3" json:"vindexes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tables               map[string]*Table  `protobuf:"bytes,3,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Keyspace) Reset()         { *m = Keyspace{} }
func (m *Keyspace) String() string { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()    {}
func (*Keyspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{2}
}
func (m *Keyspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyspace.Unmarshal(m, b)
}
func (m *Keyspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyspace.Marshal(b, m, deterministic)
}
func (dst *Keyspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyspace.Merge(dst, src)
}
func (m *Keyspace) XXX_Size() int {
	return xxx_messageInfo_Keyspace.Size(m)
}
func (m *Keyspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyspace.DiscardUnknown(m)
}

var xxx_messageInfo_Keyspace proto.InternalMessageInfo

func (m *Keyspace) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *Keyspace) GetVindexes() map[string]*Vindex {
	if m != nil {
		return m.Vindexes
	}
	return nil
}

func (m *Keyspace) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

// Vindex is the vindex info for a Keyspace.
type Vindex struct {
	// The type must match one of the predefined
	// (or plugged in) vindex names.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// params is a map of attribute value pairs
	// that must be defined as required by the
	// vindex constructors. The values can only
	// be strings.
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A lookup vindex can have an owner table defined.
	// If so, rows in the lookup table are created or
	// deleted in sync with corresponding rows in the
	// owner table.
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vindex) Reset()         { *m = Vindex{} }
func (m *Vindex) String() string { return proto.CompactTextString(m) }
func (*Vindex) ProtoMessage()    {}
func (*Vindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{3}
}
func (m *Vindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vindex.Unmarshal(m, b)
}
func (m *Vindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vindex.Marshal(b, m, deterministic)
}
func (dst *Vindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vindex.Merge(dst, src)
}
func (m *Vindex) XXX_Size() int {
	return xxx_messageInfo_Vindex.Size(m)
}
func (m *Vindex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vindex.DiscardUnknown(m)
}

var xxx_messageInfo_Vindex proto.InternalMessageInfo

func (m *Vindex) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Vindex) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Vindex) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// Table is the table info for a Keyspace.
type Table struct {
	// If the table is a sequence, type must be
	// "sequence". Otherwise, it should be empty.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// column_vindexes associates columns to vindexes.
	ColumnVindexes []*ColumnVindex `protobuf:"bytes,2,rep,name=column_vindexes,json=columnVindexes,proto3" json:"column_vindexes,omitempty"`
	// auto_increment is specified if a column needs
	// to be associated with a sequence.
	AutoIncrement *AutoIncrement `protobuf:"bytes,3,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	// columns lists the columns for the table.
	Columns []*Column `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	// pinned pins an unsharded table to a specific
	// shard, as dictated by the keyspace id.
	// The keyspace id is represened in hex form
	// like in keyranges.
	Pinned string `protobuf:"bytes,5,opt,name=pinned,proto3" json:"pinned,omitempty"`
	// column_list_authoritative is set to true if columns is
	// an authoritative list for the table. This allows
	// us to expand 'select *' expressions.
	ColumnListAuthoritative bool     `protobuf:"varint,6,opt,name=column_list_authoritative,json=columnListAuthoritative,proto3" json:"column_list_authoritative,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{4}
}
func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (dst *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(dst, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Table) GetColumnVindexes() []*ColumnVindex {
	if m != nil {
		return m.ColumnVindexes
	}
	return nil
}

func (m *Table) GetAutoIncrement() *AutoIncrement {
	if m != nil {
		return m.AutoIncrement
	}
	return nil
}

func (m *Table) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Table) GetPinned() string {
	if m != nil {
		return m.Pinned
	}
	return ""
}

func (m *Table) GetColumnListAuthoritative() bool {
	if m != nil {
		return m.ColumnListAuthoritative
	}
	return false
}

// ColumnVindex is used to associate a column to a vindex.
type ColumnVindex struct {
	// Legacy implemenation, moving forward all vindexes should define a list of columns.
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// The name must match a vindex defined in Keyspace.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// List of columns that define this Vindex
	Columns              []string `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnVindex) Reset()         { *m = ColumnVindex{} }
func (m *ColumnVindex) String() string { return proto.CompactTextString(m) }
func (*ColumnVindex) ProtoMessage()    {}
func (*ColumnVindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{5}
}
func (m *ColumnVindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnVindex.Unmarshal(m, b)
}
func (m *ColumnVindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnVindex.Marshal(b, m, deterministic)
}
func (dst *ColumnVindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnVindex.Merge(dst, src)
}
func (m *ColumnVindex) XXX_Size() int {
	return xxx_messageInfo_ColumnVindex.Size(m)
}
func (m *ColumnVindex) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnVindex.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnVindex proto.InternalMessageInfo

func (m *ColumnVindex) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *ColumnVindex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnVindex) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Autoincrement is used to designate a column as auto-inc.
type AutoIncrement struct {
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// The sequence must match a table of type SEQUENCE.
	Sequence             string   `protobuf:"bytes,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoIncrement) Reset()         { *m = AutoIncrement{} }
func (m *AutoIncrement) String() string { return proto.CompactTextString(m) }
func (*AutoIncrement) ProtoMessage()    {}
func (*AutoIncrement) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{6}
}
func (m *AutoIncrement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoIncrement.Unmarshal(m, b)
}
func (m *AutoIncrement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoIncrement.Marshal(b, m, deterministic)
}
func (dst *AutoIncrement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoIncrement.Merge(dst, src)
}
func (m *AutoIncrement) XXX_Size() int {
	return xxx_messageInfo_AutoIncrement.Size(m)
}
func (m *AutoIncrement) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoIncrement.DiscardUnknown(m)
}

var xxx_messageInfo_AutoIncrement proto.InternalMessageInfo

func (m *AutoIncrement) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *AutoIncrement) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

// Column describes a column.
type Column struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 query.Type `protobuf:"varint,2,opt,name=type,proto3,enum=query.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{7}
}
func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (dst *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(dst, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetType() query.Type {
	if m != nil {
		return m.Type
	}
	return query.Type_NULL_TYPE
}

// SrvVSchema is the roll-up of all the Keyspace schema for a cell.
type SrvVSchema struct {
	// keyspaces is a map of keyspace name -> Keyspace object.
	Keyspaces            map[string]*Keyspace `protobuf:"bytes,1,rep,name=keyspaces,proto3" json:"keyspaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RoutingRules         *RoutingRules        `protobuf:"bytes,2,opt,name=routing_rules,json=routingRules,proto3" json:"routing_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SrvVSchema) Reset()         { *m = SrvVSchema{} }
func (m *SrvVSchema) String() string { return proto.CompactTextString(m) }
func (*SrvVSchema) ProtoMessage()    {}
func (*SrvVSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_ddae95e7e0992d00, []int{8}
}
func (m *SrvVSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvVSchema.Unmarshal(m, b)
}
func (m *SrvVSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvVSchema.Marshal(b, m, deterministic)
}
func (dst *SrvVSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvVSchema.Merge(dst, src)
}
func (m *SrvVSchema) XXX_Size() int {
	return xxx_messageInfo_SrvVSchema.Size(m)
}
func (m *SrvVSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvVSchema.DiscardUnknown(m)
}

var xxx_messageInfo_SrvVSchema proto.InternalMessageInfo

func (m *SrvVSchema) GetKeyspaces() map[string]*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func (m *SrvVSchema) GetRoutingRules() *RoutingRules {
	if m != nil {
		return m.RoutingRules
	}
	return nil
}

func init() {
	proto.RegisterType((*RoutingRules)(nil), "vschema.RoutingRules")
	proto.RegisterType((*RoutingRule)(nil), "vschema.RoutingRule")
	proto.RegisterType((*Keyspace)(nil), "vschema.Keyspace")
	proto.RegisterMapType((map[string]*Table)(nil), "vschema.Keyspace.TablesEntry")
	proto.RegisterMapType((map[string]*Vindex)(nil), "vschema.Keyspace.VindexesEntry")
	proto.RegisterType((*Vindex)(nil), "vschema.Vindex")
	proto.RegisterMapType((map[string]string)(nil), "vschema.Vindex.ParamsEntry")
	proto.RegisterType((*Table)(nil), "vschema.Table")
	proto.RegisterType((*ColumnVindex)(nil), "vschema.ColumnVindex")
	proto.RegisterType((*AutoIncrement)(nil), "vschema.AutoIncrement")
	proto.RegisterType((*Column)(nil), "vschema.Column")
	proto.RegisterType((*SrvVSchema)(nil), "vschema.SrvVSchema")
	proto.RegisterMapType((map[string]*Keyspace)(nil), "vschema.SrvVSchema.KeyspacesEntry")
}

func init() { proto.RegisterFile("vschema.proto", fileDescriptor_vschema_ddae95e7e0992d00) }

var fileDescriptor_vschema_ddae95e7e0992d00 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xda, 0x35, 0x6b, 0x4f, 0xd6, 0x0e, 0xac, 0x6d, 0x84, 0x4e, 0xd3, 0xaa, 0x68, 0x40,
	0xe1, 0xa2, 0x95, 0x3a, 0x21, 0x41, 0xd1, 0x10, 0x63, 0xe2, 0x62, 0x62, 0x12, 0x28, 0x9b, 0x76,
	0xc1, 0x4d, 0xe5, 0xb5, 0x66, 0x8b, 0xd6, 0xc6, 0x99, 0xed, 0x04, 0xf2, 0x28, 0xdc, 0xf2, 0x06,
	0x3c, 0x0f, 0x2f, 0x83, 0xe2, 0x9f, 0xcc, 0xe9, 0xca, 0x9d, 0x3f, 0x9f, 0xf3, 0x7d, 0xe7, 0xf3,
	0xb1, 0x7d, 0xa0, 0x9d, 0xf1, 0xe9, 0x0d, 0x59, 0xe0, 0x41, 0xc2, 0xa8, 0xa0, 0x68, 0x5d, 0xc3,
	0xae, 0x77, 0x97, 0x12, 0x96, 0xab, 0xdd, 0x60, 0x0c, 0x1b, 0x21, 0x4d, 0x45, 0x14, 0x5f, 0x87,
	0xe9, 0x9c, 0x70, 0xf4, 0x0a, 0x1a, 0xac, 0x58, 0xf8, 0x4e, 0xaf, 0xde, 0xf7, 0x46, 0x5b, 0x03,
	0x23, 0x62, 0x65, 0x85, 0x2a, 0x25, 0x38, 0x05, 0xcf, 0xda, 0x45, 0x7b, 0x00, 0xdf, 0x19, 0x5d,
	0x4c, 0x04, 0xbe, 0x9a, 0x13, 0xdf, 0xe9, 0x39, 0xfd, 0x56, 0xd8, 0x2a, 0x76, 0x2e, 0x8a, 0x0d,
	0xb4, 0x0b, 0x2d, 0x41, 0x55, 0x90, 0xfb, 0xb5, 0x5e, 0xbd, 0xdf, 0x0a, 0x9b, 0x82, 0xca, 0x18,
	0x0f, 0xfe, 0xd4, 0xa0, 0xf9, 0x99, 0xe4, 0x3c, 0xc1, 0x53, 0x82, 0x7c, 0x58, 0xe7, 0x37, 0x98,
	0xcd, 0xc8, 0x4c, 0xaa, 0x34, 0x43, 0x03, 0xd1, 0x3b, 0x68, 0x66, 0x51, 0x3c, 0x23, 0x3f, 0xb5,
	0x84, 0x37, 0xda, 0x2f, 0x0d, 0x1a, 0xfa, 0xe0, 0x52, 0x67, 0x7c, 0x8a, 0x05, 0xcb, 0xc3, 0x92,
	0x80, 0x5e, 0x83, 0xab, 0xab, 0xd7, 0x25, 0x75, 0xef, 0x21, 0x55, 0xb9, 0x51, 0x44, 0x9d, 0xdc,
	0x3d, 0x83, 0x76, 0x45, 0x11, 0x3d, 0x82, 0xfa, 0x2d, 0xc9, 0xf5, 0x01, 0x8b, 0x25, 0x7a, 0x06,
	0x8d, 0x0c, 0xcf, 0x53, 0xe2, 0xd7, 0x7a, 0x4e, 0xdf, 0x1b, 0x6d, 0x96, 0xc2, 0x8a, 0x18, 0xaa,
	0xe8, 0xb8, 0xf6, 0xc6, 0xe9, 0x9e, 0x82, 0x67, 0x15, 0x59, 0xa1, 0x75, 0x50, 0xd5, 0xea, 0x94,
	0x5a, 0x92, 0x66, 0x49, 0x05, 0xbf, 0x1d, 0x70, 0x55, 0x01, 0x84, 0x60, 0x4d, 0xe4, 0x89, 0x69,
	0xba, 0x5c, 0xa3, 0x43, 0x70, 0x13, 0xcc, 0xf0, 0xc2, 0x74, 0x6a, 0x77, 0xc9, 0xd5, 0xe0, 0xab,
	0x8c, 0xea, 0xc3, 0xaa, 0x54, 0xb4, 0x05, 0x0d, 0xfa, 0x23, 0x26, 0xcc, 0xaf, 0x4b, 0x25, 0x05,
	0xba, 0x6f, 0xc1, 0xb3, 0x92, 0x57, 0x98, 0xde, 0xb2, 0x4d, 0xb7, 0x6c, 0x93, 0xbf, 0x6a, 0xd0,
	0x50, 0xf7, 0xbf, 0xca, 0xe3, 0x7b, 0xd8, 0x9c, 0xd2, 0x79, 0xba, 0x88, 0x27, 0x4b, 0xd7, 0xba,
	0x5d, 0x9a, 0x3d, 0x91, 0x71, 0xdd, 0xc8, 0xce, 0xd4, 0x42, 0x84, 0xa3, 0x23, 0xe8, 0xe0, 0x54,
	0xd0, 0x49, 0x14, 0x4f, 0x19, 0x59, 0x90, 0x58, 0x48, 0xdf, 0xde, 0x68, 0xa7, 0xa4, 0x1f, 0xa7,
	0x82, 0x9e, 0x9a, 0x68, 0xd8, 0xc6, 0x36, 0x44, 0x2f, 0x61, 0x5d, 0x09, 0x72, 0x7f, 0x4d, 0x96,
	0xdd, 0x5c, 0x2a, 0x1b, 0x9a, 0x38, 0xda, 0x01, 0x37, 0x89, 0xe2, 0x98, 0xcc, 0xfc, 0x86, 0xf4,
	0xaf, 0x11, 0x1a, 0xc3, 0x53, 0x7d, 0x82, 0x79, 0xc4, 0xc5, 0x04, 0xa7, 0xe2, 0x86, 0xb2, 0x48,
	0x60, 0x11, 0x65, 0xc4, 0x77, 0xe5, 0xeb, 0x7d, 0xa2, 0x12, 0xce, 0x22, 0x2e, 0x8e, 0xed, 0x70,
	0x70, 0x01, 0x1b, 0xf6, 0xe9, 0x8a, 0x1a, 0x2a, 0x55, 0xf7, 0x48, 0xa3, 0xa2, 0x73, 0x31, 0x5e,
	0x98, 0xe6, 0xca, 0x75, 0xf1, 0x47, 0x8c, 0xf5, 0xba, 0xfc, 0x4b, 0x06, 0x06, 0x27, 0xd0, 0xae,
	0x1c, 0xfa, 0xbf, 0xb2, 0x5d, 0x68, 0x72, 0x72, 0x97, 0x92, 0x78, 0x6a, 0xa4, 0x4b, 0x1c, 0x1c,
	0x81, 0x7b, 0x52, 0x2d, 0xee, 0x58, 0xc5, 0xf7, 0xf5, 0x55, 0x16, 0xac, 0xce, 0xc8, 0x1b, 0xa8,
	0x81, 0x72, 0x91, 0x27, 0x44, 0xdd, 0x6b, 0xf0, 0xd7, 0x01, 0x38, 0x67, 0xd9, 0xe5, 0xb9, 0x6c,
	0x26, 0xfa, 0x00, 0xad, 0x5b, 0xfd, 0xc5, 0xcc, 0x60, 0x09, 0xca, 0x4e, 0xdf, 0xe7, 0x95, 0xff,
	0x50, 0x3f, 0xca, 0x7b, 0x12, 0x1a, 0x43, 0x9b, 0xa9, 0x51, 0x33, 0x51, 0xe3, 0x49, 0xfd, 0x8e,
	0xed, 0x55, 0xe3, 0x89, 0x87, 0x1b, 0xcc, 0x42, 0xdd, 0x2f, 0xd0, 0xa9, 0x0a, 0xaf, 0x78, 0xc0,
	0x2f, 0xaa, 0xbf, 0xee, 0xf1, 0x83, 0xd1, 0x60, 0xbd, 0xe9, 0x8f, 0xcf, 0xbf, 0x1d, 0x64, 0x91,
	0x20, 0x9c, 0x0f, 0x22, 0x3a, 0x54, 0xab, 0xe1, 0x35, 0x1d, 0x66, 0x62, 0x28, 0x67, 0xea, 0x50,
	0x73, 0xaf, 0x5c, 0x09, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x0d, 0x93, 0x48, 0x89,
	0x05, 0x00, 0x00,
}
