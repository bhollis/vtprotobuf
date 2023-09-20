// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: pool/pool_with_oneof.proto

package pool

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *OneofTest_Test1) CloneVT() *OneofTest_Test1 {
	if m == nil {
		return (*OneofTest_Test1)(nil)
	}
	r := OneofTest_Test1FromVTPool()
	r.A = m.A
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OneofTest_Test1) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *OneofTest_Test2) CloneVT() *OneofTest_Test2 {
	if m == nil {
		return (*OneofTest_Test2)(nil)
	}
	r := OneofTest_Test2FromVTPool()
	if rhs := m.B; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.B = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OneofTest_Test2) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *OneofTest_Test3_Element2) CloneVT() *OneofTest_Test3_Element2 {
	if m == nil {
		return (*OneofTest_Test3_Element2)(nil)
	}
	r := OneofTest_Test3_Element2FromVTPool()
	r.D = m.D
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OneofTest_Test3_Element2) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *OneofTest_Test3) CloneVT() *OneofTest_Test3 {
	if m == nil {
		return (*OneofTest_Test3)(nil)
	}
	r := OneofTest_Test3FromVTPool()
	r.C = m.C.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OneofTest_Test3) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *OneofTest) CloneVT() *OneofTest {
	if m == nil {
		return (*OneofTest)(nil)
	}
	r := OneofTestFromVTPool()
	if m.Test != nil {
		r.Test = m.Test.(interface{ CloneVT() isOneofTest_Test }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OneofTest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *OneofTest_Test1_) CloneVT() isOneofTest_Test {
	if m == nil {
		return (*OneofTest_Test1_)(nil)
	}
	r := new(OneofTest_Test1_)
	r.Test1 = m.Test1.CloneVT()
	return r
}

func (m *OneofTest_Test2_) CloneVT() isOneofTest_Test {
	if m == nil {
		return (*OneofTest_Test2_)(nil)
	}
	r := new(OneofTest_Test2_)
	r.Test2 = m.Test2.CloneVT()
	return r
}

func (m *OneofTest_Test3_) CloneVT() isOneofTest_Test {
	if m == nil {
		return (*OneofTest_Test3_)(nil)
	}
	r := new(OneofTest_Test3_)
	r.Test3 = m.Test3.CloneVT()
	return r
}

func (this *OneofTest_Test1) EqualVT(that *OneofTest_Test1) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.A != that.A {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OneofTest_Test1) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OneofTest_Test1)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *OneofTest_Test2) EqualVT(that *OneofTest_Test2) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.B) != len(that.B) {
		return false
	}
	for i, vx := range this.B {
		vy := that.B[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OneofTest_Test2) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OneofTest_Test2)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *OneofTest_Test3_Element2) EqualVT(that *OneofTest_Test3_Element2) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.D != that.D {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OneofTest_Test3_Element2) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OneofTest_Test3_Element2)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *OneofTest_Test3) EqualVT(that *OneofTest_Test3) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.C.EqualVT(that.C) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OneofTest_Test3) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OneofTest_Test3)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *OneofTest) EqualVT(that *OneofTest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Test == nil && that.Test != nil {
		return false
	} else if this.Test != nil {
		if that.Test == nil {
			return false
		}
		if !this.Test.(interface{ EqualVT(isOneofTest_Test) bool }).EqualVT(that.Test) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OneofTest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*OneofTest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *OneofTest_Test1_) EqualVT(thatIface isOneofTest_Test) bool {
	that, ok := thatIface.(*OneofTest_Test1_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Test1, that.Test1; p != q {
		if p == nil {
			p = &OneofTest_Test1{}
		}
		if q == nil {
			q = &OneofTest_Test1{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *OneofTest_Test2_) EqualVT(thatIface isOneofTest_Test) bool {
	that, ok := thatIface.(*OneofTest_Test2_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Test2, that.Test2; p != q {
		if p == nil {
			p = &OneofTest_Test2{}
		}
		if q == nil {
			q = &OneofTest_Test2{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *OneofTest_Test3_) EqualVT(thatIface isOneofTest_Test) bool {
	that, ok := thatIface.(*OneofTest_Test3_)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Test3, that.Test3; p != q {
		if p == nil {
			p = &OneofTest_Test3{}
		}
		if q == nil {
			q = &OneofTest_Test3{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (m *OneofTest_Test1) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test1) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test1) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.A != 0 {
		i = encodeVarint(dAtA, i, uint64(m.A))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test2) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test2) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test2) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.B) > 0 {
		for iNdEx := len(m.B) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.B[iNdEx])
			copy(dAtA[i:], m.B[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.B[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test3_Element2) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test3_Element2) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test3_Element2) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.D != 0 {
		i = encodeVarint(dAtA, i, uint64(m.D))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test3) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test3) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test3) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.C != nil {
		size, err := m.C.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Test.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test1_) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test1_) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test1 != nil {
		size, err := m.Test1.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *OneofTest_Test2_) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test2_) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test2 != nil {
		size, err := m.Test2.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *OneofTest_Test3_) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OneofTest_Test3_) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test3 != nil {
		size, err := m.Test3.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *OneofTest_Test1) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test1) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test1) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.A != 0 {
		i = encodeVarint(dAtA, i, uint64(m.A))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test2) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test2) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test2) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.B) > 0 {
		for iNdEx := len(m.B) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.B[iNdEx])
			copy(dAtA[i:], m.B[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.B[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test3_Element2) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test3_Element2) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test3_Element2) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.D != 0 {
		i = encodeVarint(dAtA, i, uint64(m.D))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test3) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest_Test3) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test3) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.C != nil {
		size, err := m.C.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneofTest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.Test.(*OneofTest_Test3_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Test.(*OneofTest_Test2_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Test.(*OneofTest_Test1_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *OneofTest_Test1_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test1_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test1 != nil {
		size, err := m.Test1.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *OneofTest_Test2_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test2_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test2 != nil {
		size, err := m.Test2.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *OneofTest_Test3_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OneofTest_Test3_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Test3 != nil {
		size, err := m.Test3.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

var vtprotoPool_OneofTest_Test1 = sync.Pool{
	New: func() interface{} {
		return &OneofTest_Test1{}
	},
}

func (m *OneofTest_Test1) ResetVT() {
	if m != nil {
		m.Reset()
	}
}
func (m *OneofTest_Test1) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_OneofTest_Test1.Put(m)
	}
}
func OneofTest_Test1FromVTPool() *OneofTest_Test1 {
	return vtprotoPool_OneofTest_Test1.Get().(*OneofTest_Test1)
}

var vtprotoPool_OneofTest_Test2 = sync.Pool{
	New: func() interface{} {
		return &OneofTest_Test2{}
	},
}

func (m *OneofTest_Test2) ResetVT() {
	if m != nil {
		f0 := m.B[:0]
		m.Reset()
		m.B = f0
	}
}
func (m *OneofTest_Test2) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_OneofTest_Test2.Put(m)
	}
}
func OneofTest_Test2FromVTPool() *OneofTest_Test2 {
	return vtprotoPool_OneofTest_Test2.Get().(*OneofTest_Test2)
}

var vtprotoPool_OneofTest_Test3_Element2 = sync.Pool{
	New: func() interface{} {
		return &OneofTest_Test3_Element2{}
	},
}

func (m *OneofTest_Test3_Element2) ResetVT() {
	if m != nil {
		m.Reset()
	}
}
func (m *OneofTest_Test3_Element2) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_OneofTest_Test3_Element2.Put(m)
	}
}
func OneofTest_Test3_Element2FromVTPool() *OneofTest_Test3_Element2 {
	return vtprotoPool_OneofTest_Test3_Element2.Get().(*OneofTest_Test3_Element2)
}

var vtprotoPool_OneofTest_Test3 = sync.Pool{
	New: func() interface{} {
		return &OneofTest_Test3{}
	},
}

func (m *OneofTest_Test3) ResetVT() {
	if m != nil {
		m.C.ReturnToVTPool()
		m.Reset()
	}
}
func (m *OneofTest_Test3) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_OneofTest_Test3.Put(m)
	}
}
func OneofTest_Test3FromVTPool() *OneofTest_Test3 {
	return vtprotoPool_OneofTest_Test3.Get().(*OneofTest_Test3)
}

var vtprotoPool_OneofTest = sync.Pool{
	New: func() interface{} {
		return &OneofTest{}
	},
}

func (m *OneofTest) ResetVT() {
	if m != nil {
		if oneof, ok := m.Test.(*OneofTest_Test1_); ok {
			oneof.Test1.ReturnToVTPool()
		}
		if oneof, ok := m.Test.(*OneofTest_Test2_); ok {
			oneof.Test2.ReturnToVTPool()
		}
		if oneof, ok := m.Test.(*OneofTest_Test3_); ok {
			oneof.Test3.ReturnToVTPool()
		}
		m.Reset()
	}
}
func (m *OneofTest) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_OneofTest.Put(m)
	}
}
func OneofTestFromVTPool() *OneofTest {
	return vtprotoPool_OneofTest.Get().(*OneofTest)
}
func (m *OneofTest_Test1) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.A != 0 {
		n += 1 + sov(uint64(m.A))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OneofTest_Test2) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.B) > 0 {
		for _, s := range m.B {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *OneofTest_Test3_Element2) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.D != 0 {
		n += 1 + sov(uint64(m.D))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OneofTest_Test3) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.C != nil {
		l = m.C.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OneofTest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Test.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *OneofTest_Test1_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Test1 != nil {
		l = m.Test1.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *OneofTest_Test2_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Test2 != nil {
		l = m.Test2.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *OneofTest_Test3_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Test3 != nil {
		l = m.Test3.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *OneofTest_Test1) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: OneofTest_Test1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest_Test1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OneofTest_Test2) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: OneofTest_Test2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest_Test2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = append(m.B, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OneofTest_Test3_Element2) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: OneofTest_Test3_Element2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest_Test3_Element2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			m.D = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.D |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OneofTest_Test3) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: OneofTest_Test3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest_Test3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.C == nil {
				m.C = OneofTest_Test3_Element2FromVTPool()
			}
			if err := m.C.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OneofTest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: OneofTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneofTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Test.(*OneofTest_Test1_); ok {
				if err := oneof.Test1.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := OneofTest_Test1FromVTPool()
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Test = &OneofTest_Test1_{Test1: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Test.(*OneofTest_Test2_); ok {
				if err := oneof.Test2.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := OneofTest_Test2FromVTPool()
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Test = &OneofTest_Test2_{Test2: v}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test3", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Test.(*OneofTest_Test3_); ok {
				if err := oneof.Test3.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := OneofTest_Test3FromVTPool()
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Test = &OneofTest_Test3_{Test3: v}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
