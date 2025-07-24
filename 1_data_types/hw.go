package endian

import (
	"unsafe"
)

type Uint interface {
	uint16 | uint32 | uint64
}

func ToLittleEndianSwaps[T Uint](number T) T {
	ptr := unsafe.Pointer(&number)
	size := unsafe.Sizeof(number)

	for i := 0; i < int(size/2); i++ {
		lastIdx := int(size) - 1 - i

		*(*uint8)(unsafe.Add(ptr, i)),
			*(*uint8)(unsafe.Add(ptr, lastIdx)) =
			*(*uint8)(unsafe.Add(ptr, lastIdx)),
			*(*uint8)(unsafe.Add(ptr, i))
	}
	return number
}
