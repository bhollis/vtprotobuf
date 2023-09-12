// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: google/protobuf/empty.proto

package emptypb

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func Empty_MarshalVT(m *emptypb.Empty) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := Empty_SizeVT(m)
	dAtA = make([]byte, size)
	n, err := Empty_MarshalToSizedBufferVT(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func Empty_MarshalToVT(m *emptypb.Empty, dAtA []byte) (int, error) {
	size := Empty_SizeVT(m)
	return Empty_MarshalToSizedBufferVT(m, dAtA[:size])
}

func Empty_MarshalToSizedBufferVT(m *emptypb.Empty, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func Empty_SizeVT(m *emptypb.Empty) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
