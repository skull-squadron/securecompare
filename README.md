# securecompare for Go

### Don't use `bytes.Equal()` for comparing hashes!
### BAD!

This leads to side-channel attacks.

### All securecompare operations use constant-time constructs (e.g., math).

#### Note: minor differences per anecdotal `go test -bench` are due to timing inaccuracies.  Across many samples they are equal.

```shell
PASS
BenchmarkBoolToIntWarmUp     100000000           12.2 ns/op
BenchmarkBoolToIntFalse      500000000            6.18 ns/op
BenchmarkBoolToIntTrue       500000000            6.00 ns/op
BenchmarkBoolToInt64WarmUp   100000000           11.3 ns/op
BenchmarkBoolToInt64False    100000000           14.3 ns/op
BenchmarkBoolToInt64True     200000000            5.97 ns/op
BenchmarkBoolToInt8WarmUp    500000000            6.08 ns/op
BenchmarkBoolToInt8False     500000000            5.88 ns/op
BenchmarkBoolToInt8True      500000000            5.90 ns/op
BenchmarkDupBitToInt64       500000000            5.93 ns/op
BenchmarkDupBitToInt         500000000            5.94 ns/op
BenchmarkDupBitToInt8       1000000000            2.10 ns/op
BenchmarkEqual1MiB0               1000      2511385 ns/op
BenchmarkEqual1MiB1               1000      2515872 ns/op
BenchmarkEqual1MiB2               1000      2522780 ns/op
BenchmarkEqual1MiB3               1000      2584363 ns/op
BenchmarkEqual1MiB4               1000      2536941 ns/op
BenchmarkCompare1MiB0               50     49718942 ns/op
BenchmarkCompare1MiB1               50     50606172 ns/op
BenchmarkCompare1MiB2               50     50603049 ns/op
BenchmarkCompare1MiB3               50     50142368 ns/op
BenchmarkCompare1MiB4               50     49548024 ns/op
ok
```

## Download

    go get github.com/steakknife/securecompare

## Usage

```go
    
import "github.com/steakknife/securecompare"

// ...

if securecompare.Equal(a, b) {
    // ..
}
    
```

## License

MIT

## Patches welcome!
