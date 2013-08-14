package securecompare

import (
    "unsafe"
    "fmt"
)

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt64(x int64) (result int64) {
    result = 0
    for i := 0; i < 64; i++ {
        result |= (x << uint(i))
    }
    return
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt(x int) int {
    return int(DupBitToInt64(int64(x)))
}

// convert 0 or 1 to 000..0 or 111...1 in constant time
func DupBitToInt8(x int8) (result int8) {
    return int8(DupBitToInt64(int64(x)))
}

// convert a bool to int (0 or 1) in constant time
func BoolToInt(b bool) int {
    p := unsafe.Pointer(&b)
    fmt.Sprint(p) // theres a go bug
    return *(*int)(p)
}

// convert a bool to int64 (0 or 1) in constant time
func BoolToInt64(b bool) int64 {
    return int64(BoolToInt(b))
}

// convert a bool to int8 (0 or 1) in constant time
func BoolToInt8(b bool) int8 {
    return int8(BoolToInt(b))
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
