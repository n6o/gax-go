package gax

import (
	"encoding/json"
	"errors"
	"io"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
)

var (
	arrayOpen     = json.Delim('[')
	arrayClose    = json.Delim(']')
	errBadOpening = errors.New("unexpected opening token, expected '['")
)

type ProtoJSONStream struct {
	first, closed bool
	reader        io.ReadCloser
	stream        *json.Decoder
	typ           protoreflect.MessageType
}

func gologoo__NewProtoJSONStreamReader_3fc19dfe32fbbaa0999e790fccb054d0(rc io.ReadCloser, typ protoreflect.MessageType) *ProtoJSONStream {
	return &ProtoJSONStream{first: true, reader: rc, stream: json.NewDecoder(rc), typ: typ}
}
func (s *ProtoJSONStream) gologoo__Recv_3fc19dfe32fbbaa0999e790fccb054d0() (proto.Message, error) {
	if s.closed {
		return nil, io.EOF
	}
	if s.first {
		s.first = false
		if t, err := s.stream.Token(); err != nil {
			return nil, err
		} else if t != arrayOpen {
			return nil, errBadOpening
		}
	}
	var raw json.RawMessage
	if err := s.stream.Decode(&raw); err != nil {
		e := err
		if t, _ := s.stream.Token(); t == arrayClose {
			e = io.EOF
		}
		return nil, e
	}
	m := s.typ.New().Interface()
	err := protojson.Unmarshal(raw, m)
	return m, err
}
func (s *ProtoJSONStream) gologoo__Close_3fc19dfe32fbbaa0999e790fccb054d0() error {
	s.stream = nil
	s.closed = true
	return s.reader.Close()
}
func NewProtoJSONStreamReader(rc io.ReadCloser, typ protoreflect.MessageType) *ProtoJSONStream {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__NewProtoJSONStreamReader_3fc19dfe32fbbaa0999e790fccb054d0")
	log.Printf("Input : %v %v\n", rc, typ)
	r0 := gologoo__NewProtoJSONStreamReader_3fc19dfe32fbbaa0999e790fccb054d0(rc, typ)
	log.Printf("Output: %v\n", r0)
	return r0
}
func (s *ProtoJSONStream) Recv() (proto.Message, error) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Recv_3fc19dfe32fbbaa0999e790fccb054d0")
	log.Printf("Input : (none)\n")
	r0, r1 := s.gologoo__Recv_3fc19dfe32fbbaa0999e790fccb054d0()
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func (s *ProtoJSONStream) Close() error {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Close_3fc19dfe32fbbaa0999e790fccb054d0")
	log.Printf("Input : (none)\n")
	r0 := s.gologoo__Close_3fc19dfe32fbbaa0999e790fccb054d0()
	log.Printf("Output: %v\n", r0)
	return r0
}
