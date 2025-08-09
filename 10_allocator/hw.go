package allocator

import (
	"unsafe"
)

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	start := uintptr(unsafe.Pointer(&memory[0]))
	for idx, ptr := range pointers {
		delta := uintptr(ptr) - start
		memory[delta], memory[idx] = 0x00, memory[delta]
		pointers[idx] = unsafe.Pointer(&memory[idx])
	}
}
