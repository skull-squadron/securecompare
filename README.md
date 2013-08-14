# securecompare for Go

### Don't use `bytes.Equal()` for comparing hashes!
### BAD!

This leads to side-channel attacks.

### All securecompare operations use constant-time constructs (e.g., math).

## Download

    go get github.com/steakknife/securecompare

## Usage

    ```go
    import "github.com/steakknife/securecompare"

    ...
    securecompare.Equal(a, b)
    ```

## License

MIT

## Patches welcome!
