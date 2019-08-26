// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miam/events/v1/events.proto

package eventsv1

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// EventType enumerates all event type values.
type EventType int32

const (
	// Default value when no enumeration is specified.
	EventType_EVENT_TYPE_INVALID EventType = 0
	// Explicitly Unknown object value.
	EventType_EVENT_TYPE_UNKNOWN                   EventType = 1
	EventType_EVENT_TYPE_APPLICATION_CREATED       EventType = 2
	EventType_EVENT_TYPE_APPLICATION_DELETED       EventType = 3
	EventType_EVENT_TYPE_APPLICATION_LABEL_UPDATED EventType = 4
	EventType_EVENT_TYPE_APPLICATION_ACTIVATED     EventType = 5
	EventType_EVENT_TYPE_APPLICATION_DEACTIVATED   EventType = 6
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_INVALID",
	1: "EVENT_TYPE_UNKNOWN",
	2: "EVENT_TYPE_APPLICATION_CREATED",
	3: "EVENT_TYPE_APPLICATION_DELETED",
	4: "EVENT_TYPE_APPLICATION_LABEL_UPDATED",
	5: "EVENT_TYPE_APPLICATION_ACTIVATED",
	6: "EVENT_TYPE_APPLICATION_DEACTIVATED",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_INVALID":                   0,
	"EVENT_TYPE_UNKNOWN":                   1,
	"EVENT_TYPE_APPLICATION_CREATED":       2,
	"EVENT_TYPE_APPLICATION_DELETED":       3,
	"EVENT_TYPE_APPLICATION_LABEL_UPDATED": 4,
	"EVENT_TYPE_APPLICATION_ACTIVATED":     5,
	"EVENT_TYPE_APPLICATION_DEACTIVATED":   6,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dbaf9ff3d9926480, []int{0}
}

// Event wrapper for broker.
type Event struct {
	EventType     EventType  `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=miam.events.v1.EventType" json:"event_type,omitempty"`
	EventId       string     `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateType string     `protobuf:"bytes,3,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	AggregateId   string     `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	Meta          *types.Any `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Event_ApplicationCreated
	Payload              isEvent_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbaf9ff3d9926480, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}

func (m *Event) XXX_Size() int {
	return m.Size()
}

func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Payload interface {
	isEvent_Payload()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event_ApplicationCreated struct {
	ApplicationCreated *ApplicationCreated `protobuf:"bytes,10,opt,name=application_created,json=applicationCreated,proto3,oneof"`
}

func (*Event_ApplicationCreated) isEvent_Payload() {}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_EVENT_TYPE_INVALID
}

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetMeta() *types.Any {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Event) GetApplicationCreated() *ApplicationCreated {
	if x, ok := m.GetPayload().(*Event_ApplicationCreated); ok {
		return x.ApplicationCreated
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_ApplicationCreated)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// payload
	switch x := m.Payload.(type) {
	case *Event_ApplicationCreated:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApplicationCreated); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Payload has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 10: // payload.application_created
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApplicationCreated)
		err := b.DecodeMessage(msg)
		m.Payload = &Event_ApplicationCreated{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// payload
	switch x := m.Payload.(type) {
	case *Event_ApplicationCreated:
		s := proto.Size(x.ApplicationCreated)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ApplicationCreated is raised on application entity creation.
type ApplicationCreated struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationCreated) Reset()         { *m = ApplicationCreated{} }
func (m *ApplicationCreated) String() string { return proto.CompactTextString(m) }
func (*ApplicationCreated) ProtoMessage()    {}
func (*ApplicationCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbaf9ff3d9926480, []int{1}
}

func (m *ApplicationCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *ApplicationCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *ApplicationCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationCreated.Merge(m, src)
}

func (m *ApplicationCreated) XXX_Size() int {
	return m.Size()
}

func (m *ApplicationCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationCreated.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationCreated proto.InternalMessageInfo

func (m *ApplicationCreated) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ApplicationCreated) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func init() {
	proto.RegisterEnum("miam.events.v1.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "miam.events.v1.Event")
	proto.RegisterType((*ApplicationCreated)(nil), "miam.events.v1.ApplicationCreated")
}

func init() { proto.RegisterFile("miam/events/v1/events.proto", fileDescriptor_dbaf9ff3d9926480) }

var fileDescriptor_dbaf9ff3d9926480 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0xbb, 0xf9, 0x69, 0xc9, 0x04, 0xa2, 0x68, 0xa9, 0x90, 0x43, 0xa5, 0x28, 0x58, 0x05,
	0x59, 0x1c, 0x1c, 0xb9, 0x5c, 0x38, 0x70, 0x71, 0x92, 0x95, 0xb0, 0x70, 0x5d, 0xcb, 0x72, 0xcc,
	0x8f, 0x2a, 0x59, 0x9b, 0x7a, 0x31, 0x46, 0x8e, 0x6d, 0x39, 0x6e, 0x24, 0xf3, 0x38, 0x1c, 0x39,
	0xf2, 0x14, 0x88, 0x13, 0x8f, 0x80, 0xf2, 0x1a, 0x5c, 0x90, 0x77, 0xd3, 0x04, 0x52, 0xa5, 0xb7,
	0xd9, 0x6f, 0xbf, 0x99, 0xd5, 0xcc, 0x0e, 0x9c, 0xcc, 0x23, 0x3a, 0x1f, 0xb2, 0x25, 0x4b, 0x8a,
	0xc5, 0x70, 0xa9, 0xad, 0x23, 0x35, 0xcb, 0xd3, 0x22, 0xc5, 0x9d, 0xea, 0x52, 0x5d, 0xa3, 0xa5,
	0xf6, 0xb8, 0x17, 0xa6, 0x69, 0x18, 0xb3, 0x21, 0xbf, 0x9d, 0x5d, 0x7f, 0x1c, 0xd2, 0xa4, 0x14,
	0xaa, 0xfc, 0xbd, 0x06, 0x4d, 0x52, 0x89, 0xf8, 0x25, 0x00, 0xcf, 0xf0, 0x8b, 0x32, 0x63, 0x12,
	0x1a, 0x20, 0xa5, 0x73, 0xd6, 0x53, 0xff, 0xaf, 0xa4, 0x72, 0xd5, 0x2d, 0x33, 0xe6, 0xb4, 0xd8,
	0x4d, 0x88, 0x7b, 0x70, 0x4f, 0x64, 0x46, 0x81, 0x54, 0x1b, 0x20, 0xa5, 0xe5, 0x1c, 0xf1, 0xb3,
	0x11, 0xe0, 0xa7, 0xd0, 0xa1, 0x61, 0x98, 0xb3, 0x90, 0x16, 0x4c, 0x14, 0xae, 0x73, 0xe1, 0xc1,
	0x86, 0xf2, 0x0a, 0x4f, 0xe0, 0xfe, 0x56, 0x8b, 0x02, 0xa9, 0xc1, 0xa5, 0xf6, 0x86, 0x19, 0x01,
	0x56, 0xa0, 0x31, 0x67, 0x05, 0x95, 0x9a, 0x03, 0xa4, 0xb4, 0xcf, 0x8e, 0x55, 0xd1, 0x92, 0x7a,
	0xd3, 0x92, 0xaa, 0x27, 0xa5, 0xc3, 0x0d, 0x3c, 0x85, 0x87, 0x34, 0xcb, 0xe2, 0xe8, 0x8a, 0x16,
	0x51, 0x9a, 0xf8, 0x57, 0x39, 0xa3, 0x05, 0x0b, 0x24, 0xe0, 0x89, 0xf2, 0x6e, 0x47, 0xfa, 0x56,
	0x1d, 0x0b, 0xf3, 0xf5, 0x81, 0x83, 0xe9, 0x2d, 0x3a, 0x6a, 0xc1, 0x51, 0x46, 0xcb, 0x38, 0xa5,
	0x81, 0xfc, 0x0a, 0xf0, 0xed, 0x34, 0xdc, 0x85, 0xfa, 0x75, 0x9e, 0xf0, 0xc9, 0xb5, 0x9c, 0x2a,
	0xc4, 0xc7, 0xd0, 0x8c, 0xe9, 0x8c, 0xc5, 0xeb, 0xa9, 0x88, 0xc3, 0xf3, 0x3f, 0x08, 0x5a, 0x9b,
	0x39, 0xe2, 0x47, 0x80, 0x89, 0x47, 0x2c, 0xd7, 0x77, 0xdf, 0xdb, 0xc4, 0x37, 0x2c, 0x4f, 0x37,
	0x8d, 0x49, 0xf7, 0x60, 0x87, 0x4f, 0xad, 0x37, 0xd6, 0xc5, 0x5b, 0xab, 0x8b, 0xb0, 0x0c, 0xfd,
	0x7f, 0xb8, 0x6e, 0xdb, 0xa6, 0x31, 0xd6, 0x5d, 0xe3, 0xc2, 0xf2, 0xc7, 0x0e, 0xd1, 0x5d, 0x32,
	0xe9, 0xd6, 0xee, 0x70, 0x26, 0xc4, 0x24, 0x95, 0x53, 0xc7, 0x0a, 0x9c, 0xee, 0x71, 0x4c, 0x7d,
	0x44, 0x4c, 0x7f, 0x6a, 0x4f, 0x78, 0xb5, 0x06, 0x3e, 0x85, 0xc1, 0x1e, 0x53, 0x1f, 0xbb, 0x86,
	0xc7, 0xad, 0x26, 0x7e, 0x06, 0xf2, 0xde, 0x37, 0xb7, 0xde, 0xe1, 0xe8, 0xf3, 0x8f, 0x55, 0x1f,
	0xfd, 0x5a, 0xf5, 0xd1, 0xef, 0x55, 0x1f, 0xc1, 0x49, 0x9a, 0x87, 0xea, 0x17, 0x96, 0x44, 0xc5,
	0x27, 0x9a, 0xef, 0x7c, 0xcd, 0xa8, 0xcd, 0xa7, 0xb4, 0xb0, 0xab, 0x2f, 0xb6, 0xd1, 0x07, 0xb1,
	0x64, 0x8b, 0xa5, 0xf6, 0xb5, 0x56, 0x3f, 0x27, 0xef, 0xbe, 0xd5, 0x3a, 0xe7, 0x95, 0x2f, 0x24,
	0xd5, 0xd3, 0x7e, 0x0a, 0x70, 0x29, 0xc0, 0xa5, 0xa7, 0xcd, 0x0e, 0xf9, 0x76, 0xbc, 0xf8, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0xd4, 0x52, 0x05, 0x2d, 0x03, 0x00, 0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EventType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.EventType))
	}
	if len(m.EventId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EventId)))
		i += copy(dAtA[i:], m.EventId)
	}
	if len(m.AggregateType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AggregateType)))
		i += copy(dAtA[i:], m.AggregateType)
	}
	if len(m.AggregateId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AggregateId)))
		i += copy(dAtA[i:], m.AggregateId)
	}
	if m.Meta != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Meta.Size()))
		n1, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Payload != nil {
		nn2, err := m.Payload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Event_ApplicationCreated) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ApplicationCreated != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.ApplicationCreated.Size()))
		n3, err := m.ApplicationCreated.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *ApplicationCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationCreated) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Urn) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Urn)))
		i += copy(dAtA[i:], m.Urn)
	}
	if len(m.Label) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventType != 0 {
		n += 1 + sovEvents(uint64(m.EventType))
	}
	l = len(m.EventId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.AggregateType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.AggregateId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Payload != nil {
		n += m.Payload.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Event_ApplicationCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplicationCreated != nil {
		l = m.ApplicationCreated.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *ApplicationCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Urn)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= EventType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &types.Any{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationCreated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ApplicationCreated{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &Event_ApplicationCreated{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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

func (m *ApplicationCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApplicationCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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

func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvents
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
				next, err := skipEvents(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEvents
				}
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
	ErrInvalidLengthEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents   = fmt.Errorf("proto: integer overflow")
)
