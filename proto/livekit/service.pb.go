// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package livekit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateRoomRequest struct {
	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// number of seconds the room should cleanup after being empty
	EmptyTimeout         uint32   `protobuf:"varint,2,opt,name=empty_timeout,json=emptyTimeout,proto3" json:"empty_timeout,omitempty"`
	MaxParticipants      uint32   `protobuf:"varint,3,opt,name=max_participants,json=maxParticipants,proto3" json:"max_participants,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *CreateRoomRequest) GetEmptyTimeout() uint32 {
	if m != nil {
		return m.EmptyTimeout
	}
	return 0
}

func (m *CreateRoomRequest) GetMaxParticipants() uint32 {
	if m != nil {
		return m.MaxParticipants
	}
	return 0
}

type GetRoomRequest struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoomRequest) Reset()         { *m = GetRoomRequest{} }
func (m *GetRoomRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoomRequest) ProtoMessage()    {}
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomRequest.Unmarshal(m, b)
}
func (m *GetRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomRequest.Marshal(b, m, deterministic)
}
func (m *GetRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomRequest.Merge(m, src)
}
func (m *GetRoomRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoomRequest.Size(m)
}
func (m *GetRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomRequest proto.InternalMessageInfo

func (m *GetRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type RoomInfo struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	NodeIp               string   `protobuf:"bytes,2,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	NodeRtcPort          uint32   `protobuf:"varint,3,opt,name=node_rtc_port,json=nodeRtcPort,proto3" json:"node_rtc_port,omitempty"`
	CreationTime         int64    `protobuf:"varint,4,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *RoomInfo) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *RoomInfo) GetNodeRtcPort() uint32 {
	if m != nil {
		return m.NodeRtcPort
	}
	return 0
}

func (m *RoomInfo) GetCreationTime() int64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *RoomInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteRoomRequest struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomRequest) Reset()         { *m = DeleteRoomRequest{} }
func (m *DeleteRoomRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomRequest) ProtoMessage()    {}
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *DeleteRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomRequest.Unmarshal(m, b)
}
func (m *DeleteRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomRequest.Merge(m, src)
}
func (m *DeleteRoomRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomRequest.Size(m)
}
func (m *DeleteRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomRequest proto.InternalMessageInfo

func (m *DeleteRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type DeleteRoomResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomResponse) Reset()         { *m = DeleteRoomResponse{} }
func (m *DeleteRoomResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomResponse) ProtoMessage()    {}
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *DeleteRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomResponse.Unmarshal(m, b)
}
func (m *DeleteRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomResponse.Merge(m, src)
}
func (m *DeleteRoomResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomResponse.Size(m)
}
func (m *DeleteRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomResponse proto.InternalMessageInfo

type JoinRequest struct {
	RoomId               string              `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Token                string              `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PeerId               string              `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Offer                *SessionDescription `protobuf:"bytes,4,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *JoinRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *JoinRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *JoinRequest) GetOffer() *SessionDescription {
	if m != nil {
		return m.Offer
	}
	return nil
}

type JoinResponse struct {
	Answer               *SessionDescription `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetAnswer() *SessionDescription {
	if m != nil {
		return m.Answer
	}
	return nil
}

type TrickleRequest struct {
	Candidate            string   `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrickleRequest) Reset()         { *m = TrickleRequest{} }
func (m *TrickleRequest) String() string { return proto.CompactTextString(m) }
func (*TrickleRequest) ProtoMessage()    {}
func (*TrickleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *TrickleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrickleRequest.Unmarshal(m, b)
}
func (m *TrickleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrickleRequest.Marshal(b, m, deterministic)
}
func (m *TrickleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrickleRequest.Merge(m, src)
}
func (m *TrickleRequest) XXX_Size() int {
	return xxx_messageInfo_TrickleRequest.Size(m)
}
func (m *TrickleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrickleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrickleRequest proto.InternalMessageInfo

func (m *TrickleRequest) GetCandidate() string {
	if m != nil {
		return m.Candidate
	}
	return ""
}

type TrickleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrickleResponse) Reset()         { *m = TrickleResponse{} }
func (m *TrickleResponse) String() string { return proto.CompactTextString(m) }
func (*TrickleResponse) ProtoMessage()    {}
func (*TrickleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *TrickleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrickleResponse.Unmarshal(m, b)
}
func (m *TrickleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrickleResponse.Marshal(b, m, deterministic)
}
func (m *TrickleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrickleResponse.Merge(m, src)
}
func (m *TrickleResponse) XXX_Size() int {
	return xxx_messageInfo_TrickleResponse.Size(m)
}
func (m *TrickleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrickleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrickleResponse proto.InternalMessageInfo

type SessionDescription struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sdp                  []byte   `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDescription) Reset()         { *m = SessionDescription{} }
func (m *SessionDescription) String() string { return proto.CompactTextString(m) }
func (*SessionDescription) ProtoMessage()    {}
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *SessionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDescription.Unmarshal(m, b)
}
func (m *SessionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDescription.Marshal(b, m, deterministic)
}
func (m *SessionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDescription.Merge(m, src)
}
func (m *SessionDescription) XXX_Size() int {
	return xxx_messageInfo_SessionDescription.Size(m)
}
func (m *SessionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDescription proto.InternalMessageInfo

func (m *SessionDescription) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionDescription) GetSdp() []byte {
	if m != nil {
		return m.Sdp
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoomRequest)(nil), "livekit.CreateRoomRequest")
	proto.RegisterType((*GetRoomRequest)(nil), "livekit.GetRoomRequest")
	proto.RegisterType((*RoomInfo)(nil), "livekit.RoomInfo")
	proto.RegisterType((*DeleteRoomRequest)(nil), "livekit.DeleteRoomRequest")
	proto.RegisterType((*DeleteRoomResponse)(nil), "livekit.DeleteRoomResponse")
	proto.RegisterType((*JoinRequest)(nil), "livekit.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "livekit.JoinResponse")
	proto.RegisterType((*TrickleRequest)(nil), "livekit.TrickleRequest")
	proto.RegisterType((*TrickleResponse)(nil), "livekit.TrickleResponse")
	proto.RegisterType((*SessionDescription)(nil), "livekit.SessionDescription")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x2f, 0x32, 0x49, 0xda, 0x66, 0x15, 0x14, 0xcb, 0xe5, 0x10, 0x99, 0x4b, 0x2a,
	0xa1, 0x48, 0x24, 0xe2, 0x02, 0x5c, 0x20, 0x45, 0x28, 0x5c, 0xa8, 0xdc, 0x9c, 0xb8, 0x44, 0xc6,
	0x9e, 0x48, 0xab, 0xc4, 0xbb, 0xcb, 0xee, 0xb6, 0xb4, 0x12, 0x3f, 0x80, 0xff, 0xc0, 0xbf, 0xe1,
	0xc4, 0xcf, 0x42, 0xbb, 0x6b, 0x3b, 0x2e, 0x29, 0x6d, 0x6f, 0x9e, 0x99, 0xe7, 0xdd, 0xf7, 0xe6,
	0x3d, 0x2d, 0xf4, 0x14, 0xca, 0x2b, 0x9a, 0xe0, 0x44, 0x48, 0xae, 0x39, 0x69, 0x6d, 0xe9, 0x15,
	0x6e, 0xa8, 0x0e, 0x7f, 0x40, 0x7f, 0x2e, 0x31, 0xd6, 0x18, 0x71, 0x9e, 0x45, 0xf8, 0xed, 0x12,
	0x95, 0x26, 0x43, 0x68, 0x49, 0xce, 0xb3, 0x15, 0x4d, 0x7d, 0x6f, 0xe4, 0x8d, 0xdb, 0x51, 0xd3,
	0x94, 0x8b, 0x94, 0x3c, 0x87, 0x1e, 0x66, 0x42, 0xdf, 0xac, 0x34, 0xcd, 0x90, 0x5f, 0x6a, 0xff,
	0x60, 0xe4, 0x8d, 0x7b, 0x51, 0xd7, 0x36, 0x97, 0xae, 0x47, 0x4e, 0xe1, 0x38, 0x8b, 0xaf, 0x57,
	0x22, 0x96, 0x9a, 0x26, 0x54, 0xc4, 0x4c, 0x2b, 0xbf, 0x66, 0x71, 0x47, 0x59, 0x7c, 0x7d, 0x5e,
	0x69, 0x87, 0xa7, 0x70, 0xf8, 0x11, 0xf5, 0x63, 0xae, 0x0e, 0x7f, 0x79, 0xf0, 0xc4, 0x00, 0x17,
	0x6c, 0xcd, 0xff, 0x4f, 0x70, 0x08, 0x2d, 0xc6, 0x53, 0x5c, 0x51, 0x61, 0xa9, 0xb5, 0xa3, 0xa6,
	0x29, 0x17, 0x82, 0x84, 0xd0, 0xb3, 0x03, 0xa9, 0x93, 0x95, 0xe0, 0x52, 0xe7, 0x8c, 0x3a, 0xa6,
	0x19, 0xe9, 0xe4, 0x9c, 0x4b, 0x6d, 0xd4, 0x25, 0x66, 0x17, 0x94, 0x33, 0x2b, 0xd0, 0xaf, 0x8f,
	0xbc, 0x71, 0x2d, 0xea, 0x16, 0x4d, 0x23, 0x90, 0x0c, 0xa0, 0xa1, 0xf9, 0x06, 0x99, 0xdf, 0xb0,
	0xe7, 0xbb, 0x22, 0x7c, 0x01, 0xfd, 0x33, 0xdc, 0xe2, 0xe3, 0xd6, 0x18, 0x0e, 0x80, 0x54, 0xd1,
	0x4a, 0x70, 0xa6, 0x30, 0xfc, 0xe9, 0x41, 0xe7, 0x13, 0xa7, 0xec, 0x41, 0x17, 0x4a, 0x0a, 0x07,
	0x15, 0x0a, 0x06, 0x2e, 0x10, 0xa5, 0x81, 0xd7, 0x1c, 0xdc, 0x94, 0x8b, 0x94, 0xbc, 0x84, 0x06,
	0x5f, 0xaf, 0x51, 0x5a, 0x39, 0x9d, 0xe9, 0xc9, 0x24, 0xf7, 0x7e, 0x72, 0x81, 0x4a, 0x51, 0xce,
	0xce, 0x50, 0x25, 0x92, 0x0a, 0xa3, 0x30, 0x72, 0xc8, 0x70, 0x0e, 0x5d, 0xc7, 0xc4, 0x51, 0x23,
	0x33, 0x68, 0xc6, 0x4c, 0x7d, 0x47, 0x69, 0x99, 0x3c, 0x70, 0x46, 0x0e, 0x0d, 0x27, 0x70, 0xb8,
	0x94, 0x34, 0xd9, 0x6c, 0xb1, 0x50, 0xf4, 0x0c, 0xda, 0x49, 0xcc, 0x52, 0x9a, 0xc6, 0x1a, 0x73,
	0x4d, 0xbb, 0x46, 0xd8, 0x87, 0xa3, 0x12, 0x9f, 0xaf, 0xe4, 0x35, 0x90, 0xfd, 0x0b, 0x08, 0x81,
	0xba, 0xbe, 0x11, 0xc5, 0x09, 0xf6, 0x9b, 0x1c, 0x43, 0x4d, 0xa5, 0xce, 0xf4, 0x6e, 0x64, 0x3e,
	0xa7, 0x7f, 0x3c, 0xe8, 0x98, 0xfd, 0x5e, 0xb8, 0xe0, 0x93, 0x37, 0x00, 0xbb, 0xa4, 0x93, 0xa0,
	0x54, 0xb0, 0x17, 0xff, 0xa0, 0x5f, 0xce, 0xca, 0xc0, 0xbd, 0x82, 0x56, 0x1e, 0x54, 0x32, 0x2c,
	0xa7, 0xb7, 0xa3, 0x7b, 0xd7, 0x6f, 0x1f, 0x00, 0x76, 0x46, 0x57, 0xee, 0xdc, 0xcb, 0x4a, 0x70,
	0x72, 0xe7, 0xcc, 0xad, 0x61, 0xfa, 0xdb, 0x03, 0x88, 0x96, 0xf3, 0x42, 0xc9, 0x0c, 0xea, 0xc6,
	0x1d, 0x32, 0x28, 0xff, 0xa9, 0xc4, 0x26, 0x78, 0xfa, 0x4f, 0x37, 0xb7, 0xf0, 0x1d, 0x34, 0x3e,
	0x1b, 0x6f, 0xc9, 0x7d, 0xde, 0x05, 0xf7, 0x0d, 0xc9, 0x5b, 0x68, 0xe5, 0x06, 0x55, 0x96, 0x70,
	0xdb, 0xe2, 0xc0, 0xdf, 0x1f, 0x38, 0x02, 0xef, 0xdb, 0x5f, 0x8a, 0x47, 0xe7, 0x6b, 0xd3, 0x3e,
	0x42, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd8, 0xff, 0xa8, 0x95, 0x04, 0x00, 0x00,
}