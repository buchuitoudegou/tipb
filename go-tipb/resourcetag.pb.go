// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resourcetag.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResourceGroupTagLabel int32

const (
	ResourceGroupTagLabel_ResourceGroupTagLabelUnknown ResourceGroupTagLabel = 0
	ResourceGroupTagLabel_ResourceGroupTagLabelRow     ResourceGroupTagLabel = 1
	ResourceGroupTagLabel_ResourceGroupTagLabelIndex   ResourceGroupTagLabel = 2
)

var ResourceGroupTagLabel_name = map[int32]string{
	0: "ResourceGroupTagLabelUnknown",
	1: "ResourceGroupTagLabelRow",
	2: "ResourceGroupTagLabelIndex",
}
var ResourceGroupTagLabel_value = map[string]int32{
	"ResourceGroupTagLabelUnknown": 0,
	"ResourceGroupTagLabelRow":     1,
	"ResourceGroupTagLabelIndex":   2,
}

func (x ResourceGroupTagLabel) Enum() *ResourceGroupTagLabel {
	p := new(ResourceGroupTagLabel)
	*p = x
	return p
}
func (x ResourceGroupTagLabel) String() string {
	return proto.EnumName(ResourceGroupTagLabel_name, int32(x))
}
func (x *ResourceGroupTagLabel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResourceGroupTagLabel_value, data, "ResourceGroupTagLabel")
	if err != nil {
		return err
	}
	*x = ResourceGroupTagLabel(value)
	return nil
}
func (ResourceGroupTagLabel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorResourcetag, []int{0}
}

type ResourceGroupTag struct {
	SqlDigest  []byte `protobuf:"bytes,1,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	PlanDigest []byte `protobuf:"bytes,2,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	// Types that are valid to be assigned to TagOneof:
	//	*ResourceGroupTag_Sql
	//	*ResourceGroupTag_BackupCmd
	//	*ResourceGroupTag_RestoreCmd
	//	*ResourceGroupTag_Import
	//	*ResourceGroupTag_CdcJob
	TagOneof isResourceGroupTag_TagOneof `protobuf_oneof:"tag_oneof"`
	// Use to label the handling kv type of the request.
	// This is for TiKV resource_metering to collect execution information by the key label.
	Label            *ResourceGroupTagLabel `protobuf:"varint,8,opt,name=label,enum=tipb.ResourceGroupTagLabel" json:"label,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ResourceGroupTag) Reset()                    { *m = ResourceGroupTag{} }
func (m *ResourceGroupTag) String() string            { return proto.CompactTextString(m) }
func (*ResourceGroupTag) ProtoMessage()               {}
func (*ResourceGroupTag) Descriptor() ([]byte, []int) { return fileDescriptorResourcetag, []int{0} }

type isResourceGroupTag_TagOneof interface {
	isResourceGroupTag_TagOneof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ResourceGroupTag_Sql struct {
	Sql *ResourceGroupTagSQL `protobuf:"bytes,3,opt,name=sql,oneof"`
}
type ResourceGroupTag_BackupCmd struct {
	BackupCmd []byte `protobuf:"bytes,4,opt,name=backup_cmd,json=backupCmd,oneof"`
}
type ResourceGroupTag_RestoreCmd struct {
	RestoreCmd []byte `protobuf:"bytes,5,opt,name=restore_cmd,json=restoreCmd,oneof"`
}
type ResourceGroupTag_Import struct {
	Import []byte `protobuf:"bytes,6,opt,name=import,oneof"`
}
type ResourceGroupTag_CdcJob struct {
	CdcJob []byte `protobuf:"bytes,7,opt,name=cdc_job,json=cdcJob,oneof"`
}

func (*ResourceGroupTag_Sql) isResourceGroupTag_TagOneof()        {}
func (*ResourceGroupTag_BackupCmd) isResourceGroupTag_TagOneof()  {}
func (*ResourceGroupTag_RestoreCmd) isResourceGroupTag_TagOneof() {}
func (*ResourceGroupTag_Import) isResourceGroupTag_TagOneof()     {}
func (*ResourceGroupTag_CdcJob) isResourceGroupTag_TagOneof()     {}

func (m *ResourceGroupTag) GetTagOneof() isResourceGroupTag_TagOneof {
	if m != nil {
		return m.TagOneof
	}
	return nil
}

func (m *ResourceGroupTag) GetSqlDigest() []byte {
	if m != nil {
		return m.SqlDigest
	}
	return nil
}

func (m *ResourceGroupTag) GetPlanDigest() []byte {
	if m != nil {
		return m.PlanDigest
	}
	return nil
}

func (m *ResourceGroupTag) GetSql() *ResourceGroupTagSQL {
	if x, ok := m.GetTagOneof().(*ResourceGroupTag_Sql); ok {
		return x.Sql
	}
	return nil
}

func (m *ResourceGroupTag) GetBackupCmd() []byte {
	if x, ok := m.GetTagOneof().(*ResourceGroupTag_BackupCmd); ok {
		return x.BackupCmd
	}
	return nil
}

func (m *ResourceGroupTag) GetRestoreCmd() []byte {
	if x, ok := m.GetTagOneof().(*ResourceGroupTag_RestoreCmd); ok {
		return x.RestoreCmd
	}
	return nil
}

func (m *ResourceGroupTag) GetImport() []byte {
	if x, ok := m.GetTagOneof().(*ResourceGroupTag_Import); ok {
		return x.Import
	}
	return nil
}

func (m *ResourceGroupTag) GetCdcJob() []byte {
	if x, ok := m.GetTagOneof().(*ResourceGroupTag_CdcJob); ok {
		return x.CdcJob
	}
	return nil
}

func (m *ResourceGroupTag) GetLabel() ResourceGroupTagLabel {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ResourceGroupTagLabel_ResourceGroupTagLabelUnknown
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ResourceGroupTag) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ResourceGroupTag_OneofMarshaler, _ResourceGroupTag_OneofUnmarshaler, _ResourceGroupTag_OneofSizer, []interface{}{
		(*ResourceGroupTag_Sql)(nil),
		(*ResourceGroupTag_BackupCmd)(nil),
		(*ResourceGroupTag_RestoreCmd)(nil),
		(*ResourceGroupTag_Import)(nil),
		(*ResourceGroupTag_CdcJob)(nil),
	}
}

func _ResourceGroupTag_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ResourceGroupTag)
	// tag_oneof
	switch x := m.TagOneof.(type) {
	case *ResourceGroupTag_Sql:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sql); err != nil {
			return err
		}
	case *ResourceGroupTag_BackupCmd:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.BackupCmd)
	case *ResourceGroupTag_RestoreCmd:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.RestoreCmd)
	case *ResourceGroupTag_Import:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Import)
	case *ResourceGroupTag_CdcJob:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.CdcJob)
	case nil:
	default:
		return fmt.Errorf("ResourceGroupTag.TagOneof has unexpected type %T", x)
	}
	return nil
}

func _ResourceGroupTag_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ResourceGroupTag)
	switch tag {
	case 3: // tag_oneof.sql
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResourceGroupTagSQL)
		err := b.DecodeMessage(msg)
		m.TagOneof = &ResourceGroupTag_Sql{msg}
		return true, err
	case 4: // tag_oneof.backup_cmd
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TagOneof = &ResourceGroupTag_BackupCmd{x}
		return true, err
	case 5: // tag_oneof.restore_cmd
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TagOneof = &ResourceGroupTag_RestoreCmd{x}
		return true, err
	case 6: // tag_oneof.import
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TagOneof = &ResourceGroupTag_Import{x}
		return true, err
	case 7: // tag_oneof.cdc_job
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TagOneof = &ResourceGroupTag_CdcJob{x}
		return true, err
	default:
		return false, nil
	}
}

func _ResourceGroupTag_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ResourceGroupTag)
	// tag_oneof
	switch x := m.TagOneof.(type) {
	case *ResourceGroupTag_Sql:
		s := proto.Size(x.Sql)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResourceGroupTag_BackupCmd:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BackupCmd)))
		n += len(x.BackupCmd)
	case *ResourceGroupTag_RestoreCmd:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RestoreCmd)))
		n += len(x.RestoreCmd)
	case *ResourceGroupTag_Import:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Import)))
		n += len(x.Import)
	case *ResourceGroupTag_CdcJob:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.CdcJob)))
		n += len(x.CdcJob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResourceGroupTagSQL struct {
	SqlDigest        []byte                 `protobuf:"bytes,1,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	PlanDigest       []byte                 `protobuf:"bytes,2,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	Label            *ResourceGroupTagLabel `protobuf:"varint,3,opt,name=label,enum=tipb.ResourceGroupTagLabel" json:"label,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ResourceGroupTagSQL) Reset()                    { *m = ResourceGroupTagSQL{} }
func (m *ResourceGroupTagSQL) String() string            { return proto.CompactTextString(m) }
func (*ResourceGroupTagSQL) ProtoMessage()               {}
func (*ResourceGroupTagSQL) Descriptor() ([]byte, []int) { return fileDescriptorResourcetag, []int{1} }

func (m *ResourceGroupTagSQL) GetSqlDigest() []byte {
	if m != nil {
		return m.SqlDigest
	}
	return nil
}

func (m *ResourceGroupTagSQL) GetPlanDigest() []byte {
	if m != nil {
		return m.PlanDigest
	}
	return nil
}

func (m *ResourceGroupTagSQL) GetLabel() ResourceGroupTagLabel {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ResourceGroupTagLabel_ResourceGroupTagLabelUnknown
}

func init() {
	proto.RegisterType((*ResourceGroupTag)(nil), "tipb.ResourceGroupTag")
	proto.RegisterType((*ResourceGroupTagSQL)(nil), "tipb.ResourceGroupTagSQL")
	proto.RegisterEnum("tipb.ResourceGroupTagLabel", ResourceGroupTagLabel_name, ResourceGroupTagLabel_value)
}
func (m *ResourceGroupTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceGroupTag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SqlDigest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.SqlDigest)))
		i += copy(dAtA[i:], m.SqlDigest)
	}
	if m.PlanDigest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.PlanDigest)))
		i += copy(dAtA[i:], m.PlanDigest)
	}
	if m.TagOneof != nil {
		nn1, err := m.TagOneof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Label != nil {
		dAtA[i] = 0x40
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(*m.Label))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourceGroupTag_Sql) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Sql != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(m.Sql.Size()))
		n2, err := m.Sql.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *ResourceGroupTag_BackupCmd) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.BackupCmd != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.BackupCmd)))
		i += copy(dAtA[i:], m.BackupCmd)
	}
	return i, nil
}
func (m *ResourceGroupTag_RestoreCmd) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.RestoreCmd != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.RestoreCmd)))
		i += copy(dAtA[i:], m.RestoreCmd)
	}
	return i, nil
}
func (m *ResourceGroupTag_Import) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Import != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.Import)))
		i += copy(dAtA[i:], m.Import)
	}
	return i, nil
}
func (m *ResourceGroupTag_CdcJob) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CdcJob != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.CdcJob)))
		i += copy(dAtA[i:], m.CdcJob)
	}
	return i, nil
}
func (m *ResourceGroupTagSQL) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceGroupTagSQL) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SqlDigest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.SqlDigest)))
		i += copy(dAtA[i:], m.SqlDigest)
	}
	if m.PlanDigest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(len(m.PlanDigest)))
		i += copy(dAtA[i:], m.PlanDigest)
	}
	if m.Label != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintResourcetag(dAtA, i, uint64(*m.Label))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintResourcetag(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceGroupTag) Size() (n int) {
	var l int
	_ = l
	if m.SqlDigest != nil {
		l = len(m.SqlDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(m.PlanDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.TagOneof != nil {
		n += m.TagOneof.Size()
	}
	if m.Label != nil {
		n += 1 + sovResourcetag(uint64(*m.Label))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourceGroupTag_Sql) Size() (n int) {
	var l int
	_ = l
	if m.Sql != nil {
		l = m.Sql.Size()
		n += 1 + l + sovResourcetag(uint64(l))
	}
	return n
}
func (m *ResourceGroupTag_BackupCmd) Size() (n int) {
	var l int
	_ = l
	if m.BackupCmd != nil {
		l = len(m.BackupCmd)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	return n
}
func (m *ResourceGroupTag_RestoreCmd) Size() (n int) {
	var l int
	_ = l
	if m.RestoreCmd != nil {
		l = len(m.RestoreCmd)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	return n
}
func (m *ResourceGroupTag_Import) Size() (n int) {
	var l int
	_ = l
	if m.Import != nil {
		l = len(m.Import)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	return n
}
func (m *ResourceGroupTag_CdcJob) Size() (n int) {
	var l int
	_ = l
	if m.CdcJob != nil {
		l = len(m.CdcJob)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	return n
}
func (m *ResourceGroupTagSQL) Size() (n int) {
	var l int
	_ = l
	if m.SqlDigest != nil {
		l = len(m.SqlDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(m.PlanDigest)
		n += 1 + l + sovResourcetag(uint64(l))
	}
	if m.Label != nil {
		n += 1 + sovResourcetag(uint64(*m.Label))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResourcetag(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourcetag(x uint64) (n int) {
	return sovResourcetag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceGroupTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourcetag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceGroupTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceGroupTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SqlDigest = append(m.SqlDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.SqlDigest == nil {
				m.SqlDigest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanDigest = append(m.PlanDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.PlanDigest == nil {
				m.PlanDigest = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sql", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResourceGroupTagSQL{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TagOneof = &ResourceGroupTag_Sql{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupCmd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TagOneof = &ResourceGroupTag_BackupCmd{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestoreCmd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TagOneof = &ResourceGroupTag_RestoreCmd{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Import", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TagOneof = &ResourceGroupTag_Import{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdcJob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TagOneof = &ResourceGroupTag_CdcJob{v}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var v ResourceGroupTagLabel
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (ResourceGroupTagLabel(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Label = &v
		default:
			iNdEx = preIndex
			skippy, err := skipResourcetag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourcetag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResourceGroupTagSQL) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourcetag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceGroupTagSQL: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceGroupTagSQL: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SqlDigest = append(m.SqlDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.SqlDigest == nil {
				m.SqlDigest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcetag
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanDigest = append(m.PlanDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.PlanDigest == nil {
				m.PlanDigest = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var v ResourceGroupTagLabel
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (ResourceGroupTagLabel(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Label = &v
		default:
			iNdEx = preIndex
			skippy, err := skipResourcetag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourcetag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResourcetag(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourcetag
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResourcetag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResourcetag
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourcetag
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipResourcetag(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthResourcetag = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourcetag   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("resourcetag.proto", fileDescriptorResourcetag) }

var fileDescriptorResourcetag = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xdf, 0x4e, 0xdb, 0x30,
	0x18, 0xc5, 0xe3, 0xf4, 0xdf, 0xfa, 0x65, 0x9a, 0x32, 0x6f, 0x95, 0xdc, 0xae, 0x4b, 0xd3, 0xee,
	0xa6, 0x9b, 0xb4, 0x4c, 0xab, 0xb4, 0x17, 0xe8, 0x90, 0x28, 0xa8, 0x17, 0x10, 0xe0, 0x3a, 0x4a,
	0x1c, 0x13, 0x85, 0x26, 0xb1, 0x9b, 0xb8, 0x2a, 0x8f, 0xc0, 0x23, 0xf0, 0x44, 0x88, 0x4b, 0x1e,
	0x01, 0x95, 0x17, 0x41, 0x49, 0x53, 0xc4, 0x9f, 0x48, 0xdc, 0xd9, 0xe7, 0xfc, 0xfc, 0xf9, 0x7c,
	0x07, 0x3e, 0xa7, 0x2c, 0xe3, 0xab, 0x94, 0x32, 0xe9, 0x06, 0x96, 0x48, 0xb9, 0xe4, 0xb8, 0x2e,
	0x43, 0xe1, 0xf5, 0xbe, 0x06, 0x3c, 0xe0, 0x85, 0xf0, 0x27, 0x3f, 0x6d, 0xbd, 0xd1, 0x8d, 0x0a,
	0xba, 0x5d, 0xbe, 0xd8, 0x4f, 0xf9, 0x4a, 0x9c, 0xba, 0x01, 0x1e, 0x02, 0x64, 0xcb, 0xc8, 0xf1,
	0xc3, 0x80, 0x65, 0x92, 0x20, 0x13, 0x8d, 0x3f, 0x4e, 0x55, 0x82, 0xec, 0x76, 0xb6, 0x8c, 0xf6,
	0x0a, 0x11, 0xff, 0x00, 0x4d, 0x44, 0x6e, 0xb2, 0x63, 0xd4, 0x27, 0x06, 0x72, 0xb9, 0x84, 0x7e,
	0x43, 0x2d, 0x5b, 0x46, 0xa4, 0x66, 0xa2, 0xb1, 0x36, 0xe9, 0x5a, 0x79, 0x0c, 0xeb, 0xf5, 0x67,
	0x27, 0xc7, 0xf3, 0x99, 0x62, 0xe7, 0x1c, 0x1e, 0x00, 0x78, 0x2e, 0x5d, 0xac, 0x84, 0x43, 0x63,
	0x9f, 0xd4, 0xf3, 0x91, 0x33, 0xc5, 0x6e, 0x6f, 0xb5, 0xff, 0xb1, 0x8f, 0x87, 0xa0, 0xa5, 0x2c,
	0x93, 0x3c, 0x65, 0x05, 0xd1, 0x28, 0x09, 0x28, 0xc5, 0x1c, 0x21, 0xd0, 0x0c, 0x63, 0xc1, 0x53,
	0x49, 0x9a, 0xa5, 0x5b, 0xde, 0x71, 0x17, 0x5a, 0xd4, 0xa7, 0xce, 0x05, 0xf7, 0x48, 0x6b, 0x67,
	0x51, 0x9f, 0x1e, 0x72, 0x0f, 0xff, 0x83, 0x46, 0xe4, 0x7a, 0x2c, 0x22, 0x1f, 0x4c, 0x34, 0xfe,
	0x34, 0xf9, 0x56, 0x9d, 0x74, 0x9e, 0x23, 0xc5, 0x8e, 0x5b, 0x7a, 0xaa, 0x41, 0x5b, 0xba, 0x81,
	0xc3, 0x13, 0xc6, 0xcf, 0x47, 0x57, 0x08, 0xbe, 0x54, 0xec, 0x86, 0xbf, 0xbf, 0xed, 0xf2, 0x79,
	0x8f, 0x83, 0x8a, 0x1e, 0x5f, 0x74, 0xf8, 0x77, 0x97, 0xad, 0xf6, 0x6e, 0xb6, 0x32, 0xd7, 0xaf,
	0x35, 0x74, 0x2a, 0x7d, 0x6c, 0x42, 0xbf, 0xd2, 0x38, 0x4b, 0x16, 0x09, 0x5f, 0x27, 0xba, 0x82,
	0xfb, 0x40, 0xaa, 0x47, 0xf3, 0xb5, 0x8e, 0xb0, 0x01, 0xbd, 0x4a, 0xf7, 0x20, 0xf1, 0xd9, 0xa5,
	0xae, 0x4e, 0x7f, 0xde, 0x6e, 0x0c, 0x74, 0xb7, 0x31, 0xd0, 0xfd, 0xc6, 0x40, 0xd7, 0x0f, 0x86,
	0x02, 0x1d, 0xca, 0x63, 0x4b, 0x84, 0x49, 0x40, 0x5d, 0x61, 0xc9, 0xd0, 0xf7, 0x8a, 0xf8, 0x47,
	0xe8, 0x31, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xe4, 0xc9, 0x32, 0xa7, 0x02, 0x00, 0x00,
}
