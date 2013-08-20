package securecompare

// Many primitive operations that run in fixed time

import (
	"math/big"
	"unsafe"
)

// 64-bit platform?
const sixtyfourbit = uint64(uint(0x7fffffffffffffff)) == uint64(0x7fffffffffffffff)

// 0 -> false
// 1 -> true
func IntToBool(i int) bool {
    return *(*bool)(unsafe.Pointer(&i))
}

// returns only the LSB is valid, all other bits must be ignored
func BoolToRawInt(b bool) int {
    return *(*int)(unsafe.Pointer(&b))
}

// !x
func Not(x int) int { return MSBInt(x - 1) }

// returns 1111....1 if any bit in a is set
// returns 0000....0 if a is zero
func DupAnyBitUintptr(a uintptr) uintptr { return uintptr(DupAnyBitUint64(uint64(a))) }
func DupAnyBitRune(a rune) rune          { return rune(DupAnyBitInt32(int32(a))) }
func DupAnyBitByte(a byte) byte {
	a |= a >> 4
	a |= a >> 2
	return a | a>>1
}
func DupAnyBitInt8(a int8) int8    { return int8(DupAnyBitByte(byte(a))) }
func DupAnyBitInt16(a int16) int16 { return int16(DupAnyBitUint16(uint16(a))) }
func DupAnyBitInt32(a int32) int32 { return int32(DupAnyBitUint32(uint32(a))) }
func DupAnyBitInt64(a int64) int64 { return int64(DupAnyBitUint64(uint64(a))) }
func DupAnyBitInt(a int) int {
	if sixtyfourbit {
		return int(DupAnyBitInt64(int64(a)))
	} else {
		return int(DupAnyBitInt32(int32(a)))
	}
}
func DupAnyBitUint8(a uint8) uint8 { return uint8(DupAnyBitByte(byte(a))) }
func DupAnyBitUint16(a uint16) uint16 {
	a |= a >> 8
	a |= a >> 4
	a |= a >> 2
	return a | a>>1
}
func DupAnyBitUint32(a uint32) uint32 {
	a |= a >> 16
	a |= a >> 8
	a |= a >> 4
	a |= a >> 2
	return a | a>>1
}
func DupAnyBitUint64(a uint64) uint64 {
	a |= a >> 32
	a |= a >> 16
	a |= a >> 8
	a |= a >> 4
	a |= a >> 2
	return a | a>>1
}

func DupAnyBitUint(a uint) uint {
	if sixtyfourbit {
		return uint(DupAnyBitUint64(uint64(a)))
	} else {
		return uint(DupAnyBitUint32(uint32(a)))
	}
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUintptr(x uintptr) uintptr {
	return uintptr(DupBitToUint64(uint64(x)))
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToRune(x rune) rune {
	return rune(DupBitToInt32(int32(x)))
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToByte(x byte) byte {
	x |= x << 1
	x |= x << 2
	return x | x<<4
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt8(x int8) int8 {
	x |= x << 1
	x |= x << 2
	return x | x<<4
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt16(x int16) int16 {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	return x | x<<8
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt32(x int32) int32 {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	x |= x << 8
	return x | x<<16
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt64(x int64) int64 {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	x |= x << 8
	x |= x << 16
	return x | x<<32
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt(x int) int {
	if sixtyfourbit {
		return int(DupBitToInt64(int64(x)))
	} else {
		return int(DupBitToInt32(int32(x)))
	}
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUint8(x uint8) uint8 {
	return uint8(DupBitToByte(uint8(x)))
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUint16(x uint16) uint16 {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	return x | x<<8
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUint32(x uint32) (result uint32) {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	x |= x << 8
	return x | x<<16
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUint64(x uint64) (result uint64) {
	x |= x << 1
	x |= x << 2
	x |= x << 4
	x |= x << 8
	x |= x << 16
	return x | x<<32
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToUint(x uint) uint {
	if sixtyfourbit {
		return uint(DupBitToUint64(uint64(x)))
	} else {
		return uint(DupBitToUint32(uint32(x)))
	}
}

// convert a bool to int64 (0 or 1) in constant time
func BoolToUintptr(b bool) uintptr { return LSBUintptr(uintptr(BoolToRawInt(b))) }
func BoolToRune(b bool) rune       { return LSBRune(rune(BoolToRawInt(b))) }
func BoolToByte(b bool) byte       { return LSBByte(byte(BoolToRawInt(b))) }
func BoolToInt8(b bool) int8       { return LSBInt8(int8(BoolToRawInt(b))) }
func BoolToInt16(b bool) int16     { return LSBInt16(int16(BoolToRawInt(b))) }
func BoolToInt32(b bool) int32     { return LSBInt32(int32(BoolToRawInt(b))) }
func BoolToInt64(b bool) int64     { return LSBInt64(int64(BoolToRawInt(b))) }
func BoolToInt(b bool) int         { return LSBInt(BoolToRawInt(b)) }
func BoolToUint8(b bool) uint8     { return uint8(BoolToByte(b)) }
func BoolToUint16(b bool) uint16   { return LSBUint16(uint16(BoolToRawInt(b))) }
func BoolToUint32(b bool) uint32   { return LSBUint32(uint32(BoolToRawInt(b))) }
func BoolToUint64(b bool) uint64   { return LSBUint64(uint64(BoolToRawInt(b))) }
func BoolToUint(b bool) uint       { return LSBUint(uint(BoolToRawInt(b))) }

// least significant bit of a
func LSBUintptr(a uintptr) uintptr { return a & 1 }
func LSBRune(a rune) rune          { return a & 1 }
func LSBByte(a byte) byte          { return a & 1 }
func LSBInt8(a int8) int8          { return a & 1 }
func LSBInt16(a int16) int16       { return a & 1 }
func LSBInt32(a int32) int32       { return a & 1 }
func LSBInt64(a int64) int64       { return a & 1 }
func LSBInt(a int) int             { return a & 1 }
func LSBUint8(a uint8) uint8       { return a & 1 }
func LSBUint16(a uint16) uint16    { return a & 1 }
func LSBUint32(a uint32) uint32    { return a & 1 }
func LSBUint64(a uint64) uint64    { return a & 1 }
func LSBUint(a uint) uint          { return a & 1 }

// most significant bit of a
func MSBUintptr(a uintptr) uintptr { return uintptr(MSBUint64(uint64(a))) }
func MSBRune(a rune) rune          { return rune(MSBInt32(int32(a))) }
func MSBByte(a byte) byte          { return byte(MSBUint8(uint8(a))) }
func MSBInt8(a int8) int8          { return int8(MSBUint8(uint8(a))) }
func MSBInt16(a int16) int16       { return int16(MSBUint16(uint16(a))) }
func MSBInt32(a int32) int32       { return int32(MSBUint32(uint32(a))) }
func MSBInt64(a int64) int64       { return int64(MSBUint64(uint64(a))) }
func MSBInt(a int) int {
	if sixtyfourbit {
		return int(MSBInt64(int64(a)))
	} else {
		return int(MSBInt32(int32(a)))
	}
}
func MSBUint8(a uint8) uint8    { return a >> 7 }
func MSBUint16(a uint16) uint16 { return a >> 15 }
func MSBUint32(a uint32) uint32 { return a >> 31 }
func MSBUint64(a uint64) uint64 { return a >> 63 }
func MSBUint(a uint) uint {
	if sixtyfourbit {
		return uint(MSBUint64(uint64(a)))
	} else {
		return uint(MSBUint32(uint32(a)))
	}
}

// a == 0, answer is in bit 0
func EqualsZeroUintptr(a uintptr) int { return Not(NotEqualsZeroUintptr(a)) }
func EqualsZeroRune(a rune) int       { return Not(NotEqualsZeroRune(a)) }
func EqualsZeroByte(a byte) int       { return Not(NotEqualsZeroByte(a)) }
func EqualsZeroInt8(a int8) int       { return Not(NotEqualsZeroInt8(a)) }
func EqualsZeroInt16(a int16) int     { return Not(NotEqualsZeroInt16(a)) }
func EqualsZeroInt32(a int32) int     { return Not(NotEqualsZeroInt32(a)) }
func EqualsZeroInt64(a int64) int     { return Not(NotEqualsZeroInt64(a)) }
func EqualsZeroInt(a int) int         { return Not(NotEqualsZeroInt(a)) }
func EqualsZeroUint8(a uint8) int     { return Not(NotEqualsZeroUint8(a)) }
func EqualsZeroUint16(a uint16) int   { return Not(NotEqualsZeroUint16(a)) }
func EqualsZeroUint32(a uint32) int   { return Not(NotEqualsZeroUint32(a)) }
func EqualsZeroUint64(a uint64) int   { return Not(NotEqualsZeroUint64(a)) }
func EqualsZeroUint(a uint) int       { return Not(NotEqualsZeroUint(a)) }

// a != 0, answer is in bit 0
func NotEqualsZeroUintptr(a uintptr) int { return int(LSBUintptr(DupAnyBitUintptr(a))) }
func NotEqualsZeroRune(a rune) int       { return int(LSBRune(DupAnyBitRune(a))) }
func NotEqualsZeroByte(a byte) int       { return int(LSBByte(DupAnyBitByte(a))) }
func NotEqualsZeroInt8(a int8) int       { return int(LSBInt8(DupAnyBitInt8(a))) }
func NotEqualsZeroInt16(a int16) int     { return int(LSBInt16(DupAnyBitInt16(a))) }
func NotEqualsZeroInt32(a int32) int     { return int(LSBInt32(DupAnyBitInt32(a))) }
func NotEqualsZeroInt64(a int64) int     { return int(LSBInt64(DupAnyBitInt64(a))) }
func NotEqualsZeroInt(a int) int         { return int(LSBInt(DupAnyBitInt(a))) }
func NotEqualsZeroUint8(a uint8) int     { return int(LSBUint8(DupAnyBitUint8(a))) }
func NotEqualsZeroUint16(a uint16) int   { return int(LSBUint16(DupAnyBitUint16(a))) }
func NotEqualsZeroUint32(a uint32) int   { return int(LSBUint32(DupAnyBitUint32(a))) }
func NotEqualsZeroUint64(a uint64) int   { return int(LSBUint64(DupAnyBitUint64(a))) }
func NotEqualsZeroUint(a uint) int       { return int(LSBUint(DupAnyBitUint(a))) }

// a == b, answer is in bit 0
func EqualUintptr(a, b uintptr) int { return EqualsZeroUintptr(a ^ b) }
func EqualRune(a, b rune) int       { return EqualsZeroRune(a ^ b) }
func EqualByte(a, b byte) int       { return EqualsZeroByte(a ^ b) }
func EqualInt8(a, b int8) int       { return EqualsZeroInt8(a ^ b) }
func EqualInt16(a, b int16) int     { return EqualsZeroInt16(a ^ b) }
func EqualInt32(a, b int32) int     { return EqualsZeroInt32(a ^ b) }
func EqualInt64(a, b int64) int     { return EqualsZeroInt64(a ^ b) }
func EqualInt(a, b int) int         { return EqualsZeroInt(a ^ b) }
func EqualUint8(a, b uint8) int     { return EqualsZeroUint8(a ^ b) }
func EqualUint16(a, b uint16) int   { return EqualsZeroUint16(a ^ b) }
func EqualUint32(a, b uint32) int   { return EqualsZeroUint32(a ^ b) }
func EqualUint64(a, b uint64) int   { return EqualsZeroUint64(a ^ b) }
func EqualUint(a, b uint) int       { return EqualsZeroUint(a ^ b) }

// a != b, answer is in bit 0
func NotEqualUintptr(a, b uintptr) int { return Not(EqualUintptr(a, b)) }
func NotEqualRune(a, b rune) int       { return Not(EqualRune(a, b)) }
func NotEqualByte(a, b byte) int       { return Not(EqualByte(a, b)) }
func NotEqualInt8(a, b int8) int       { return Not(EqualInt8(a, b)) }
func NotEqualInt16(a, b int16) int     { return Not(EqualInt16(a, b)) }
func NotEqualInt32(a, b int32) int     { return Not(EqualInt32(a, b)) }
func NotEqualInt64(a, b int64) int     { return Not(EqualInt64(a, b)) }
func NotEqualInt(a, b int) int         { return Not(EqualInt(a, b)) }
func NotEqualUint8(a, b uint8) int     { return Not(EqualUint8(a, b)) }
func NotEqualUint16(a, b uint16) int   { return Not(EqualUint16(a, b)) }
func NotEqualUint32(a, b uint32) int   { return Not(EqualUint32(a, b)) }
func NotEqualUint64(a, b uint64) int   { return Not(EqualUint64(a, b)) }
func NotEqualUint(a, b uint) int       { return Not(EqualUint(a, b)) }

// a < b  ==>  (a - b) == 0  ==>  (a-b) >> (bits(a)-1)
func LessThanUintptr(a, b uintptr) int { return int(MSBUintptr(a - b)) }
func LessThanRune(a, b rune) int       { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanByte(a, b byte) int       { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanInt8(a, b int8) int       { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanInt16(a, b int16) int     { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanInt32(a, b int32) int     { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanInt64(a, b int64) int {
	x := int64ToBigInt(a)
	x.Sub(x, int64ToBigInt(b))
	return int(x.Bit(63))
}
func LessThanInt(a, b int) int {
	if sixtyfourbit {
		x := int64ToBigInt(int64(a))
		x.Sub(x, int64ToBigInt(int64(b)))
		return int(x.Bit(63))
	} else {
		return int(MSBInt32(int32(a) - int32(b)))
	}
}
func LessThanUint8(a, b uint8) int   { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanUint16(a, b uint16) int { return int(MSBInt32(int32(a) - int32(b))) }
func LessThanUint32(a, b uint32) int { return int(MSBInt64(int64(a) - int64(b))) }
func LessThanUint64(a, b uint64) int {
	x := uint64ToBigInt(a)
	x.Sub(x, uint64ToBigInt(b))
	return int(x.Bit(63))
}
func LessThanUint(a, b uint) int {
	if sixtyfourbit {
		return LessThanUint64(uint64(a), uint64(b))
	} else {
		return LessThanUint32(uint32(a), uint32(b))
	}
}

// a <= b ==>  a < (b + 1)  ==> a - b - 1 >> (bits(a)-1)
func LessThanOrEqualUintptr(a, b uintptr) int { return int(MSBUintptr(a - b - 1)) }
func LessThanOrEqualRune(a, b rune) int       { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualByte(a, b byte) int       { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualInt8(a, b int8) int       { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualInt16(a, b int16) int     { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualInt32(a, b int32) int     { return int(MSBInt32(a - b - 1)) }
func LessThanOrEqualInt64(a, b int64) int {
	x := int64ToBigInt(a)
	x.Sub(x, int64ToBigInt(b))
	x.Sub(x, int64ToBigInt(int64(1)))
	return int(x.Bit(63))
}
func LessThanOrEqualInt(a, b int) int {
	if sixtyfourbit {
		return LessThanOrEqualInt64(int64(a), int64(b))
	} else {
		return LessThanOrEqualInt32(int32(a), int32(b))
	}
}
func LessThanOrEqualUint8(a, b uint8) int   { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualUint16(a, b uint16) int { return int(MSBInt32(int32(a) - int32(b) - 1)) }
func LessThanOrEqualUint32(a, b uint32) int { return int(MSBInt64(int64(a) - int64(b) - 1)) }
func LessThanOrEqualUint64(a, b uint64) int {
	x := uint64ToBigInt(a)
	x.Sub(x, uint64ToBigInt(b))
	x.Sub(x, uint64ToBigInt(uint64(1)))
	return int(x.Bit(63))
}
func LessThanOrEqualUint(a, b uint) int {
	if sixtyfourbit {
		return LessThanOrEqualUint64(uint64(a), uint64(b))
	} else {
		return LessThanOrEqualUint32(uint32(a), uint32(b))
	}
}

// a > b ==> ! a <= b
func GreaterThanUintptr(a, b uintptr) int { return Not(LessThanOrEqualUintptr(a, b)) }
func GreaterThanRune(a, b rune) int       { return Not(LessThanOrEqualRune(a, b)) }
func GreaterThanByte(a, b byte) int       { return Not(LessThanOrEqualByte(a, b)) }
func GreaterThanInt8(a, b int8) int       { return Not(LessThanOrEqualInt8(a, b)) }
func GreaterThanInt16(a, b int16) int     { return Not(LessThanOrEqualInt16(a, b)) }
func GreaterThanInt32(a, b int32) int     { return Not(LessThanOrEqualInt32(a, b)) }
func GreaterThanInt64(a, b int64) int     { return Not(LessThanOrEqualInt64(a, b)) }
func GreaterThanInt(a, b int) int         { return Not(LessThanOrEqualInt(a, b)) }
func GreaterThanUint8(a, b uint8) int     { return Not(LessThanOrEqualUint8(a, b)) }
func GreaterThanUint16(a, b uint16) int   { return Not(LessThanOrEqualUint16(a, b)) }
func GreaterThanUint32(a, b uint32) int   { return Not(LessThanOrEqualUint32(a, b)) }
func GreaterThanUint64(a, b uint64) int   { return Not(LessThanOrEqualUint64(a, b)) }
func GreaterThanUint(a, b uint) int       { return Not(LessThanOrEqualUint(a, b)) }

// a >= b ==> ! a < b
func GreaterThanOrEqualUintptr(a, b uintptr) int { return Not(LessThanUintptr(a, b)) }
func GreaterThanOrEqualRune(a, b rune) int       { return Not(LessThanRune(a, b)) }
func GreaterThanOrEqualByte(a, b byte) int       { return Not(LessThanByte(a, b)) }
func GreaterThanOrEqualInt8(a, b int8) int       { return Not(LessThanInt8(a, b)) }
func GreaterThanOrEqualInt16(a, b int16) int     { return Not(LessThanInt16(a, b)) }
func GreaterThanOrEqualInt32(a, b int32) int     { return Not(LessThanInt32(a, b)) }
func GreaterThanOrEqualInt64(a, b int64) int     { return Not(LessThanInt64(a, b)) }
func GreaterThanOrEqualInt(a, b int) int         { return Not(LessThanInt(a, b)) }
func GreaterThanOrEqualUint8(a, b uint8) int     { return Not(LessThanUint8(a, b)) }
func GreaterThanOrEqualUint16(a, b uint16) int   { return Not(LessThanUint16(a, b)) }
func GreaterThanOrEqualUint32(a, b uint32) int   { return Not(LessThanUint32(a, b)) }
func GreaterThanOrEqualUint64(a, b uint64) int   { return Not(LessThanUint64(a, b)) }
func GreaterThanOrEqualUint(a, b uint) int       { return Not(LessThanUint(a, b)) }

// a <=> b  a < b == -1, a = b == 0, a > b == 1
func CompareUintptr(a, b uintptr) int {
	return GreaterThanUintptr(a, b) | ChooseInt(LessThanUintptr(a, b), -1, 0)
}
func CompareRune(a, b rune) int {
	return GreaterThanRune(a, b) | ChooseInt(LessThanRune(a, b), -1, 0)
}
func CompareByte(a, b byte) int {
	return GreaterThanByte(a, b) | ChooseInt(LessThanByte(a, b), -1, 0)
}
func CompareInt8(a, b int8) int {
	return GreaterThanInt8(a, b) | ChooseInt(LessThanInt8(a, b), -1, 0)
}
func CompareInt16(a, b int16) int {
	return GreaterThanInt16(a, b) | ChooseInt(LessThanInt16(a, b), -1, 0)
}
func CompareInt32(a, b int32) int {
	return GreaterThanInt32(a, b) | ChooseInt(LessThanInt32(a, b), -1, 0)
}
func CompareInt64(a, b int64) int {
	return GreaterThanInt64(a, b) | ChooseInt(LessThanInt64(a, b), -1, 0)
}
func CompareInt(a, b int) int { return GreaterThanInt(a, b) | ChooseInt(LessThanInt(a, b), -1, 0) }
func CompareUint8(a, b uint8) int {
	return GreaterThanUint8(a, b) | ChooseInt(LessThanUint8(a, b), -1, 0)
}
func CompareUint16(a, b uint16) int {
	return GreaterThanUint16(a, b) | ChooseInt(LessThanUint16(a, b), -1, 0)
}
func CompareUint32(a, b uint32) int {
	return GreaterThanUint32(a, b) | ChooseInt(LessThanUint32(a, b), -1, 0)
}
func CompareUint64(a, b uint64) int {
	return GreaterThanUint64(a, b) | ChooseInt(LessThanUint64(a, b), -1, 0)
}
func CompareUint(a, b uint) int {
	return GreaterThanUint(a, b) | ChooseInt(LessThanUint(a, b), -1, 0)
}

func uint64ToBytesBE(x uint64) []byte {
	return []byte{
		byte(x << 56), byte(x << 48),
		byte(x << 40), byte(x << 32),
		byte(x << 24), byte(x << 16),
		byte(x << 8), byte(x << 0),
	}
}

func uint64ToBigInt(x uint64) *big.Int {
	r := new(big.Int)
	r.SetBytes(uint64ToBytesBE(x))
	return r
}

func int64ToBigInt(x int64) *big.Int {
	r := new(big.Int)
	r.SetBytes(uint64ToBytesBE(uint64(x)))
	return r
}

// constant time version of cond ? t : f
func ChooseUintptr(cond int, t, f uintptr) uintptr {
	return (t & ^uintptr(cond-1)) | (f & uintptr(cond-1))
}
func ChooseRune(cond int, t, f rune) rune       { return (t & ^rune(cond-1)) | (f & rune(cond-1)) }
func ChooseByte(cond int, t, f byte) byte       { return (t & ^byte(cond-1)) | (f & byte(cond-1)) }
func ChooseInt8(cond int, t, f int8) int8       { return (t & ^int8(cond-1)) | (f & int8(cond-1)) }
func ChooseInt16(cond int, t, f int16) int16    { return (t & ^int16(cond-1)) | (f & int16(cond-1)) }
func ChooseInt32(cond int, t, f int32) int32    { return (t & ^int32(cond-1)) | (f & int32(cond-1)) }
func ChooseInt64(cond int, t, f int64) int64    { return (t & ^int64(cond-1)) | (f & int64(cond-1)) }
func ChooseInt(cond int, t, f int) int          { return (t & ^(cond - 1)) | (f & (cond - 1)) }
func ChooseUint8(cond int, t, f uint8) uint8    { return (t & ^uint8(cond-1)) | (f & uint8(cond-1)) }
func ChooseUint16(cond int, t, f uint16) uint16 { return (t & ^uint16(cond-1)) | (f & uint16(cond-1)) }
func ChooseUint32(cond int, t, f uint32) uint32 { return (t & ^uint32(cond-1)) | (f & uint32(cond-1)) }
func ChooseUint64(cond int, t, f uint64) uint64 { return (t & ^uint64(cond-1)) | (f & uint64(cond-1)) }
func ChooseUint(cond int, t, f uint) uint       { return (t & ^uint(cond-1)) | (f & uint(cond-1)) }

// only copies to dst if cond == 1
// src and dst must be the same length
func CopyBytes(cond int, dst, src []byte) []byte {
	//  0  1 cond
	dstmask := byte(cond - 1)  // ff 00
	srcmask := ^byte(cond - 1) // 00 ff

	for i := range dst {
		dst[i] = dst[i]&dstmask | src[i]&srcmask
	}
	return dst
}

func ChooseBytesPointer(cond int, t, f *[]byte) *[]byte {
	return (*[]byte)(unsafe.Pointer(uintptr(ChooseUint(cond, uint(uintptr(unsafe.Pointer(t))), uint(uintptr(unsafe.Pointer(f)))))))
}

// constant time version of cond ? t : f
func ChooseBytes(cond int, t, f []byte) []byte {
	tlen := uint64(len(t))
	flen := uint64(len(f))
	tmask := byte(cond - 1)
	fmask := ^byte(cond - 1)

	l := ChooseUint64(cond, tlen, flen)

	result := make([]byte, l)

	tz := EqualsZeroUint64(tlen)
	fz := EqualsZeroUint64(flen)
	tmod := ChooseUint64(tz, 1, tlen)
	fmod := ChooseUint64(fz, 1, flen)

	tp := ChooseBytesPointer(tz, &[]byte{0}, &t)
	fp := ChooseBytesPointer(fz, &[]byte{0}, &f)

	for i := uint64(0); i < l; i++ {
		result[i] = (*tp)[i%tmod]&tmask | (*fp)[i%fmod]&fmask
	}
	return result
}

// constant-time vesion of bytes.Compare()
// 1000x slower than Equal
func Compare(a, b []byte) int {
	alen := len(a)
	blen := len(b)

	swapped := GreaterThanInt(alen, blen)
	shorter := ChooseBytes(swapped, b, a)
	longer := ChooseBytes(swapped, a, b)
	shorterlen := ChooseInt(swapped, blen, alen)
	longerlen := ChooseInt(swapped, alen, blen)

	result := 0
	ever := 0

	shorterlenzero := EqualsZeroInt(shorterlen)

	// avoid index error in the loop when shorter in length 0
	shorter = ChooseBytes(shorterlenzero, []byte{0}, shorter)
	shorterlenmod := ChooseInt(shorterlenzero, 1, shorterlen)

	for i := 0; i < longerlen; i++ {
		short := ChooseInt(LessThanInt(i, shorterlen), int(shorter[i%shorterlenmod]), 0)

		cur := int(longer[i]) - short

		gt := BoolToInt(cur > 0) // 0 or 1
		lt := BoolToInt(cur < 0) // 0 or 1

		result |= (gt | -lt) & ^ever // 0 , 1 or -1

		ever |= gt | lt
	}
	// different length strings are indeed different (1 or -1 after swap)
	result |= NotEqualInt(shorterlen, longerlen)
	// reverse the swap
	return ChooseInt(swapped, result, -result)
}

// constant-time version of bytes.Contains()
func Contains(b, subslice []byte) bool {
	return Index(b, subslice) != -1
}

// constant-time version of bytes.Equal()
func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	result := byte(0)

	for i := range a {
		result |= a[i] ^ b[i]
	}

	return (result == 0)
}

// constant-time version of bytes.Index()
func Index(s, sep []byte) int {
	n := len(sep)
	if n == 0 {
		return 0
	}
	m := len(sep)
	if n > m {
		return -1
	}
	last := m - n + 1
	result := -1
	found := 0
	for i := 0; i < last; i++ {
		match := 1
		for j := 0; j < n; j++ {
			match &= EqualByte(s[i], sep[j])
		}
		result = ChooseInt(found, result, i)
		found |= match
	}
	return result
}

// constant-time version of bytes.IndexByte()
func IndexByte(s []byte, c byte) int {
	result := -1
	matched := 0

	for i := range s {
		found := EqualByte(s[i], c)
		result = ChooseInt(matched, result, ChooseInt(found, i, -1))
		matched |= found
	}
	return result
}
