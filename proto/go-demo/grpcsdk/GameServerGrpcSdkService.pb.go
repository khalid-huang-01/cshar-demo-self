// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GameServerGrpcSdkService.proto

package grpcsdk

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// 定时拉取游戏进程健康检查
type HealthCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

type HealthCheckResponse struct {
	HealthStatus         bool     `protobuf:"varint,1,opt,name=healthStatus,proto3" json:"healthStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetHealthStatus() bool {
	if m != nil {
		return m.HealthStatus
	}
	return false
}

// 游戏属性详情
type GameProperty struct {
	// 属性名称（键）
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 属性值（值）
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameProperty) Reset()         { *m = GameProperty{} }
func (m *GameProperty) String() string { return proto.CompactTextString(m) }
func (*GameProperty) ProtoMessage()    {}
func (*GameProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{2}
}

func (m *GameProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameProperty.Unmarshal(m, b)
}
func (m *GameProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameProperty.Marshal(b, m, deterministic)
}
func (m *GameProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameProperty.Merge(m, src)
}
func (m *GameProperty) XXX_Size() int {
	return xxx_messageInfo_GameProperty.Size(m)
}
func (m *GameProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_GameProperty.DiscardUnknown(m)
}

var xxx_messageInfo_GameProperty proto.InternalMessageInfo

func (m *GameProperty) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GameProperty) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// game server session
type GameServerSession struct {
	GameServerSessionId   string          `protobuf:"bytes,1,opt,name=gameServerSessionId,proto3" json:"gameServerSessionId,omitempty"`
	FleetId               string          `protobuf:"bytes,2,opt,name=fleetId,proto3" json:"fleetId,omitempty"`
	Name                  string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	MaxPlayers            int32           `protobuf:"varint,4,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
	Joinable              bool            `protobuf:"varint,5,opt,name=joinable,proto3" json:"joinable,omitempty"`
	GameProperties        []*GameProperty `protobuf:"bytes,6,rep,name=gameProperties,proto3" json:"gameProperties,omitempty"`
	Port                  int32           `protobuf:"varint,7,opt,name=port,proto3" json:"port,omitempty"`
	IpAddress             string          `protobuf:"bytes,8,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	GameServerSessionData string          `protobuf:"bytes,9,opt,name=gameServerSessionData,proto3" json:"gameServerSessionData,omitempty"`
	MatchmakerData        string          `protobuf:"bytes,10,opt,name=matchmakerData,proto3" json:"matchmakerData,omitempty"`
	DnsName               string          `protobuf:"bytes,11,opt,name=dnsName,proto3" json:"dnsName,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *GameServerSession) Reset()         { *m = GameServerSession{} }
func (m *GameServerSession) String() string { return proto.CompactTextString(m) }
func (*GameServerSession) ProtoMessage()    {}
func (*GameServerSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{3}
}

func (m *GameServerSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServerSession.Unmarshal(m, b)
}
func (m *GameServerSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServerSession.Marshal(b, m, deterministic)
}
func (m *GameServerSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServerSession.Merge(m, src)
}
func (m *GameServerSession) XXX_Size() int {
	return xxx_messageInfo_GameServerSession.Size(m)
}
func (m *GameServerSession) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServerSession.DiscardUnknown(m)
}

var xxx_messageInfo_GameServerSession proto.InternalMessageInfo

func (m *GameServerSession) GetGameServerSessionId() string {
	if m != nil {
		return m.GameServerSessionId
	}
	return ""
}

func (m *GameServerSession) GetFleetId() string {
	if m != nil {
		return m.FleetId
	}
	return ""
}

func (m *GameServerSession) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameServerSession) GetMaxPlayers() int32 {
	if m != nil {
		return m.MaxPlayers
	}
	return 0
}

func (m *GameServerSession) GetJoinable() bool {
	if m != nil {
		return m.Joinable
	}
	return false
}

func (m *GameServerSession) GetGameProperties() []*GameProperty {
	if m != nil {
		return m.GameProperties
	}
	return nil
}

func (m *GameServerSession) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *GameServerSession) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *GameServerSession) GetGameServerSessionData() string {
	if m != nil {
		return m.GameServerSessionData
	}
	return ""
}

func (m *GameServerSession) GetMatchmakerData() string {
	if m != nil {
		return m.MatchmakerData
	}
	return ""
}

func (m *GameServerSession) GetDnsName() string {
	if m != nil {
		return m.DnsName
	}
	return ""
}

// 分配gameserversession到游戏进程
type StartGameServerSessionRequest struct {
	GameServerSession    *GameServerSession `protobuf:"bytes,1,opt,name=gameServerSession,proto3" json:"gameServerSession,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StartGameServerSessionRequest) Reset()         { *m = StartGameServerSessionRequest{} }
func (m *StartGameServerSessionRequest) String() string { return proto.CompactTextString(m) }
func (*StartGameServerSessionRequest) ProtoMessage()    {}
func (*StartGameServerSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{4}
}

func (m *StartGameServerSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameServerSessionRequest.Unmarshal(m, b)
}
func (m *StartGameServerSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameServerSessionRequest.Marshal(b, m, deterministic)
}
func (m *StartGameServerSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameServerSessionRequest.Merge(m, src)
}
func (m *StartGameServerSessionRequest) XXX_Size() int {
	return xxx_messageInfo_StartGameServerSessionRequest.Size(m)
}
func (m *StartGameServerSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameServerSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameServerSessionRequest proto.InternalMessageInfo

func (m *StartGameServerSessionRequest) GetGameServerSession() *GameServerSession {
	if m != nil {
		return m.GameServerSession
	}
	return nil
}

// 结束游戏进程
type ProcessTerminateRequest struct {
	TerminationTime      int64    `protobuf:"varint,1,opt,name=terminationTime,proto3" json:"terminationTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessTerminateRequest) Reset()         { *m = ProcessTerminateRequest{} }
func (m *ProcessTerminateRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessTerminateRequest) ProtoMessage()    {}
func (*ProcessTerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b60b42a1666dea1c, []int{5}
}

func (m *ProcessTerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessTerminateRequest.Unmarshal(m, b)
}
func (m *ProcessTerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessTerminateRequest.Marshal(b, m, deterministic)
}
func (m *ProcessTerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessTerminateRequest.Merge(m, src)
}
func (m *ProcessTerminateRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessTerminateRequest.Size(m)
}
func (m *ProcessTerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessTerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessTerminateRequest proto.InternalMessageInfo

func (m *ProcessTerminateRequest) GetTerminationTime() int64 {
	if m != nil {
		return m.TerminationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "tencentcloud.gse.grpcsdk.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "tencentcloud.gse.grpcsdk.HealthCheckResponse")
	proto.RegisterType((*GameProperty)(nil), "tencentcloud.gse.grpcsdk.GameProperty")
	proto.RegisterType((*GameServerSession)(nil), "tencentcloud.gse.grpcsdk.GameServerSession")
	proto.RegisterType((*StartGameServerSessionRequest)(nil), "tencentcloud.gse.grpcsdk.StartGameServerSessionRequest")
	proto.RegisterType((*ProcessTerminateRequest)(nil), "tencentcloud.gse.grpcsdk.ProcessTerminateRequest")
}

func init() { proto.RegisterFile("GameServerGrpcSdkService.proto", fileDescriptor_b60b42a1666dea1c) }

var fileDescriptor_b60b42a1666dea1c = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6f, 0x12, 0x41,
	0x10, 0x96, 0x52, 0x5a, 0x18, 0x6a, 0x2b, 0xd3, 0x9a, 0x6e, 0x88, 0x36, 0xe4, 0x12, 0x1b, 0x12,
	0x2d, 0x2a, 0x1a, 0x8d, 0x8f, 0x5a, 0x13, 0xec, 0x4b, 0x21, 0x47, 0x5f, 0xf4, 0x6d, 0x7b, 0x37,
	0xc2, 0xc9, 0xdd, 0xee, 0xb9, 0xbb, 0x34, 0x62, 0xfc, 0x87, 0xfc, 0x17, 0x7d, 0x32, 0xb7, 0x5c,
	0x0b, 0xdc, 0x71, 0xa4, 0x6f, 0x33, 0xdf, 0x7d, 0xf3, 0xeb, 0x9b, 0x9d, 0x83, 0x93, 0x1e, 0x8f,
	0x68, 0x48, 0xea, 0x86, 0x54, 0x4f, 0xc5, 0xde, 0xd0, 0x9f, 0x24, 0x4e, 0xe0, 0x51, 0x27, 0x56,
	0xd2, 0x48, 0x64, 0x86, 0x84, 0x47, 0xc2, 0x78, 0xa1, 0x9c, 0xfa, 0x9d, 0x91, 0xa6, 0xce, 0x48,
	0xc5, 0x9e, 0xf6, 0x27, 0xcd, 0xe3, 0x9e, 0xa6, 0x75, 0x21, 0xce, 0x11, 0xe0, 0x17, 0xe2, 0xa1,
	0x19, 0x9f, 0x8f, 0xc9, 0x9b, 0xb8, 0xf4, 0x73, 0x4a, 0xda, 0x38, 0x1f, 0xe0, 0x70, 0x05, 0xd5,
	0xb1, 0x14, 0x9a, 0xd0, 0x81, 0xbd, 0xb1, 0x85, 0x87, 0x86, 0x9b, 0xa9, 0x66, 0xa5, 0x56, 0xa9,
	0x5d, 0x75, 0x57, 0x30, 0xe7, 0x1d, 0xec, 0x25, 0x5d, 0x0e, 0x94, 0x8c, 0x49, 0x99, 0x19, 0x3e,
	0x82, 0xf2, 0x84, 0x66, 0x96, 0x5a, 0x73, 0x13, 0x13, 0x8f, 0xa0, 0x72, 0xc3, 0xc3, 0x29, 0xb1,
	0x2d, 0x8b, 0xcd, 0x1d, 0xe7, 0x6f, 0x19, 0x1a, 0x8b, 0xf1, 0x86, 0xa4, 0x75, 0x20, 0x05, 0xbe,
	0x82, 0xc3, 0x51, 0x16, 0xbc, 0xf0, 0xd3, 0x6c, 0xeb, 0x3e, 0x21, 0x83, 0xdd, 0xef, 0x21, 0x91,
	0xb9, 0xf0, 0xd3, 0xfc, 0xb7, 0x2e, 0x22, 0x6c, 0x0b, 0x1e, 0x11, 0x2b, 0x5b, 0xd8, 0xda, 0x78,
	0x02, 0x10, 0xf1, 0x5f, 0x83, 0x90, 0xcf, 0x48, 0x69, 0xb6, 0xdd, 0x2a, 0xb5, 0x2b, 0xee, 0x12,
	0x82, 0x4d, 0xa8, 0xfe, 0x90, 0x81, 0xe0, 0xd7, 0x21, 0xb1, 0x8a, 0x9d, 0xf6, 0xce, 0xc7, 0x4b,
	0xd8, 0x1f, 0x2d, 0x26, 0x0d, 0x48, 0xb3, 0x9d, 0x56, 0xb9, 0x5d, 0xef, 0x9e, 0x76, 0x8a, 0xd6,
	0xd0, 0x59, 0x56, 0xc6, 0xcd, 0x44, 0x27, 0xfd, 0xc5, 0x52, 0x19, 0xb6, 0x6b, 0xbb, 0xb0, 0x36,
	0x3e, 0x81, 0x5a, 0x10, 0x7f, 0xf4, 0x7d, 0x45, 0x5a, 0xb3, 0xaa, 0x6d, 0x7c, 0x01, 0xe0, 0x5b,
	0x78, 0x9c, 0x93, 0xe0, 0x33, 0x37, 0x9c, 0xd5, 0x2c, 0x73, 0xfd, 0x47, 0x3c, 0x85, 0xfd, 0x88,
	0x1b, 0x6f, 0x1c, 0xf1, 0x09, 0x29, 0x4b, 0x07, 0x4b, 0xcf, 0xa0, 0x89, 0x92, 0xbe, 0xd0, 0x97,
	0x89, 0x64, 0xf5, 0xb9, 0x92, 0xa9, 0xeb, 0xfc, 0x86, 0xa7, 0x43, 0xc3, 0x95, 0xc9, 0xed, 0x2b,
	0x7d, 0x3f, 0xf8, 0x15, 0x1a, 0xb9, 0xda, 0x76, 0x69, 0xf5, 0xee, 0xf3, 0xcd, 0xea, 0xac, 0xa6,
	0xcb, 0x67, 0x71, 0xce, 0xe1, 0x78, 0xa0, 0xa4, 0x47, 0x5a, 0x5f, 0x91, 0x8a, 0x02, 0xc1, 0x0d,
	0xdd, 0x56, 0x6d, 0xc3, 0x81, 0x49, 0xb1, 0x40, 0x8a, 0xab, 0x20, 0x22, 0x5b, 0xb3, 0xec, 0x66,
	0xe1, 0xee, 0xbf, 0x2d, 0x60, 0x45, 0xb7, 0x84, 0x02, 0x1e, 0xf6, 0xc5, 0xd2, 0xf3, 0xc7, 0x17,
	0xc5, 0x2d, 0xe7, 0x6f, 0xa7, 0x79, 0x76, 0x4f, 0xf6, 0xfc, 0xa6, 0x9c, 0x07, 0xf8, 0x07, 0x58,
	0x5f, 0xac, 0xd7, 0x13, 0xdf, 0x17, 0x27, 0xdb, 0xb8, 0x81, 0xe6, 0xb3, 0x0d, 0x32, 0x6b, 0x5a,
	0xaa, 0x1e, 0x03, 0xf6, 0x45, 0x56, 0x51, 0x7c, 0x5d, 0x1c, 0x5e, 0xa0, 0xfe, 0xbd, 0x2b, 0x7e,
	0x6a, 0x7c, 0x3b, 0x18, 0xc9, 0x33, 0x9f, 0x22, 0xf9, 0x32, 0x25, 0x5c, 0xef, 0xd8, 0x9f, 0xd1,
	0x9b, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x96, 0x79, 0xc6, 0xe1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServerGrpcSdkServiceClient is the client API for GameServerGrpcSdkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServerGrpcSdkServiceClient interface {
	// 接收健康检查请求
	OnHealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	// 接收游戏会话
	OnStartGameServerSession(ctx context.Context, in *StartGameServerSessionRequest, opts ...grpc.CallOption) (*GseResponse, error)
	// 结束游戏进程
	OnProcessTerminate(ctx context.Context, in *ProcessTerminateRequest, opts ...grpc.CallOption) (*GseResponse, error)
}

type gameServerGrpcSdkServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServerGrpcSdkServiceClient(cc *grpc.ClientConn) GameServerGrpcSdkServiceClient {
	return &gameServerGrpcSdkServiceClient{cc}
}

func (c *gameServerGrpcSdkServiceClient) OnHealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerGrpcSdkServiceClient) OnStartGameServerSession(ctx context.Context, in *StartGameServerSessionRequest, opts ...grpc.CallOption) (*GseResponse, error) {
	out := new(GseResponse)
	err := c.cc.Invoke(ctx, "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnStartGameServerSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerGrpcSdkServiceClient) OnProcessTerminate(ctx context.Context, in *ProcessTerminateRequest, opts ...grpc.CallOption) (*GseResponse, error) {
	out := new(GseResponse)
	err := c.cc.Invoke(ctx, "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnProcessTerminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServerGrpcSdkServiceServer is the server API for GameServerGrpcSdkService service.
type GameServerGrpcSdkServiceServer interface {
	// 接收健康检查请求
	OnHealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	// 接收游戏会话
	OnStartGameServerSession(context.Context, *StartGameServerSessionRequest) (*GseResponse, error)
	// 结束游戏进程
	OnProcessTerminate(context.Context, *ProcessTerminateRequest) (*GseResponse, error)
}

func RegisterGameServerGrpcSdkServiceServer(s *grpc.Server, srv GameServerGrpcSdkServiceServer) {
	s.RegisterService(&_GameServerGrpcSdkService_serviceDesc, srv)
}

func _GameServerGrpcSdkService_OnHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerGrpcSdkServiceServer).OnHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerGrpcSdkServiceServer).OnHealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerGrpcSdkService_OnStartGameServerSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameServerSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerGrpcSdkServiceServer).OnStartGameServerSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnStartGameServerSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerGrpcSdkServiceServer).OnStartGameServerSession(ctx, req.(*StartGameServerSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerGrpcSdkService_OnProcessTerminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessTerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerGrpcSdkServiceServer).OnProcessTerminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencentcloud.gse.grpcsdk.GameServerGrpcSdkService/OnProcessTerminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerGrpcSdkServiceServer).OnProcessTerminate(ctx, req.(*ProcessTerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameServerGrpcSdkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tencentcloud.gse.grpcsdk.GameServerGrpcSdkService",
	HandlerType: (*GameServerGrpcSdkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnHealthCheck",
			Handler:    _GameServerGrpcSdkService_OnHealthCheck_Handler,
		},
		{
			MethodName: "OnStartGameServerSession",
			Handler:    _GameServerGrpcSdkService_OnStartGameServerSession_Handler,
		},
		{
			MethodName: "OnProcessTerminate",
			Handler:    _GameServerGrpcSdkService_OnProcessTerminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GameServerGrpcSdkService.proto",
}