// Code generated by protoc-gen-gogo.
// source: riak_yokozuna.proto
// DO NOT EDIT!

/*
	Package rpbc is a generated protocol buffer package.

	It is generated from these files:
		riak_yokozuna.proto

	It has these top-level messages:
		RpbYokozunaIndex
		RpbYokozunaIndexGetReq
		RpbYokozunaIndexGetResp
		RpbYokozunaIndexPutReq
		RpbYokozunaIndexDeleteReq
		RpbYokozunaSchema
		RpbYokozunaSchemaPutReq
		RpbYokozunaSchemaGetReq
		RpbYokozunaSchemaGetResp
*/
package rpbc

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "gogo.pb"
// discarding unused import rpbc1 "riak.pb"

import io1 "io"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"

import bytes1 "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type RpbYokozunaIndex struct {
	Name             []byte  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Schema           []byte  `protobuf:"bytes,2,opt,name=schema" json:"schema,omitempty"`
	NVal             *uint32 `protobuf:"varint,3,opt,name=n_val" json:"n_val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbYokozunaIndex) Reset()         { *m = RpbYokozunaIndex{} }
func (m *RpbYokozunaIndex) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndex) ProtoMessage()    {}

func (m *RpbYokozunaIndex) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaIndex) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RpbYokozunaIndex) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

// GET request - If a name is given, return matching index, else return all
type RpbYokozunaIndexGetReq struct {
	Name             []byte `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexGetReq) Reset()         { *m = RpbYokozunaIndexGetReq{} }
func (m *RpbYokozunaIndexGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaIndexGetResp struct {
	Index            []*RpbYokozunaIndex `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RpbYokozunaIndexGetResp) Reset()         { *m = RpbYokozunaIndexGetResp{} }
func (m *RpbYokozunaIndexGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetResp) ProtoMessage()    {}

func (m *RpbYokozunaIndexGetResp) GetIndex() []*RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

// PUT request - Create a new index
type RpbYokozunaIndexPutReq struct {
	Index            *RpbYokozunaIndex `protobuf:"bytes,1,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RpbYokozunaIndexPutReq) Reset()         { *m = RpbYokozunaIndexPutReq{} }
func (m *RpbYokozunaIndexPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexPutReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexPutReq) GetIndex() *RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

// DELETE request - Remove an index
type RpbYokozunaIndexDeleteReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexDeleteReq) Reset()         { *m = RpbYokozunaIndexDeleteReq{} }
func (m *RpbYokozunaIndexDeleteReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexDeleteReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexDeleteReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchema struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Content          []byte `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchema) Reset()         { *m = RpbYokozunaSchema{} }
func (m *RpbYokozunaSchema) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchema) ProtoMessage()    {}

func (m *RpbYokozunaSchema) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaSchema) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// PUT request - create or potentially update a new schema
type RpbYokozunaSchemaPutReq struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaPutReq) Reset()         { *m = RpbYokozunaSchemaPutReq{} }
func (m *RpbYokozunaSchemaPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaPutReq) ProtoMessage()    {}

func (m *RpbYokozunaSchemaPutReq) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

// GET request - Return matching schema by name
type RpbYokozunaSchemaGetReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchemaGetReq) Reset()         { *m = RpbYokozunaSchemaGetReq{} }
func (m *RpbYokozunaSchemaGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetReq) ProtoMessage()    {}

func (m *RpbYokozunaSchemaGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchemaGetResp struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaGetResp) Reset()         { *m = RpbYokozunaSchemaGetResp{} }
func (m *RpbYokozunaSchemaGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetResp) ProtoMessage()    {}

func (m *RpbYokozunaSchemaGetResp) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
}
func (m *RpbYokozunaIndex) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Name = append(m.Name, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Schema = append(m.Schema, data[index:postIndex]...)
			index = postIndex
		case 3:
			if wireType != 0 {
				return ErrWrongType
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NVal = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaIndexGetReq) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Name = append(m.Name, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaIndexGetResp) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Index = append(m.Index, &RpbYokozunaIndex{})
			m.Index[len(m.Index)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaIndexPutReq) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.Index == nil {
				m.Index = &RpbYokozunaIndex{}
			}
			if err := m.Index.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaIndexDeleteReq) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Name = append(m.Name, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaSchema) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Name = append(m.Name, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaSchemaPutReq) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.Schema == nil {
				m.Schema = &RpbYokozunaSchema{}
			}
			if err := m.Schema.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaSchemaGetReq) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Name = append(m.Name, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaSchemaGetResp) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.Schema == nil {
				m.Schema = &RpbYokozunaSchema{}
			}
			if err := m.Schema.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *RpbYokozunaIndex) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaIndex) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Schema != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Schema)))
		i += copy(data[i:], m.Schema)
	}
	if m.NVal != nil {
		data[i] = 0x18
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(*m.NVal))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaIndexGetReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaIndexGetReq) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaIndexGetResp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaIndexGetResp) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		for _, msg := range m.Index {
			data[i] = 0xa
			i++
			i = encodeVarintRiakYokozuna(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaIndexPutReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaIndexPutReq) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(m.Index.Size()))
		n1, err := m.Index.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaIndexDeleteReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaIndexDeleteReq) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaSchema) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaSchema) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Content != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Content)))
		i += copy(data[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaSchemaPutReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaSchemaPutReq) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Schema != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(m.Schema.Size()))
		n2, err := m.Schema.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaSchemaGetReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaSchemaGetReq) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *RpbYokozunaSchemaGetResp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RpbYokozunaSchemaGetResp) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Schema != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRiakYokozuna(data, i, uint64(m.Schema.Size()))
		n3, err := m.Schema.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64RiakYokozuna(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RiakYokozuna(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRiakYokozuna(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RpbYokozunaIndex) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(m.Name)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.Schema != nil {
		l = len(m.Schema)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.NVal != nil {
		n += 1 + sovRiakYokozuna(uint64(*m.NVal))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaIndexGetReq) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(m.Name)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaIndexGetResp) Size() (n int) {
	var l int
	_ = l
	if len(m.Index) > 0 {
		for _, e := range m.Index {
			l = e.Size()
			n += 1 + l + sovRiakYokozuna(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaIndexPutReq) Size() (n int) {
	var l int
	_ = l
	if m.Index != nil {
		l = m.Index.Size()
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaIndexDeleteReq) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(m.Name)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaSchema) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(m.Name)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.Content != nil {
		l = len(m.Content)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaSchemaPutReq) Size() (n int) {
	var l int
	_ = l
	if m.Schema != nil {
		l = m.Schema.Size()
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaSchemaGetReq) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(m.Name)
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *RpbYokozunaSchemaGetResp) Size() (n int) {
	var l int
	_ = l
	if m.Schema != nil {
		l = m.Schema.Size()
		n += 1 + l + sovRiakYokozuna(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRiakYokozuna(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRiakYokozuna(x uint64) (n int) {
	return sovRiakYokozuna(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedRpbYokozunaIndex(r randyRiakYokozuna, easy bool) *RpbYokozunaIndex {
	this := &RpbYokozunaIndex{}
	v1 := r.Intn(100)
	this.Name = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Name[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(100)
		this.Schema = make([]byte, v2)
		for i := 0; i < v2; i++ {
			this.Schema[i] = byte(r.Intn(256))
		}
	}
	if r.Intn(10) != 0 {
		v3 := r.Uint32()
		this.NVal = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 4)
	}
	return this
}

func NewPopulatedRpbYokozunaIndexGetReq(r randyRiakYokozuna, easy bool) *RpbYokozunaIndexGetReq {
	this := &RpbYokozunaIndexGetReq{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(100)
		this.Name = make([]byte, v4)
		for i := 0; i < v4; i++ {
			this.Name[i] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaIndexGetResp(r randyRiakYokozuna, easy bool) *RpbYokozunaIndexGetResp {
	this := &RpbYokozunaIndexGetResp{}
	if r.Intn(10) != 0 {
		v5 := r.Intn(10)
		this.Index = make([]*RpbYokozunaIndex, v5)
		for i := 0; i < v5; i++ {
			this.Index[i] = NewPopulatedRpbYokozunaIndex(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaIndexPutReq(r randyRiakYokozuna, easy bool) *RpbYokozunaIndexPutReq {
	this := &RpbYokozunaIndexPutReq{}
	this.Index = NewPopulatedRpbYokozunaIndex(r, easy)
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaIndexDeleteReq(r randyRiakYokozuna, easy bool) *RpbYokozunaIndexDeleteReq {
	this := &RpbYokozunaIndexDeleteReq{}
	v6 := r.Intn(100)
	this.Name = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.Name[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaSchema(r randyRiakYokozuna, easy bool) *RpbYokozunaSchema {
	this := &RpbYokozunaSchema{}
	v7 := r.Intn(100)
	this.Name = make([]byte, v7)
	for i := 0; i < v7; i++ {
		this.Name[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v8 := r.Intn(100)
		this.Content = make([]byte, v8)
		for i := 0; i < v8; i++ {
			this.Content[i] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 3)
	}
	return this
}

func NewPopulatedRpbYokozunaSchemaPutReq(r randyRiakYokozuna, easy bool) *RpbYokozunaSchemaPutReq {
	this := &RpbYokozunaSchemaPutReq{}
	this.Schema = NewPopulatedRpbYokozunaSchema(r, easy)
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaSchemaGetReq(r randyRiakYokozuna, easy bool) *RpbYokozunaSchemaGetReq {
	this := &RpbYokozunaSchemaGetReq{}
	v9 := r.Intn(100)
	this.Name = make([]byte, v9)
	for i := 0; i < v9; i++ {
		this.Name[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

func NewPopulatedRpbYokozunaSchemaGetResp(r randyRiakYokozuna, easy bool) *RpbYokozunaSchemaGetResp {
	this := &RpbYokozunaSchemaGetResp{}
	this.Schema = NewPopulatedRpbYokozunaSchema(r, easy)
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRiakYokozuna(r, 2)
	}
	return this
}

type randyRiakYokozuna interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRiakYokozuna(r randyRiakYokozuna) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringRiakYokozuna(r randyRiakYokozuna) string {
	v10 := r.Intn(100)
	tmps := make([]rune, v10)
	for i := 0; i < v10; i++ {
		tmps[i] = randUTF8RuneRiakYokozuna(r)
	}
	return string(tmps)
}
func randUnrecognizedRiakYokozuna(r randyRiakYokozuna, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldRiakYokozuna(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldRiakYokozuna(data []byte, r randyRiakYokozuna, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateRiakYokozuna(data, uint64(key))
		v11 := r.Int63()
		if r.Intn(2) == 0 {
			v11 *= -1
		}
		data = encodeVarintPopulateRiakYokozuna(data, uint64(v11))
	case 1:
		data = encodeVarintPopulateRiakYokozuna(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateRiakYokozuna(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateRiakYokozuna(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateRiakYokozuna(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateRiakYokozuna(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (this *RpbYokozunaIndex) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaIndex)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes1.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes1.Equal(this.Schema, that1.Schema) {
		return false
	}
	if this.NVal != nil && that1.NVal != nil {
		if *this.NVal != *that1.NVal {
			return false
		}
	} else if this.NVal != nil {
		return false
	} else if that1.NVal != nil {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaIndexGetReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaIndexGetReq)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes1.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaIndexGetResp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaIndexGetResp)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Index) != len(that1.Index) {
		return false
	}
	for i := range this.Index {
		if !this.Index[i].Equal(that1.Index[i]) {
			return false
		}
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaIndexPutReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaIndexPutReq)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Index.Equal(that1.Index) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaIndexDeleteReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaIndexDeleteReq)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes1.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaSchema) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaSchema)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes1.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes1.Equal(this.Content, that1.Content) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaSchemaPutReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaSchemaPutReq)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Schema.Equal(that1.Schema) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaSchemaGetReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaSchemaGetReq)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes1.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpbYokozunaSchemaGetResp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RpbYokozunaSchemaGetResp)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Schema.Equal(that1.Schema) {
		return false
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
