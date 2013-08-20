package securecompare

import (
    "encoding/hex"
    "testing"
    "fmt"
    "bytes"
)

const (
    uintptrz uintptr = 0
    uintptr1 uintptr = 1
    uintptrn = uintptr(uintn)
    runez rune = 0
    rune1 rune = 1
    runen rune = -1
    bytez byte = 0
    byte1 byte = 1
    byten byte = 0xff
    int8z int8 = 0
    int81 int8 = 1
    int8n int8 = -1
    int16z int16 = 0
    int161 int16 = 1
    int16n int16 = -1
    int32z int32 = 0
    int321 int32 = 1
    int32n int32 = -1
    int64z int64 = 0
    int641 int64 = 1
    int64n int64 = -1
    // intn and intz are pointless
    uint8z uint8 = 0
    uint81 uint8 = 1
    uint8n uint8 = 0xff
    uint16z uint16 = 0
    uint161 uint16 = 1
    uint16n uint16 = 0xffff
    uint32z uint32 = 0
    uint321 uint32 = 1
    uint32n uint32 = 0xffffffff
    uint64z uint64 = 0
    uint641 uint64 = 1
    uint64n uint64 = 0xffffffffffffffff
    uintz uint = 0
    uint1 uint = 1
    uintn uint = 0xffffffffffffffff
)

func TestIntToBool(t *testing.T) {
    var (msg string; actual, expected interface{})

    if msg, actual, expected = "IntToBool", IntToBool(0), false; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "IntToBool", IntToBool(1), true; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
}

func TestBoolToRawInt(t *testing.T) {
    var (msg string; actual, expected interface{})

    if msg, actual, expected = "BoolToRawInt", BoolToRawInt(false) & 1, 0 ; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToRawInt", BoolToRawInt(true) & 1, 1 ; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
}

func TestNot(t *testing.T) {
    var (msg string; actual, expected interface{})

    if msg, actual, expected = "Not", Not(0), 1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "Not", Not(1), 0; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
}

// DupBitTo...
func TestDupBitToRune(t *testing.T) {
    if DupBitToRune(0) != 0 {
        t.Error("DupBitToRune 0 fails")
    }
    if DupBitToByte(1) != byten {
        t.Error("DupBitToRune 1 fails")
    }
}


func TestDupBitToByte(t *testing.T) {
    if DupBitToByte(0) != 0 {
        t.Error("DupBitToByte 0 fails")
    }
    if DupBitToByte(1) != byten {
        t.Error("DupBitToByte 1 fails")
    }
}

func TestDupBitToInt8(t *testing.T) {
    if DupBitToInt8(0) != 0 {
        t.Error("DupBitToInt8 0 fails")
    }
    if DupBitToInt8(1) != int8n {
        t.Error("DupBitToInt8 1 fails")
    }
}

func TestDupBitTo(t *testing.T) {
    var (msg string; actual, expected interface{})

    if msg, actual, expected = "DupBitToUintptr", DupBitToUintptr(0), uintptrz; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUintptr", DupBitToUintptr(0), uintptrz; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUintptr", DupBitToUintptr(1), uintptrn; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToRune", DupBitToRune(0), runez; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToRune", DupBitToRune(1), runen; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToByte", DupBitToByte(0), bytez; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToByte", DupBitToByte(1), byten; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt8", DupBitToInt8(0), int8z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt8", DupBitToInt8(1), int8n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt16", DupBitToInt16(0), int16z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt16", DupBitToInt16(1), int16n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt32", DupBitToInt32(0), int32z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt32", DupBitToInt32(1), int32n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt64", DupBitToInt64(0), int64z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt64", DupBitToInt64(1), int64n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt", DupBitToInt(0), 0; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToInt", DupBitToInt(1), -1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint8", DupBitToUint8(0), uint8z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint8", DupBitToUint8(1), uint8n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint16", DupBitToUint16(0), uint16z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint16", DupBitToUint16(1), uint16n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint32", DupBitToUint32(0), uint32z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint32", DupBitToUint32(1), uint32n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint64", DupBitToUint64(0), uint64z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint64", DupBitToUint64(1), uint64n; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint", DupBitToUint(0), uintz; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "DupBitToUint", DupBitToUint(1), uintn; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
}


// BoolTo...
func TestBoolTo(t *testing.T) {
    var (msg string; actual, expected interface{})

    if msg, actual, expected = "BoolToUintptr", BoolToUintptr(false), uintptrz; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUintptr", BoolToUintptr(true), uintptr1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToRune", BoolToRune(false), runez; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToRune", BoolToRune(true), rune1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToByte", BoolToByte(false), bytez; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToByte", BoolToByte(true), byte1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt8", BoolToInt8(false), int8z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt8", BoolToInt8(true), int81; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt16", BoolToInt16(false), int16z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt16", BoolToInt16(true), int161; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt32", BoolToInt32(false), int32z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt32", BoolToInt32(true), int321; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt64", BoolToInt64(false), int64z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt64", BoolToInt64(true), int641; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt", BoolToInt(false), 0; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToInt", BoolToInt(true), 1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint8", BoolToUint8(false), uint8z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint8", BoolToUint8(true), uint81; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint16", BoolToUint16(false), uint16z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint16", BoolToUint16(true), uint161; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint32", BoolToUint32(false), uint32z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint32", BoolToUint32(true), uint321; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint64", BoolToUint64(false), uint64z; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint64", BoolToUint64(true), uint641; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint", BoolToUint(false), uintz; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
    if msg, actual, expected = "BoolToUint", BoolToUint(true), uint1; actual != expected { t.Error(msg, "fails expected=", expected, "actual=", actual) }
}


// LSB...

var lsbTests = [][]int64 {
    // value expected
    []int64{ 0, 0 },
    []int64{ 1, 1 },
    []int64{ 2, 0 },
    []int64{ 0xabcd00, 0 },
    []int64{ 0xabcdff, 1 },
    []int64{ -2, 0 },
    []int64{ -1, 1 },
}

func TestLSB(t *testing.T) {

}

func TestLSBRune(t *testing.T) {
    for i, test := range lsbTests {
        if LSBRune(rune(test[0])) != rune(test[1]) {
            t.Error("LSBRune fails", i)
        }
    }
}

func TestLSBByte(t *testing.T) {
    for i, test := range lsbTests {
        if LSBByte(byte(test[0])) != byte(test[1]) {
            t.Error("LSBByte fails", i)
        }
    }
}

func TestLSBInt8(t *testing.T) {
    for i, test := range lsbTests {
        if LSBInt8(int8(test[0])) != int8(test[1]) {
            t.Error("LSBInt8 fails", i)
        }
    }
}

func TestLSBInt16(t *testing.T) {
    for i, test := range lsbTests {
        if LSBInt16(int16(test[0])) != int16(test[1]) {
            t.Error("LSBInt16 fails", i)
        }
    }
}

func TestLSBInt32(t *testing.T) {
    for i, test := range lsbTests {
        if LSBInt32(int32(test[0])) != int32(test[1]) {
            t.Error("LSBInt32 fails", i)
        }
    }
}

func TestLSBInt64(t *testing.T) {
    for i, test := range lsbTests {
        if LSBInt64(int64(test[0])) != int64(test[1]) {
            t.Error("LSBInt64 fails", i)
        }
    }
}

func TestLSBInt(t *testing.T) {
    for i, test := range lsbTests {
        if LSBInt(int(test[0])) != int(test[1]) {
            t.Error("LSBInt fails", i)
        }
    }
}

func TestLSBUint8(t *testing.T) {
    for i, test := range lsbTests {
        if LSBUint8(uint8(test[0])) != uint8(test[1]) {
            t.Error("LSBUint8 fails", i)
        }
    }
}

func TestLSBUint16(t *testing.T) {
    for i, test := range lsbTests {
        if LSBUint16(uint16(test[0])) != uint16(test[1]) {
            t.Error("LSBUint16 fails", i)
        }
    }
}

func TestLSBUint32(t *testing.T) {
    for i, test := range lsbTests {
        if LSBUint32(uint32(test[0])) != uint32(test[1]) {
            t.Error("LSBUint32 fails", i)
        }
    }
}

func TestLSBUint64(t *testing.T) {
    for i, test := range lsbTests {
        if LSBUint64(uint64(test[0])) != uint64(test[1]) {
            t.Error("LSBUint64 fails", i)
        }
    }
}

func TestLSBUint(t *testing.T) {
    for i, test := range lsbTests {
        if LSBUint(uint(test[0])) != uint(test[1]) {
            t.Error("LSBUint fails", i)
        }
    }
}

var msbTests = []int64 {
    0,
    1,
    2,
    0x55,
    0x7f,
}

const (
    msbuint8 = uint8(0x80)
    msbuint16 = uint16(0x8000)
    msbuint32 = uint32(0x80000000)
    msbuint64 = uint64(0x8000000000000000)
    msbint8 = int8(-128)
    msbint16 = int16(-32768)
    msbint32 = int32(-2147483648)
    msbint64 = int64(-9223372036854775808)
    msbbyte = byte(msbuint8)
    msbrune = rune(msbint32)
)

func TestMSBRune(t *testing.T) {
    for i, test := range msbTests {
        if MSBRune(rune(test)) != 0 {
            t.Error("MSBRune unset fails", i)
        }
    }
    for i, test := range msbTests {
        x := rune(test) | msbrune
        actual := MSBRune(x)
        expected := rune(1)
        if actual != expected {
            t.Error("MSBRune set fails test=", i, "value=", x, "actual=", actual, "expected=", expected)
        }
    }
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


func TestChooseInt(t *testing.T) {
    if ChooseInt(0, 0, 0) != 0 {
        t.Error("ChooseInt fails 0")
    }
    if ChooseInt(1, 0, 0) != 0 {
        t.Error("ChooseInt fails 1")
    }
    if ChooseInt(0, 0, 1) != 1 {
        t.Error("ChooseInt fails 2")
    }
    if ChooseInt(1, 1, 0) != 1 {
        t.Error("ChooseInt fails 3")
    }
    if ChooseInt(1, 1, 1) != 1 {
        t.Error("ChooseInt fails 4")
    }
    if ChooseInt(0, 1, 2) != 2 {
        t.Error("ChooseInt fails 5")
    }
    if ChooseInt(1, 2, 1) != 2 {
        t.Error("ChooseInt fails 6")
    }
}

func TestChooseInt8(t *testing.T) {
    if ChooseInt8(0, 0, 0) != 0 {
        t.Error("ChooseInt8 fails 0")
    }
    if ChooseInt8(1, 0, 0) != 0 {
        t.Error("ChooseInt8 fails 1")
    }
    if ChooseInt8(0, 0, 1) != 1 {
        t.Error("ChooseInt8 fails 2")
    }
    if ChooseInt8(1, 1, 0) != 1 {
        t.Error("ChooseInt8 fails 3")
    }
    if ChooseInt8(1, 1, 1) != 1 {
        t.Error("ChooseInt8 fails 4")
    }
    if ChooseInt8(0, 1, 2) != 2 {
        t.Error("ChooseInt8 fails 5")
    }
    if ChooseInt8(1, 2, 1) != 2 {
        t.Error("ChooseInt8 fails 6")
    }
}

func TestChooseInt64(t *testing.T) {
    if ChooseInt64(0, 0, 0) != 0 {
        t.Error("ChooseInt64 fails 0")
    }
    if ChooseInt64(1, 0, 0) != 0 {
        t.Error("ChooseInt64 fails 1")
    }
    if ChooseInt64(0, 0, 1) != 1 {
        t.Error("ChooseInt64 fails 2")
    }
    if ChooseInt64(1, 1, 0) != 1 {
        t.Error("ChooseInt64 fails 3")
    }
    if ChooseInt64(1, 1, 1) != 1 {
        t.Error("ChooseInt64 fails 4")
    }
    if ChooseInt64(0, 1, 2) != 2 {
        t.Error("ChooseInt64 fails 5")
    }
    if ChooseInt64(1, 2, 1) != 2 {
        t.Error("ChooseInt64 fails 6")
    }
}


func TestChooseUint64(t *testing.T) {
    if ChooseUint64(0, 0, 0) != 0 {
        t.Error("ChooseUint64 0 fails")
    }
    if ChooseUint64(0, 1, 0) != 0 {
        t.Error("ChooseUint64 1 fails")
    }
    if ChooseUint64(0, 0xff, 0) != 0 {
        t.Error("ChooseUint64 2 fails")
    }
    if ChooseUint64(0, uint64n, 0) != 0 {
        t.Error("ChooseUint64 3 fails")
    }
    if ChooseUint64(0, 0, 1) != 1 {
        t.Error("ChooseUint64 4 fails")
    }
    if ChooseUint64(0, 1, 1) != 1 {
        t.Error("ChooseUint64 5 fails")
    }
    if ChooseUint64(0, 0xff, 1) != 1 {
        t.Error("ChooseUint64 6 fails")
    }
    if ChooseUint64(0, uint64n, 1) != 1 {
        t.Error("ChooseUint64 7 fails")
    }
    if ChooseUint64(1, 0, 0) != 0 {
        t.Error("ChooseUint64 8 fails")
    }
    if ChooseUint64(1, 1, 0) != 1 {
        t.Error("ChooseUint64 9 fails")
    }
    if ChooseUint64(1, 0xff, 0) != 0xff {
        t.Error("ChooseUint64 a fails")
    }
    if ChooseUint64(1, uint64n, 0) != uint64n {
        t.Error("ChooseUint64 b fails")
    }
    if ChooseUint64(1, 0, 1) != 0 {
        t.Error("ChooseUint64 c fails")
    }
    if ChooseUint64(1, 1, 1) != 1 {
        t.Error("ChooseUint64 d fails")
    }
    if ChooseUint64(1, 0xff, 1) != 0xff {
        t.Error("ChooseUint64 e fails")
    }
    if ChooseUint64(1, uint64n, 1) != uint64n {
        t.Error("ChooseUint64 f fails")
    }
}


/*
func testChooseBytes(t *testing.T, T, F []byte) {
    actual := ChooseBytes(1, T, F)
    if bytes.Compare(actual, T) != 0 {
        t.Error("ChooseBytes fails on true expected=", T," actual=", actual)
    }
    actual = ChooseBytes(0, T, F)
    if bytes.Compare(actual, F) != 0 {
        t.Error("ChooseBytes fails on false expected=", F, " actual=", actual)
    }
}

func TestChooseBytes(t *testing.T) {
    testChooseBytes(t, data1, data1)
    testChooseBytes(t, data1, data2)
    testChooseBytes(t, data2, data1)
    testChooseBytes(t, data2, data2)
    testChooseBytes(t, data2, data3)
    testChooseBytes(t, data3, data2)
    testChooseBytes(t, data3, data3)
}
*/

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

/*
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
*/

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

func make1mib(value byte) []byte {
        result := make([]byte, 1024*1024)
        for i := range result {
            result[i] = value
        }
        return result
}
func make1mibint32(value int32) []int32 {
        result := make([]int32, 1024*1024/4)
        for i := range result {
            result[i] = value
        }
        return result
}

func BenchmarkMSBInt32ForWarmup0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(0)
    }
}

func BenchmarkMSBInt32For0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(0)
    }
}

func BenchmarkMSBInt32ForWarmup1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(1)
    }
}

func BenchmarkMSBInt32For1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(1)
    }
}

func BenchmarkMSBInt32ForWarmupOther(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(0x01234567)
    }
}

func BenchmarkMSBInt32ForOther(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MSBInt32(0x01234567)
    }
}
/*

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

*/
