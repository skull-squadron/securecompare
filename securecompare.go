package securecompare

func Compare(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    result := byte(0)

    for i := range a {
        result |= a[i] ^ b[i]
    }

    return (result == 0)
}
