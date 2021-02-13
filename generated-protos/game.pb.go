// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/game.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Format int32

const (
	Format_PS4             Format = 0
	Format_XBOX            Format = 1
	Format_NINTENDO_SWITCH Format = 2
	Format_PC              Format = 3
)

var Format_name = map[int32]string{
	0: "PS4",
	1: "XBOX",
	2: "NINTENDO_SWITCH",
	3: "PC",
}

var Format_value = map[string]int32{
	"PS4":             0,
	"XBOX":            1,
	"NINTENDO_SWITCH": 2,
	"PC":              3,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}

func (Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{0}
}

type Genre int32

const (
	Genre_ACTION    Genre = 0
	Genre_ADVENTURE Genre = 1
	Genre_SPORT     Genre = 3
	Genre_RPG       Genre = 4
)

var Genre_name = map[int32]string{
	0: "ACTION",
	1: "ADVENTURE",
	3: "SPORT",
	4: "RPG",
}

var Genre_value = map[string]int32{
	"ACTION":    0,
	"ADVENTURE": 1,
	"SPORT":     3,
	"RPG":       4,
}

func (x Genre) String() string {
	return proto.EnumName(Genre_name, int32(x))
}

func (Genre) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{1}
}

type Game_NoOfPlayers int32

const (
	Game_ONE_PLAYER Game_NoOfPlayers = 0
	Game_TWO_PLAYER Game_NoOfPlayers = 1
)

var Game_NoOfPlayers_name = map[int32]string{
	0: "ONE_PLAYER",
	1: "TWO_PLAYER",
}

var Game_NoOfPlayers_value = map[string]int32{
	"ONE_PLAYER": 0,
	"TWO_PLAYER": 1,
}

func (x Game_NoOfPlayers) String() string {
	return proto.EnumName(Game_NoOfPlayers_name, int32(x))
}

func (Game_NoOfPlayers) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{0, 0}
}

type Game struct {
	Title                string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Format               Format           `protobuf:"varint,2,opt,name=format,proto3,enum=protos.Format" json:"format,omitempty"`
	Price                float32          `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Genre                Genre            `protobuf:"varint,4,opt,name=genre,proto3,enum=protos.Genre" json:"genre,omitempty"`
	Players              Game_NoOfPlayers `protobuf:"varint,5,opt,name=players,proto3,enum=protos.Game_NoOfPlayers" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{0}
}

func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Game) GetFormat() Format {
	if m != nil {
		return m.Format
	}
	return Format_PS4
}

func (m *Game) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Game) GetGenre() Genre {
	if m != nil {
		return m.Genre
	}
	return Genre_ACTION
}

func (m *Game) GetPlayers() Game_NoOfPlayers {
	if m != nil {
		return m.Players
	}
	return Game_ONE_PLAYER
}

type InStockRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InStockRequest) Reset()         { *m = InStockRequest{} }
func (m *InStockRequest) String() string { return proto.CompactTextString(m) }
func (*InStockRequest) ProtoMessage()    {}
func (*InStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{1}
}

func (m *InStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InStockRequest.Unmarshal(m, b)
}
func (m *InStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InStockRequest.Marshal(b, m, deterministic)
}
func (m *InStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InStockRequest.Merge(m, src)
}
func (m *InStockRequest) XXX_Size() int {
	return xxx_messageInfo_InStockRequest.Size(m)
}
func (m *InStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InStockRequest proto.InternalMessageInfo

func (m *InStockRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type InStockResponse struct {
	Stock                bool     `protobuf:"varint,1,opt,name=stock,proto3" json:"stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InStockResponse) Reset()         { *m = InStockResponse{} }
func (m *InStockResponse) String() string { return proto.CompactTextString(m) }
func (*InStockResponse) ProtoMessage()    {}
func (*InStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{2}
}

func (m *InStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InStockResponse.Unmarshal(m, b)
}
func (m *InStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InStockResponse.Marshal(b, m, deterministic)
}
func (m *InStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InStockResponse.Merge(m, src)
}
func (m *InStockResponse) XXX_Size() int {
	return xxx_messageInfo_InStockResponse.Size(m)
}
func (m *InStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InStockResponse proto.InternalMessageInfo

func (m *InStockResponse) GetStock() bool {
	if m != nil {
		return m.Stock
	}
	return false
}

type SearchRequest struct {
	Format               Format   `protobuf:"varint,1,opt,name=format,proto3,enum=protos.Format" json:"format,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{3}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetFormat() Format {
	if m != nil {
		return m.Format
	}
	return Format_PS4
}

func (m *SearchRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type SearchResponse struct {
	Games                []*Game  `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3658bb65e6a43a, []int{4}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetGames() []*Game {
	if m != nil {
		return m.Games
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.Format", Format_name, Format_value)
	proto.RegisterEnum("protos.Genre", Genre_name, Genre_value)
	proto.RegisterEnum("protos.Game_NoOfPlayers", Game_NoOfPlayers_name, Game_NoOfPlayers_value)
	proto.RegisterType((*Game)(nil), "protos.Game")
	proto.RegisterType((*InStockRequest)(nil), "protos.InStockRequest")
	proto.RegisterType((*InStockResponse)(nil), "protos.InStockResponse")
	proto.RegisterType((*SearchRequest)(nil), "protos.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "protos.SearchResponse")
}

func init() { proto.RegisterFile("protos/game.proto", fileDescriptor_2a3658bb65e6a43a) }

var fileDescriptor_2a3658bb65e6a43a = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x6f, 0x94, 0x40,
	0x14, 0xdd, 0xe1, 0x6b, 0xbb, 0xb7, 0x2e, 0x1d, 0xaf, 0x5a, 0x49, 0x9f, 0x08, 0x26, 0x95, 0x34,
	0xb1, 0x26, 0xd8, 0xa8, 0xf1, 0x6d, 0xa5, 0xb8, 0x92, 0x28, 0x90, 0x01, 0x6d, 0x7d, 0x6a, 0x70,
	0x33, 0x5d, 0x1b, 0x77, 0x17, 0x64, 0xd0, 0xc4, 0x9f, 0xe0, 0xdf, 0xf4, 0x97, 0x98, 0x61, 0xa0,
	0x6e, 0xd5, 0x07, 0xdf, 0x38, 0xe7, 0x9e, 0x73, 0xb9, 0xf7, 0xdc, 0x81, 0xdb, 0x75, 0x53, 0xb5,
	0x95, 0x78, 0xbc, 0x2c, 0xd7, 0xfc, 0xb8, 0xfb, 0x46, 0x4b, 0x51, 0xde, 0x4f, 0x02, 0xc6, 0xbc,
	0x5c, 0x73, 0xbc, 0x0b, 0x66, 0x7b, 0xd5, 0xae, 0xb8, 0x43, 0x5c, 0xe2, 0x4f, 0x98, 0x02, 0x78,
	0x08, 0xd6, 0x65, 0xd5, 0xac, 0xcb, 0xd6, 0xd1, 0x5c, 0xe2, 0xdb, 0x81, 0xad, 0xec, 0xe2, 0xf8,
	0x55, 0xc7, 0xb2, 0xbe, 0x2a, 0xdd, 0x75, 0x73, 0xb5, 0xe0, 0x8e, 0xee, 0x12, 0x5f, 0x63, 0x0a,
	0xe0, 0x03, 0x30, 0x97, 0x7c, 0xd3, 0x70, 0xc7, 0xe8, 0xcc, 0xd3, 0xc1, 0x3c, 0x97, 0x24, 0x53,
	0x35, 0x0c, 0x60, 0x5c, 0xaf, 0xca, 0xef, 0xbc, 0x11, 0x8e, 0xd9, 0xc9, 0x9c, 0x6b, 0x99, 0x1c,
	0x37, 0xa9, 0xd2, 0xcb, 0x4c, 0xd5, 0xd9, 0x20, 0xf4, 0x1e, 0xc1, 0xee, 0x16, 0x8f, 0x36, 0x40,
	0x9a, 0x44, 0x17, 0xd9, 0x9b, 0xd9, 0x87, 0x88, 0xd1, 0x91, 0xc4, 0xc5, 0x59, 0x3a, 0x60, 0xe2,
	0x1d, 0x82, 0x1d, 0x6f, 0xf2, 0xb6, 0x5a, 0x7c, 0x66, 0xfc, 0xcb, 0x57, 0x2e, 0xda, 0x7f, 0x6f,
	0xeb, 0x3d, 0x84, 0xbd, 0x6b, 0x9d, 0xa8, 0xab, 0x8d, 0xe8, 0x62, 0x11, 0x92, 0xe8, 0x84, 0x3b,
	0x4c, 0x01, 0xef, 0x2d, 0x4c, 0x73, 0x5e, 0x36, 0x8b, 0x4f, 0x43, 0xbf, 0xdf, 0x39, 0x91, 0xff,
	0xcb, 0x49, 0xdb, 0xca, 0xc9, 0x3b, 0x01, 0x7b, 0x68, 0xd7, 0xff, 0xd6, 0x03, 0x53, 0x1e, 0x4b,
	0x38, 0xc4, 0xd5, 0xfd, 0xdd, 0xe0, 0xd6, 0x76, 0x24, 0x4c, 0x95, 0x8e, 0x9e, 0x83, 0xa5, 0xba,
	0xe3, 0x18, 0xf4, 0x2c, 0x3f, 0xa1, 0x23, 0xdc, 0x01, 0xe3, 0xfc, 0x65, 0x7a, 0x4e, 0x09, 0xde,
	0x81, 0xbd, 0x24, 0x4e, 0x8a, 0x28, 0x39, 0x4d, 0x2f, 0xf2, 0xb3, 0xb8, 0x08, 0x5f, 0x53, 0x0d,
	0x2d, 0xd0, 0xb2, 0x90, 0xea, 0x47, 0x4f, 0xc1, 0xec, 0x4e, 0x80, 0x00, 0xd6, 0x2c, 0x2c, 0xe2,
	0x34, 0xa1, 0x23, 0x9c, 0xc2, 0x64, 0x76, 0xfa, 0x3e, 0x4a, 0x8a, 0x77, 0x2c, 0xa2, 0x04, 0x27,
	0x60, 0xe6, 0x59, 0xca, 0x0a, 0xaa, 0xcb, 0xf6, 0x2c, 0x9b, 0x53, 0x23, 0xf8, 0x41, 0x00, 0xe5,
	0x04, 0x61, 0xd9, 0x96, 0xab, 0x6a, 0x99, 0xf3, 0xe6, 0x9b, 0x3c, 0xf3, 0x0b, 0x18, 0xf7, 0xb1,
	0xe1, 0xfe, 0x30, 0xe8, 0xcd, 0xbc, 0x0f, 0xee, 0xff, 0xc5, 0xf7, 0x8b, 0x3e, 0x03, 0x4b, 0xad,
	0x8e, 0xf7, 0x06, 0xc9, 0x8d, 0x64, 0x0f, 0xf6, 0xff, 0xa4, 0x95, 0xf1, 0xa3, 0x7a, 0xc0, 0x4f,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x00, 0x89, 0x71, 0xef, 0xdc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameCatalogServiceClient is the client API for GameCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameCatalogServiceClient interface {
	InStock(ctx context.Context, in *InStockRequest, opts ...grpc.CallOption) (*InStockResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type gameCatalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameCatalogServiceClient(cc *grpc.ClientConn) GameCatalogServiceClient {
	return &gameCatalogServiceClient{cc}
}

func (c *gameCatalogServiceClient) InStock(ctx context.Context, in *InStockRequest, opts ...grpc.CallOption) (*InStockResponse, error) {
	out := new(InStockResponse)
	err := c.cc.Invoke(ctx, "/protos.GameCatalogService/InStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameCatalogServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/protos.GameCatalogService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameCatalogServiceServer is the server API for GameCatalogService service.
type GameCatalogServiceServer interface {
	InStock(context.Context, *InStockRequest) (*InStockResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedGameCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameCatalogServiceServer struct {
}

func (*UnimplementedGameCatalogServiceServer) InStock(ctx context.Context, req *InStockRequest) (*InStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InStock not implemented")
}
func (*UnimplementedGameCatalogServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterGameCatalogServiceServer(s *grpc.Server, srv GameCatalogServiceServer) {
	s.RegisterService(&_GameCatalogService_serviceDesc, srv)
}

func _GameCatalogService_InStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameCatalogServiceServer).InStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GameCatalogService/InStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameCatalogServiceServer).InStock(ctx, req.(*InStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameCatalogService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameCatalogServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GameCatalogService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameCatalogServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameCatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GameCatalogService",
	HandlerType: (*GameCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InStock",
			Handler:    _GameCatalogService_InStock_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _GameCatalogService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/game.proto",
}
