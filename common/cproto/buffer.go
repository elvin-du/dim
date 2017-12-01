package gproto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type GBuffer struct {
	Buffer *bytes.Buffer
}

// for encoder
func NewEmptyGBuffer() *GBuffer {
	return &GBuffer{
		Buffer: new(bytes.Buffer),
	}
}

// for decoder
func NewGBuffer(buf []byte) *GBuffer {
	return &GBuffer{
		Buffer: bytes.NewBuffer(buf),
	}
}

func (gBuffer *GBuffer) Bytes() []byte {
	return gBuffer.Buffer.Bytes()
}

func (gBuffer *GBuffer) Len() int {
	return gBuffer.Buffer.Len()
}

func (gBuffer *GBuffer) WriteInt8(v int8) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

// <1> Int(Big-Endian编码)
// int8, int16, int32, int64
// uint8, uint16, uint32, uint64

func (gBuffer *GBuffer) ReadInt8(v *int8) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteInt16(v int16) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadInt16(v *int16) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteInt32(v int32) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadInt32(v *int32) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteInt64(v int64) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadInt64(v *int64) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteUInt8(v uint8) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadUInt8(v *uint8) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteUInt16(v uint16) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadUInt16(v *uint16) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteUInt32(v uint32) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadUInt32(v *uint32) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) WriteUInt64(v uint64) error {
	return binary.Write(gBuffer.Buffer, binary.BigEndian, v)
}

func (gBuffer *GBuffer) ReadUInt64(v *uint64) error {
	return binary.Read(gBuffer.Buffer, binary.BigEndian, v)
}

// <2> Bytes/String
// uint16长度 + 具体字节流
//
// 备注:
// 没有内容的时候只写入uint16长度
func (gBuffer *GBuffer) WriteBytes(b []byte) error {
	if b == nil {
		// Write Len
		if err := gBuffer.WriteUInt16(uint16(0)); err != nil {
			return fmt.Errorf("GBuffer.WriteBytes Len Error: len=%v, err=%v", len(b), err)
		}
		return nil
	}

	if len(b) == 0 {
		// Write Len
		if err := gBuffer.WriteUInt16(uint16(0)); err != nil {
			return fmt.Errorf("GBuffer.WriteBytes Len Error: len=%v, err=%v", len(b), err)
		}
		return nil
	}

	// Write Len
	if err := gBuffer.WriteUInt16(uint16(len(b))); err != nil {
		return fmt.Errorf("GBuffer.WriteBytes Len Error: len=%v, err=%v", len(b), err)
	}

	// Write bytes
	if n, _ := gBuffer.Buffer.Write(b); n == len(b) {
		// 写入成功
		return nil
	} else {
		return fmt.Errorf("GBuffer.WriteBytes Error: len=%v, n=%v", len(b), n)
	}
}

// 备注:
// 会自动分配内存
func (gBuffer *GBuffer) ReadBytes() ([]byte, error) {
	var l uint16
	if err := gBuffer.ReadUInt16(&l); err != nil {
		return nil, fmt.Errorf("GBuffer.ReadBytes Len Error: err=%v", err)
	}

	if l == 0 {
		return nil, nil
	}

	out := make([]byte, l, l)
	if n, _ := gBuffer.Buffer.Read(out); n == int(l) {
		return out, nil
	} else {
		return nil, fmt.Errorf("GBuffer.ReadBytes Error: len=%v, n=%v", l, n)
	}
}

func (gBuffer *GBuffer) WriteString(str string) error {
	return gBuffer.WriteBytes([]byte(str))
}

func (gBuffer *GBuffer) ReadString() (string, error) {
	if buf, err := gBuffer.ReadBytes(); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}

// 没有额外的长度信息
func (gBuffer *GBuffer) WriteRawBytes(b []byte) error {
	if b == nil {
		return nil
	}

	if len(b) == 0 {
		return nil
	}

	// Write bytes
	if n, _ := gBuffer.Buffer.Write(b); n == len(b) {
		return nil
	} else {
		return fmt.Errorf("GBuffer.WriteRawBytes Error: len=%v, n=%v", len(b), n)
	}
}

func (gBuffer *GBuffer) ReadRawBytes(l uint16) ([]byte, error) {
	if l == 0 {
		return nil, nil
	}

	out := make([]byte, l, l)
	if n, _ := gBuffer.Buffer.Read(out); n == int(l) {
		return out, nil
	} else {
		return nil, fmt.Errorf("GBuffer.ReadRawBytes Error: len=%v, n=%v", l, n)
	}
}

// <3> Struct虚拟的约束
// uint16长度 + 具体字节流
func (gBuffer *GBuffer) WriteStruct(version uint16, iMessage IMessage) error {
	if iMessage == nil {
		return nil
	}

	if buf, err := iMessage.Encode(version); err != nil {
		return err
	} else {
		if err = gBuffer.WriteUInt16(uint16(len(buf))); err != nil {
			return err
		}
		if err = gBuffer.WriteRawBytes(buf); err != nil {
			return err
		}
	}

	return nil
}

func (gBuffer *GBuffer) ReadStruct(version uint16, iMessage IMessage) error {
	var l uint16
	if err := gBuffer.ReadUInt16(&l); err != nil {
		return err
	} else {
		if buf, err := gBuffer.ReadRawBytes(l); err != nil {
			return err
		} else {
			if err := iMessage.Decode(version, buf); err != nil {
				return err
			}
		}
	}

	return nil
}

// <4> List
// uint16具体元素个数 + 具体元素
//
// 备注:
// arr必须是[]IMessage, 否则会Panic
func (gBuffer *GBuffer) WriteArray(version uint16, arr interface{}) error {
	if arr == nil {
		return nil
	}

	iMessages := reflect.ValueOf(arr)
	if iMessages.Kind() != reflect.Slice {
		return fmt.Errorf("GBuffer.WriteArray Arguments Error: arr=%v", arr)
	}

	l := iMessages.Len()
	if err := gBuffer.WriteUInt16(uint16(l)); err != nil {
		return err
	} else {
		// 循环 Encode iMessage
		for i := 0; i < l; i++ {
			if buf, err := (iMessages.Index(i).Interface()).(IMessage).Encode(version); err != nil {
				return err
			} else {
				if err = gBuffer.WriteUInt16(uint16(len(buf))); err != nil {
					return err
				}
				if err = gBuffer.WriteRawBytes(buf); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// 备注:
// 这只是一个Read Array的范例, 当前版本这个函数并不能使用
func (gBuffer *GBuffer) ReadArray(version uint16) ([]IMessage, error) {
	var num uint16
	if err := gBuffer.ReadUInt16(&num); err != nil {
		return nil, fmt.Errorf("GBuffer.ReadArray Len Error: err=%v", err)
	}

	if num == 0 {
		return nil, nil
	}

	iMessages := make([]IMessage, num, num)
	for i := 0; i < int(num); i++ {
		// 循环 Decode iMessage
		var l uint16
		if err := gBuffer.ReadUInt16(&l); err != nil {
			return nil, err
		} else {
			if buf, err := gBuffer.ReadRawBytes(l); err != nil {
				return nil, err
			} else {
				if err := iMessages[i].Decode(version, buf); err != nil {
					return nil, err
				}
			}
		}
	}

	return iMessages, nil
}
