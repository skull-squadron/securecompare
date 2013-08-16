package securecompare

import (
    "encoding/hex"
    "testing"
    "fmt"
    "bytes"
)

func make1mib(b byte) []byte {
        result := make([]byte, 1024*1024)
        for i := range result {
            result[i] = b
        }
        return result
}

var (
    data0, _ = hex.DecodeString("ffff")
    data1, _ = hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data2, _ = hex.DecodeString("01bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data3, _ = hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50de")
    data4, _ = hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50")
    data5, _ = hex.DecodeString("")
    data6, _ = hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d5000")
    data7, _ = hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d500000")
    data8, _ = hex.DecodeString("00")
    mib0 = make1mib(0xAA)
    mib1 = make1mib(0x55)
    mib2 = make1mib(0x00)
    mib3 = make1mib(0x00)
)

func TestDupBitToInt64(t *testing.T) {
    if DupBitToInt64(0) != 0 {
        t.Error("DupBitToInt64 0 fails")
    }
    if DupBitToInt64(1) != int64(-1) {
        t.Error("DupBitToInt64 1 fails")
    }
}

func TestDupBitToInt(t *testing.T) {
    if DupBitToInt(0) != 0 {
        t.Error("DupBitToInt 0 fails")
    }
    if DupBitToInt(1) != -1 {
        t.Error("DupBitToInt 1 fails")
    }
}

func TestDupBitToInt8(t *testing.T) {
    if DupBitToInt8(0) != 0 {
        t.Error("DupBitToInt8 0 fails")
    }
    if DupBitToInt8(1) != -1 {
        t.Error("DupBitToInt8 1 fails")
    }
}

func TestBoolToInt64(t *testing.T) {
    if BoolToInt64(false) != 0 {
        t.Error("BoolToInt64 false fails")
    }
    if BoolToInt64(true) != 1 {
        t.Error("BoolToInt64 true fails")
    }
}

func TestBoolToInt(t *testing.T) {
    if BoolToInt(false) != 0 {
        t.Error("BoolToInt false fails")
    }
    if BoolToInt(true) != 1 {
        t.Error("BoolToInt true fails")
    }
}

func TestBoolToInt8(t *testing.T) {
    if BoolToInt8(false) != 0 {
        t.Error("BoolToInt8 false fails")
    }
    if BoolToInt8(true) != 1 {
        t.Error("BoolToInt8 true fails")
    }
}

func TestChooseInt(t *testing.T) {
    if ChooseInt(false, 0, 0) != 0 {
        t.Error("ChooseInt fails 0")
    }
    if ChooseInt(true, 0, 0) != 0 {
        t.Error("ChooseInt fails 1")
    }
    if ChooseInt(false, 0, 1) != 1 {
        t.Error("ChooseInt fails 2")
    }
    if ChooseInt(true, 1, 0) != 1 {
        t.Error("ChooseInt fails 3")
    }
    if ChooseInt(true, 1, 1) != 1 {
        t.Error("ChooseInt fails 4")
    }
    if ChooseInt(false, 1, 2) != 2 {
        t.Error("ChooseInt fails 5")
    }
    if ChooseInt(true, 2, 1) != 2 {
        t.Error("ChooseInt fails 6")
    }
}

func TestChooseInt8(t *testing.T) {
    if ChooseInt8(false, 0, 0) != 0 {
        t.Error("ChooseInt8 fails 0")
    }
    if ChooseInt8(true, 0, 0) != 0 {
        t.Error("ChooseInt8 fails 1")
    }
    if ChooseInt8(false, 0, 1) != 1 {
        t.Error("ChooseInt8 fails 2")
    }
    if ChooseInt8(true, 1, 0) != 1 {
        t.Error("ChooseInt8 fails 3")
    }
    if ChooseInt8(true, 1, 1) != 1 {
        t.Error("ChooseInt8 fails 4")
    }
    if ChooseInt8(false, 1, 2) != 2 {
        t.Error("ChooseInt8 fails 5")
    }
    if ChooseInt8(true, 2, 1) != 2 {
        t.Error("ChooseInt8 fails 6")
    }
}

func TestChooseInt64(t *testing.T) {
    if ChooseInt64(false, 0, 0) != 0 {
        t.Error("ChooseInt64 fails 0")
    }
    if ChooseInt64(true, 0, 0) != 0 {
        t.Error("ChooseInt64 fails 1")
    }
    if ChooseInt64(false, 0, 1) != 1 {
        t.Error("ChooseInt64 fails 2")
    }
    if ChooseInt64(true, 1, 0) != 1 {
        t.Error("ChooseInt64 fails 3")
    }
    if ChooseInt64(true, 1, 1) != 1 {
        t.Error("ChooseInt64 fails 4")
    }
    if ChooseInt64(false, 1, 2) != 2 {
        t.Error("ChooseInt64 fails 5")
    }
    if ChooseInt64(true, 2, 1) != 2 {
        t.Error("ChooseInt64 fails 6")
    }
}

func TestChooseBytes(t *testing.T) {
    if bytes.Compare(ChooseBytes(false, data1, data1), data1) != 0 {
        t.Error("ChooseBytes fails 0")
    }
    if bytes.Compare(ChooseBytes(true, data1, data1), data1) != 0 {
        t.Error("ChooseBytes fails 1")
    }
    if bytes.Compare(ChooseBytes(false, data1, data2), data2) != 0 {
        t.Error("ChooseBytes fails 2")
    }
    if bytes.Compare(ChooseBytes(true, data2, data1), data2) != 0 {
        t.Error("ChooseBytes fails 3")
    }
    if bytes.Compare(ChooseBytes(true, data2, data2), data2) != 0 {
        t.Error("ChooseBytes fails 4")
    }
    if bytes.Compare(ChooseBytes(false, data2, data3), data3) != 0 {
        t.Error("ChooseBytes fails 5")
    }
    if bytes.Compare(ChooseBytes(true, data3, data2), data3) != 0 {
        t.Error("ChooseBytes fails 6")
    }
}

func testEqual(t *testing.T, a, b []byte) {
    if bytes.Equal(a, b) != Equal(a, b) {
        t.Error("Equal failed")
        return
    }
}

func testEqualSuccess(t *testing.T, a, b []byte) {
    testEqual(t, a, b)
}

func testEqualFail(t *testing.T, a, b []byte) {
    testEqual(t, a, b)
}

func TestEqual(t *testing.T) {
    testEqualFail(t, data1, data2)
    testEqualFail(t, data1, data3)
    testEqualFail(t, data1, data4)
    testEqualFail(t, data1, data5)
    testEqualSuccess(t, data1, data1)
    testEqualSuccess(t, data5, data8)
    testEqualSuccess(t, data5, data5)
    testEqualSuccess(t, data8, data5)
}

func testCompare(t *testing.T, a, b []byte) {
    expected := bytes.Compare(a, b)
    if actual := Compare(a, b); actual != expected {
        t.Error(fmt.Sprint("failed a=", a, " b=", b, " expected=",expected," actual=",actual))
    }
}

func TestCompare(t *testing.T) {
    testCompare(t, data1, data2)
    testCompare(t, data1, data3)
    testCompare(t, data1, data4)
    testCompare(t, data1, data5)
    testCompare(t, data1, data1)
    testCompare(t, data5, data5)
    testCompare(t, data5, data5)
    testCompare(t, data4, data6)
    testCompare(t, data6, data4)
    testCompare(t, data6, data7)
    testCompare(t, data7, data6)
    testCompare(t, data8, data7)
    testCompare(t, data8, data5)
}

func testIndex(t *testing.T, s, sep []byte) {
    expected := bytes.Index(s, sep)
    if actual := Index(s, sep); actual != expected {
        t.Error(fmt.Sprint("failed s=", s, " b=", sep, " expected=",expected," actual=",actual))
    }

}

func testIndexByte(t *testing.T, s []byte, c byte) {
    expected := bytes.IndexByte(s, c)
    if actual := IndexByte(s, c); actual != expected {
        t.Error(fmt.Sprint("failed s=", s, " c=", c, " expected=",expected," actual=",actual))
    }

}

func TestIndex(t *testing.T) {
    testIndex(t, []byte{0x11, 0x22, 0x33}, []byte{0x11})
    testIndex(t, []byte{0x11, 0x22, 0x33}, []byte{0x11, 0x22})
}

func TestIndexByte(t *testing.T) {
    testIndexByte(t, []byte{0x11, 0x22, 0x33}, 0x11)
    testIndexByte(t, []byte{0x11, 0x22, 0x33}, 0x22)
    testIndexByte(t, []byte{0x11, 0x22, 0x33}, 0x55)
}

// benchmarks

func BenchmarkBoolToIntWarmUp(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt(true)
    }
}

func BenchmarkBoolToIntFalse(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt(false)
    }
}

func BenchmarkBoolToIntTrue(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt(true)
    }
}

func BenchmarkBoolToInt64WarmUp(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt64(true)
    }
}

func BenchmarkBoolToInt64False(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt64(false)
    }
}

func BenchmarkBoolToInt64True(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt64(true)
    }
}

func BenchmarkBoolToInt8WarmUp(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt8(true)
    }
}

func BenchmarkBoolToInt8False(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt8(false)
    }
}

func BenchmarkBoolToInt8True(b *testing.B) {
    for i := 0; i < b.N; i++ {
        BoolToInt8(true)
    }
}

func BenchmarkDupBitToInt64(b *testing.B) {
    for i := 0; i < b.N; i++ {
        DupBitToInt64(0)
    }
}

func BenchmarkDupBitToInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        DupBitToInt(0)
    }
}

func BenchmarkDupBitToInt8(b *testing.B) {
    for i := 0; i < b.N; i++ {
        DupBitToInt8(0)
    }
}

func BenchmarkEqual1MiB0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Equal(mib0, mib0)
    }
}

func BenchmarkEqual1MiB1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Equal(mib0, mib1)
    }
}

func BenchmarkEqual1MiB2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Equal(mib0, mib2)
    }
}

func BenchmarkEqual1MiB3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Equal(mib1, mib2)
    }
}

func BenchmarkEqual1MiB4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Equal(mib2, mib3)
    }
}

func BenchmarkCompare1MiB0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(mib0, mib0)
    }
}

func BenchmarkCompare1MiB1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(mib0, mib1)
    }
}

func BenchmarkCompare1MiB2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(mib0, mib2)
    }
}

func BenchmarkCompare1MiB3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(mib1, mib2)
    }
}

func BenchmarkCompare1MiB4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(mib2, mib3)
    }
}
