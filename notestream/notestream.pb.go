// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notestream.proto

package notestream

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Note struct {
	Sender               string   `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea8159b2c6cc18e, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Note) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "notestream.Note")
}

func init() { proto.RegisterFile("notestream.proto", fileDescriptor_5ea8159b2c6cc18e) }

var fileDescriptor_5ea8159b2c6cc18e = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcb, 0x2f, 0x49,
	0x2d, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x59, 0x70, 0xb1, 0xf8, 0xe5, 0x97, 0xa4, 0x0a, 0x89, 0x71, 0xb1, 0x05, 0xa7, 0xe6, 0xa5,
	0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xbe,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x09, 0x18, 0xd7, 0xc8, 0x99, 0x8b, 0x0b,
	0xa4, 0x33, 0x18, 0x6c, 0x8e, 0x90, 0x29, 0x17, 0x37, 0x84, 0x05, 0x12, 0x2b, 0x16, 0x12, 0xd0,
	0x43, 0xb2, 0x15, 0x24, 0x24, 0x85, 0x21, 0xa2, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0x91,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe2, 0x43, 0x15, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoteStreamClient is the client API for notestream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoteStreamClient interface {
	// unary call
	StreamNotes(ctx context.Context, opts ...grpc.CallOption) (NoteStream_StreamNotesClient, error)
}

type noteStreamClient struct {
	cc *grpc.ClientConn
}

func NewNoteStreamClient(cc *grpc.ClientConn) NoteStreamClient {
	return &noteStreamClient{cc}
}

func (c *noteStreamClient) StreamNotes(ctx context.Context, opts ...grpc.CallOption) (NoteStream_StreamNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NoteStream_serviceDesc.Streams[0], "/notestream.notestream/StreamNotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &noteStreamStreamNotesClient{stream}
	return x, nil
}

type NoteStream_StreamNotesClient interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ClientStream
}

type noteStreamStreamNotesClient struct {
	grpc.ClientStream
}

func (x *noteStreamStreamNotesClient) Send(m *Note) error {
	return x.ClientStream.SendMsg(m)
}

func (x *noteStreamStreamNotesClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NoteStreamServer is the server API for notestream service.
type NoteStreamServer interface {
	// unary call
	StreamNotes(NoteStream_StreamNotesServer) error
}

func RegisterNoteStreamServer(s *grpc.Server, srv NoteStreamServer) {
	s.RegisterService(&_NoteStream_serviceDesc, srv)
}

func _NoteStream_StreamNotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NoteStreamServer).StreamNotes(&noteStreamStreamNotesServer{stream})
}

type NoteStream_StreamNotesServer interface {
	Send(*Note) error
	Recv() (*Note, error)
	grpc.ServerStream
}

type noteStreamStreamNotesServer struct {
	grpc.ServerStream
}

func (x *noteStreamStreamNotesServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

func (x *noteStreamStreamNotesServer) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NoteStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notestream.notestream",
	HandlerType: (*NoteStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNotes",
			Handler:       _NoteStream_StreamNotes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "notestream.proto",
}