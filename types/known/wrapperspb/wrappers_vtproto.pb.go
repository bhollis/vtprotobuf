// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: google/protobuf/wrappers.proto

package wrapperspb

import (
	binary "encoding/binary"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	math "math"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func DoubleValue_MarshalVT(m *wrapperspb.DoubleValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := DoubleValue_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := DoubleValue_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func DoubleValue_MarshalToVT(m *wrapperspb.DoubleValue, dAtA []byte) (int, error) {
	size := DoubleValue_SizeVT(m)
	return DoubleValue_MarshalToSizedBufferVT(m, dAtA[:size])
}

func DoubleValue_MarshalToSizedBufferVT(m *wrapperspb.DoubleValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func FloatValue_MarshalVT(m *wrapperspb.FloatValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := FloatValue_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := FloatValue_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func FloatValue_MarshalToVT(m *wrapperspb.FloatValue, dAtA []byte) (int, error) {
	size := FloatValue_SizeVT(m)
	return FloatValue_MarshalToSizedBufferVT(m, dAtA[:size])
}

func FloatValue_MarshalToSizedBufferVT(m *wrapperspb.FloatValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Value))))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func Int64Value_MarshalVT(m *wrapperspb.Int64Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := Int64Value_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := Int64Value_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func Int64Value_MarshalToVT(m *wrapperspb.Int64Value, dAtA []byte) (int, error) {
	size := Int64Value_SizeVT(m)
	return Int64Value_MarshalToSizedBufferVT(m, dAtA[:size])
}

func Int64Value_MarshalToSizedBufferVT(m *wrapperspb.Int64Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func UInt64Value_MarshalVT(m *wrapperspb.UInt64Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := UInt64Value_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := UInt64Value_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func UInt64Value_MarshalToVT(m *wrapperspb.UInt64Value, dAtA []byte) (int, error) {
	size := UInt64Value_SizeVT(m)
	return UInt64Value_MarshalToSizedBufferVT(m, dAtA[:size])
}

func UInt64Value_MarshalToSizedBufferVT(m *wrapperspb.UInt64Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func Int32Value_MarshalVT(m *wrapperspb.Int32Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := Int32Value_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := Int32Value_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func Int32Value_MarshalToVT(m *wrapperspb.Int32Value, dAtA []byte) (int, error) {
	size := Int32Value_SizeVT(m)
	return Int32Value_MarshalToSizedBufferVT(m, dAtA[:size])
}

func Int32Value_MarshalToSizedBufferVT(m *wrapperspb.Int32Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func UInt32Value_MarshalVT(m *wrapperspb.UInt32Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := UInt32Value_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := UInt32Value_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func UInt32Value_MarshalToVT(m *wrapperspb.UInt32Value, dAtA []byte) (int, error) {
	size := UInt32Value_SizeVT(m)
	return UInt32Value_MarshalToSizedBufferVT(m, dAtA[:size])
}

func UInt32Value_MarshalToSizedBufferVT(m *wrapperspb.UInt32Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func BoolValue_MarshalVT(m *wrapperspb.BoolValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := BoolValue_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := BoolValue_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func BoolValue_MarshalToVT(m *wrapperspb.BoolValue, dAtA []byte) (int, error) {
	size := BoolValue_SizeVT(m)
	return BoolValue_MarshalToSizedBufferVT(m, dAtA[:size])
}

func BoolValue_MarshalToSizedBufferVT(m *wrapperspb.BoolValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func StringValue_MarshalVT(m *wrapperspb.StringValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := StringValue_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := StringValue_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func StringValue_MarshalToVT(m *wrapperspb.StringValue, dAtA []byte) (int, error) {
	size := StringValue_SizeVT(m)
	return StringValue_MarshalToSizedBufferVT(m, dAtA[:size])
}

func StringValue_MarshalToSizedBufferVT(m *wrapperspb.StringValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func BytesValue_MarshalVT(m *wrapperspb.BytesValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := BytesValue_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := BytesValue_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func BytesValue_MarshalToVT(m *wrapperspb.BytesValue, dAtA []byte) (int, error) {
	size := BytesValue_SizeVT(m)
	return BytesValue_MarshalToSizedBufferVT(m, dAtA[:size])
}

func BytesValue_MarshalToSizedBufferVT(m *wrapperspb.BytesValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func DoubleValue_SizeVT(m *wrapperspb.DoubleValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func FloatValue_SizeVT(m *wrapperspb.FloatValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 5
	}
	return n
}

func Int64Value_SizeVT(m *wrapperspb.Int64Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	return n
}

func UInt64Value_SizeVT(m *wrapperspb.UInt64Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	return n
}

func Int32Value_SizeVT(m *wrapperspb.Int32Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	return n
}

func UInt32Value_SizeVT(m *wrapperspb.UInt32Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	return n
}

func BoolValue_SizeVT(m *wrapperspb.BoolValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func StringValue_SizeVT(m *wrapperspb.StringValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func BytesValue_SizeVT(m *wrapperspb.BytesValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
