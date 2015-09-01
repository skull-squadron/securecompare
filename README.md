# securecompare for Go

### Don't use `bytes.Equal()` for comparing hashes!
### BAD! (and evil naughty Zoot)

Recommended alternative: 

```go
import "crypto/subtle"

xyEqual := subtle.ConstantTimeCompare(x, y) == 1
```

This leads to side-channel attacks.

### All securecompare operations use constant-time constructs (e.g., math).

#### Note: minor differences per anecdotal `go test -bench` are due to timing inaccuracies.  Across many samples they are equal.

```shell
PASS
BenchmarkMSBInt32ForWarmup0 500000000            2.97 ns/op
BenchmarkMSBInt32For0   500000000            3.44 ns/op
BenchmarkMSBInt32ForWarmup1 500000000            2.99 ns/op
BenchmarkMSBInt32For1   500000000            3.44 ns/op
BenchmarkMSBInt32ForWarmupOther 500000000            2.98 ns/op
BenchmarkMSBInt32ForOther   500000000            3.52 ns/op
BenchmarkBoolToIntWarmUp    200000000            8.08 ns/op
BenchmarkBoolToIntFalse 200000000            8.14 ns/op
BenchmarkBoolToIntTrue  200000000            8.11 ns/op
BenchmarkBoolToInt64WarmUp  200000000            8.18 ns/op
BenchmarkBoolToInt64False   200000000            9.25 ns/op
BenchmarkBoolToInt64True    200000000            8.98 ns/op
BenchmarkBoolToInt8WarmUp   200000000            8.13 ns/op
BenchmarkBoolToInt8False    200000000            8.10 ns/op
BenchmarkBoolToInt8True 200000000            8.31 ns/op
BenchmarkDupBitToInt64  500000000            4.28 ns/op
BenchmarkDupBitToInt    500000000            6.13 ns/op
BenchmarkDupBitToInt8   1000000000           2.31 ns/op
BenchmarkEqual1MiB0     1000       2869088 ns/op
BenchmarkEqual1MiB1     1000       2845156 ns/op
BenchmarkEqual1MiB2     1000       2538090 ns/op
BenchmarkEqual1MiB3     1000       2528940 ns/op
BenchmarkEqual1MiB4     1000       2532726 ns/op
BenchmarkCompare1MiB0          1    1481007150 ns/op
BenchmarkCompare1MiB1          1    1360861457 ns/op
BenchmarkCompare1MiB2          1    1241465379 ns/op
BenchmarkCompare1MiB3          1    1361820448 ns/op
BenchmarkCompare1MiB4          1    1188699854 ns/op
ok      [redacted]/go/securecompare  64.388s
```

## Please verify the generated assembly by running

    go build -gcflags '-S' securecompare.go

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
