package main

import (
	"slices"
	"unsafe"

	"k8s.io/utils/ptr"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	return COWBuffer{
		data: data,
		refs: ptr.To(1),
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++

	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	if b == nil {
		return
	}

	*b.refs--
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index >= len(b.data) || index < 0 {
		return false
	}

	if *b.refs > 1 {
		*b.refs--
		*b = NewCOWBuffer(slices.Clone(b.data))
	}

	b.data[index] = value

	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
