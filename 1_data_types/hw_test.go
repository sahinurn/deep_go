package endian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversionUint16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: uint16(0x0000),
			result: uint16(0x0000),
		},
		"test case #2": {
			number: uint16(0x0010),
			result: uint16(0x1000),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianSwaps(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
func TestConversionUint32(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: uint32(0x00000000),
			result: uint32(0x00000000),
		},
		"test case #2": {
			number: uint32(0xFFFFFFFF),
			result: uint32(0xFFFFFFFF),
		},
		"test case #3": {
			number: uint32(0x00FF00FF),
			result: uint32(0xFF00FF00),
		},
		"test case #4": {
			number: uint32(0x0000FFFF),
			result: uint32(0xFFFF0000),
		},
		"test case #5": {
			number: uint32(0x01020304),
			result: uint32(0x04030201),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianSwaps(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversionUint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: uint64(0x0000000000000000),
			result: uint64(0x0000000000000000),
		},
		"test case #2": {
			number: uint64(0xFFFFFFFF00000000),
			result: uint64(0x00000000FFFFFFFF),
		},
		"test case #3": {
			number: uint64(0xFF00FF00FF00FF00),
			result: uint64(0x00FF00FF00FF00FF),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianSwaps(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func BenchmarkToLittleEndianSwaps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ToLittleEndianSwaps(uint64(0x00FF00FF00FF00FF))
	}
}
