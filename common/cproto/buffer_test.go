package gproto

import (
	"testing"
)

type demoStructV2 struct {
	v21 int8
	v22 []byte
	v23 uint64
	v24 string
}

func (sv2 *demoStructV2) Encode(_ uint16) ([]byte, error) {
	// New Buffer
	gBuf := NewEmptyGBuffer()

	/////////////
	// Encode  //
	/////////////
	var err error

	// demoStructV2.v21
	if err = gBuf.WriteInt8(sv2.v21); err != nil {
		return nil, err
	}

	// demoStructV2.v22
	if err = gBuf.WriteBytes(sv2.v22); err != nil {
		return nil, err
	}

	// demoStructV2.v23
	if err = gBuf.WriteUInt64(sv2.v23); err != nil {
		return nil, err
	}

	// demoStructV2.v24
	if err = gBuf.WriteString(sv2.v24); err != nil {
		return nil, err
	}

	return gBuf.Bytes(), nil
}

func (sv2 *demoStructV2) Decode(_ uint16, buf []byte) error {
	// New Buffer
	gBuf := NewGBuffer(buf)

	/////////////
	// Decode  //
	/////////////
	var err error

	// demoStructV2.v21
	if err = gBuf.ReadInt8(&sv2.v21); err != nil {
		return err
	}

	// demoStructV2.v22
	if sv2.v22, err = gBuf.ReadBytes(); err != nil {
		return err
	}

	// demoStructV2.v23
	if err = gBuf.ReadUInt64(&sv2.v23); err != nil {
		return err
	}

	// demoStructV2.v24
	if sv2.v24, err = gBuf.ReadString(); err != nil {
		return err
	}

	return nil
}

func (sv2 *demoStructV2) Equal(other *demoStructV2) bool {

	// demoStructV2.v21
	if sv2.v21 != other.v21 {
		return false
	}

	// demoStructV2.v22
	// compare bytes
	if sv2.v22 != nil {
		if other.v22 == nil {
			return false
		}

		if len(sv2.v22) != len(other.v22) {
			return false
		}

		for i := 0; i < len(sv2.v22); i++ {
			if sv2.v22[i] != other.v22[i] {
				return false
			}
		}
	}

	// demoStructV2.v23
	if sv2.v23 != other.v23 {
		return false
	}

	// demoStructV2.v24
	if sv2.v24 != other.v24 {
		return false
	}

	return true
}

func TestBufferBasic(t *testing.T) {
	v2_1 := &demoStructV2{}
	v2_1.v21 = -12
	v2_1.v22 = []byte("ni hao你好这是Unicode")
	v2_1.v23 = 1233
	v2_1.v24 = "aabbcc你好哈哈aaa"

	bufV2_1, _ := v2_1.Encode(1)

	v2_2 := &demoStructV2{}
	v2_2.Decode(1, bufV2_1)

	if !v2_1.Equal(v2_2) {
		t.Fatal("TestBufferBasic Error")
	}
}

type demoStructV1 struct {
	v1  int8
	v2  int16
	v3  int32
	v4  int64
	v5  string
	v6  []byte
	v7  uint8
	v8  *demoStructV2
	v9  uint16
	v10 []*demoStructV2
	v11 uint32
	v12 uint64
}

func (sv1 *demoStructV1) Encode(_ uint16) ([]byte, error) {
	// New Buffer
	gBuf := NewEmptyGBuffer()

	/////////////
	// Encode  //
	/////////////
	var err error

	// demoStructV1.v1
	if err = gBuf.WriteInt8(sv1.v1); err != nil {
		return nil, err
	}

	// demoStructV1.v2
	if err = gBuf.WriteInt16(sv1.v2); err != nil {
		return nil, err
	}

	// demoStructV1.v3
	if err = gBuf.WriteInt32(sv1.v3); err != nil {
		return nil, err
	}

	// demoStructV1.v4
	if err = gBuf.WriteInt64(sv1.v4); err != nil {
		return nil, err
	}

	// demoStructV1.v5
	if err = gBuf.WriteString(sv1.v5); err != nil {
		return nil, err
	}

	// demoStructV1.v5
	if err = gBuf.WriteBytes(sv1.v6); err != nil {
		return nil, err
	}

	// demoStructV1.v7
	if err = gBuf.WriteUInt8(sv1.v7); err != nil {
		return nil, err
	}

	// demoStructV1.v8
	if err = gBuf.WriteStruct(1, sv1.v8); err != nil {
		return nil, err
	}

	// demoStructV1.v9
	if err = gBuf.WriteUInt16(sv1.v9); err != nil {
		return nil, err
	}

	// demoStructV1.v10
	if err = gBuf.WriteArray(1, sv1.v10); err != nil {
		return nil, err
	}

	// demoStructV1.v11
	if err = gBuf.WriteUInt32(sv1.v11); err != nil {
		return nil, err
	}

	// demoStructV1.v12
	if err = gBuf.WriteUInt64(sv1.v12); err != nil {
		return nil, err
	}

	return gBuf.Bytes(), nil
}

func (sv1 *demoStructV1) Decode(_ uint16, buf []byte) error {
	// New Buffer
	gBuf := NewGBuffer(buf)

	/////////////
	// Decode  //
	/////////////
	var err error

	// demoStructV1.v1
	if err = gBuf.ReadInt8(&sv1.v1); err != nil {
		return err
	}

	// demoStructV1.v2
	if err = gBuf.ReadInt16(&sv1.v2); err != nil {
		return err
	}

	// demoStructV1.v3
	if err = gBuf.ReadInt32(&sv1.v3); err != nil {
		return err
	}

	// demoStructV1.v4
	if err = gBuf.ReadInt64(&sv1.v4); err != nil {
		return err
	}

	// demoStructV1.v5
	if sv1.v5, err = gBuf.ReadString(); err != nil {
		return err
	}

	// demoStructV1.v6
	if sv1.v6, err = gBuf.ReadBytes(); err != nil {
		return err
	}

	// demoStructV1.v7
	if err = gBuf.ReadUInt8(&sv1.v7); err != nil {
		return err
	}

	// demoStructV1.v8
	sv1.v8 = &demoStructV2{}
	if err = gBuf.ReadStruct(1, sv1.v8); err != nil {
		return err
	}

	// demoStructV1.v9
	if err = gBuf.ReadUInt16(&sv1.v9); err != nil {
		return err
	}

	// demoStructV1.v10
	var v10Num uint16
	if err = gBuf.ReadUInt16(&v10Num); err != nil {
		return err
	} else {
		sv1.v10 = make([]*demoStructV2, v10Num, v10Num)
		for i := 0; i < int(v10Num); i++ {
			var tmpLen uint16
			if err = gBuf.ReadUInt16(&tmpLen); err != nil {
				return err
			} else {
				if buf1, err := gBuf.ReadRawBytes(tmpLen); err != nil {
					return err
				} else {
					sv1.v10[i] = &demoStructV2{}
					if err = sv1.v10[i].Decode(1, buf1); err != nil {
						return err
					}
				}
			}
		}
	}

	// demoStructV1.v11
	if err = gBuf.ReadUInt32(&sv1.v11); err != nil {
		return err
	}

	// demoStructV1.v12
	if err = gBuf.ReadUInt64(&sv1.v12); err != nil {
		return err
	}

	return nil
}

func (sv1 *demoStructV1) Equal(other *demoStructV1) bool {

	// demoStructV1.v1
	if sv1.v1 != other.v1 {
		return false
	}

	// demoStructV1.v2
	if sv1.v2 != other.v2 {
		return false
	}

	// demoStructV1.v3
	if sv1.v3 != other.v3 {
		return false
	}

	// demoStructV1.v4
	if sv1.v4 != other.v4 {
		return false
	}

	// demoStructV1.v5
	if sv1.v5 != other.v5 {
		return false
	}

	// demoStructV1.v6
	// compare byte
	if sv1.v6 != nil {
		if other.v6 == nil {
			return false
		}

		if len(sv1.v6) != len(other.v6) {
			return false
		}

		for i := 0; i < len(sv1.v6); i++ {
			if sv1.v6[i] != other.v6[i] {
				return false
			}
		}
	}

	// demoStructV1.v7
	if sv1.v7 != other.v7 {
		return false
	}

	// demoStructV1.v8
	if !sv1.v8.Equal(other.v8) {
		return false
	}

	// demoStructV1.v9
	if sv1.v9 != other.v9 {
		return false
	}

	// demoStructV1.v10
	if len(sv1.v10) != len(other.v10) {
		return false
	} else {
		for i := 0; i < len(sv1.v10); i++ {
			if !sv1.v10[i].Equal(other.v10[i]) {
				return false
			}
		}
	}

	// demoStructV1.v11
	if sv1.v11 != other.v11 {
		return false
	}

	// demoStructV1.v12
	if sv1.v12 != other.v12 {
		return false
	}

	return true
}

func TestBufferAdvance(t *testing.T) {
	v21 := &demoStructV2{
		v21: -11,
		v22: []byte("aaa这里是字节啊，测试努力测试....啊啊啊AABB"),
		v23: 11,
		v24: "aaaa大的我们是世界最强的团队，哈哈AAABBB",
	}
	v1 := &demoStructV1{
		v1: 1,
		v2: 11,
		v3: 222,
		v4: 1111,
		v5: "aaa",
		v6: []byte("都来这里测试，非常快的编解码Encode & Decode啊啊啊"),
		v7: 1,
		v8: v21,
		v9: 111,
		v10: []*demoStructV2{
			&demoStructV2{111, []byte("啊啊啊哈哈哈aaabbb长传冲吊"), 1111, "asdfasdfd大调"},
			&demoStructV2{111, []byte("啊啊啊bbbd大调"), 2341, "大发生的发生地方ddddbbb"}},
		v11: 111,
		v12: 222,
	}

	buf_1, _ := v1.Encode(1)

	v2 := &demoStructV1{}
	v2.Decode(1, buf_1)

	if !v1.Equal(v2) {
		t.Fatal("TestBufferAdvance Error")
	}
}

func BenchmarkEncodeAndDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v21 := &demoStructV2{
			v21: -11,
			v22: []byte("aaa这里是字节啊，测试努力测试....啊啊啊AABB"),
			v23: 11,
			v24: "aaaa大的我们是世界最强的团队，哈哈AAABBB",
		}
		v1 := &demoStructV1{
			v1: 1,
			v2: 11,
			v3: 222,
			v4: 1111,
			v5: "aaa",
			v6: []byte("都来这里测试，非常快的编解码Encode & Decode啊啊啊"),
			v7: 1,
			v8: v21,
			v9: 111,
			v10: []*demoStructV2{
				&demoStructV2{111, []byte("啊啊啊哈哈哈aaabbb长传冲吊"), 1111, "asdfasdfd大调"},
				&demoStructV2{111, []byte("啊啊啊bbbd大调"), 2341, "大发生的发生地方ddddbbb"}},
			v11: 111,
			v12: 222,
		}

		buf_1, _ := v1.Encode(1)

		v2 := &demoStructV1{}
		v2.Decode(1, buf_1)

		if !v1.Equal(v2) {
			b.Fatal("TestBufferAdvance Error")
		}
	}
}
