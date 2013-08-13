package securecompare

import (
    "encoding/hex"
    "testing"
)

func testCompare(t *testing.T, a, b []byte, expected bool) {
    if expected != Compare(a, b) {
        t.Errorf("Should have failed")
        return
    }
}

func testCompareSuccess(t *testing.T, a, b []byte) {
    testCompare(t, a, b, true)
}

func testCompareFail(t *testing.T, a, b []byte) {
    testCompare(t, a, b, false)
}

func TestCompare(t *testing.T) {
    data1, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data2, _ := hex.DecodeString("01bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data3, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50de")
    data4, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50")
    data5, _ := hex.DecodeString("")

    testCompareFail(t, data1, data2)
    testCompareFail(t, data1, data3)
    testCompareFail(t, data1, data4)
    testCompareFail(t, data1, data5)
    testCompareSuccess(t, data1, data1)
    testCompareSuccess(t, data5, data5)
}


var (
    x = make([]byte, 1024*1024)
    y = make([]byte, 1024*1024)
)

func BenchmarkCompare1MiB(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Compare(x, y)
    }
}
