package securecompare

import "C"
import (
    "unsafe"
)

// 64-bit platform?
const sixtyfourbit = uint64(uint(0x7fffffffffffffff)) == uint64(0x7fffffffffffffff)

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt64(x int64) (result int64) {
    x1 := x | x<<1
    x2 := x1 | x1<<2
    x3 := x2 | x2<<4
    x4 := x3 | x3<<8
    x5 := x4 | x4<<16
    return x5 | x5<<32
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt32(x int32) (result int32) {
    x1 := x | x<<1
    x2 := x1 | x1<<2
    x3 := x2 | x2<<4
    x4 := x3 | x3<<8
    return x4 | x4<<16
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
func DupBitToInt8(x int8) int8 {
    x1 := x | x<<1
    x2 := x1 | x1<<2
    return x2 | x2<<4
}

// convert a bool to int (0 or 1) in constant time
func BoolToInt(b bool) int {
    result := int(*(*int)(unsafe.Pointer(&b)))
    if sixtyfourbit { // 64-bit platforms
        result <<= 63
        result = int(uint(result) >> 63)
    } else { // 32-bit platforms
        result <<= 31
        result = int(uint(result) >> 31)
    }
    return result
}

// convert a bool to int64 (0 or 1) in constant time
func BoolToInt64(b bool) int64 {
    result := int64(*(*int)(unsafe.Pointer(&b)))
    result <<= 63
    return int64(uint64(result) >> 63)
}

// convert a bool to int8 (0 or 1) in constant time
func BoolToInt8(b bool) int8 {
    result := int8(*(*int)(unsafe.Pointer(&b)))
    result <<= 7
    return int8(uint8(result) >> 7)
}

// constant time version of cond ? t : f
func ChooseInt(cond bool, t, f int) int {
    mask := DupBitToInt(BoolToInt(cond))
    return (t & mask) | (f & ^mask)
}

// constant time version of cond ? t : f
func ChooseInt8(cond bool, t, f int8) int8 {
    mask := DupBitToInt8(BoolToInt8(cond))
    return (t & mask) | (f & ^mask)
}

// constant time version of cond ? t : f
func ChooseInt64(cond bool, t, f int64) int64 {
    mask := DupBitToInt64(BoolToInt64(cond))
    return (t & mask) | (f & ^mask)
}

// constant time version of cond ? t : f
func ChooseBytes(cond bool, t, f []byte) []byte {
    ptrt := uintptr(unsafe.Pointer(&t))
    ptrf := uintptr(unsafe.Pointer(&f))
    mask := uintptr(DupBitToInt(BoolToInt(cond)))
    ptr := (ptrt & mask) | (ptrf & ^mask)
    return *(*[]byte)(unsafe.Pointer(ptr))
}

// constant-time vesion of bytes.Compare()
// 1000x slower than Equal
func Compare(a, b []byte) int {
    alen := len(a)
    blen := len(b)

    swapped     := (alen >= blen)
    shorter     := ChooseBytes(swapped, b, a)
    longer      := ChooseBytes(swapped, a, b)
    shorterlen  := ChooseInt(swapped, blen, alen)
    longerlen   := ChooseInt(swapped, alen, blen)

    result := 0
    ever := 0

    // avoid index error in the loop
    shorter = ChooseBytes(shorterlen == 0, []byte{0}, shorter)
    shorterlenmod := ChooseInt(shorterlen == 0, 1, shorterlen)

    for i := 0; i < longerlen; i++  {
        short := ChooseInt(i < shorterlen, int(shorter[i % shorterlenmod]), 0)

        cur := int(longer[i]) - short

        gt := BoolToInt(cur > 0) // 0 or 1
        lt := BoolToInt(cur < 0) // 0 or 1

        result |= (gt | -lt) & ^ever // 0 , 1 or -1

        ever   |= gt | lt
    }
    // different length strings are indeed different (1 or -1 after swap)
    result |= BoolToInt(longerlen != shorterlen)
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
            match &= BoolToInt(s[i] == sep[j])
        }
        result = ChooseInt(found == 0, i, result)
        found |= match
    }
    return result
}

// constant-time version of bytes.IndexByte()
func IndexByte(s []byte, c byte) int {
    result := -1
    matched := 0

    for i := range s {
        found := s[i] == c
        result = ChooseInt(matched == 0, ChooseInt(found, i, -1), result)
        matched |= BoolToInt(found)
    }
    return result
}
