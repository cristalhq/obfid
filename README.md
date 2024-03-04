# obfid

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

Obfuscating ID based on modular multiplicative inverse in Go.

## Rationale

The simplest and fastest way to encode and expose to public an integer primary key from a database can be achieved by obfuscating it with [multiplicative inverse
](https://en.wikipedia.org/wiki/Multiplicative_inverse) (aka Knuth's Hashing Algorithm).

## Features

* Simple API.
* Dependency-free.
* Clean and tested code.

See [these docs][pkg-url] or [GUIDE.md](https://github.com/cristalhq/obfid/blob/main/GUIDE.md) for more details.

## Install

Go version 1.17+

```
go get github.com/cristalhq/obfid
```

## Example

```go
prime := uint64(32_452_867)
random := uint64(123_456)
offset := uint64(1_000_000)
bits := 30

generator, err := obfid.NewGenerator(prime, random, offset, bits)
if err != nil {
	panic(err)
}

for num := uint64(0); num < 10; num++ {
	enc := generator.Encode(num)
	dec := generator.Decode(enc)

	fmt.Printf("%d => %d => %d\n", num, enc, dec)

	if num != dec {
		fmt.Printf("oops must be equal: %d != %d\n", num, dec)
	}
}

// Output:
// 0 => 1123456 => 0
// 1 => 33428867 => 1
// 2 => 65979014 => 2
// 3 => 98284425 => 3
// 4 => 130836620 => 4
// 5 => 163142031 => 5
// 6 => 195692178 => 6
// 7 => 228259733 => 7
// 8 => 260549784 => 8
// 9 => 293117339 => 9
```

See examples: [example_test.go](https://github.com/cristalhq/obfid/blob/main/example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/obfid/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/obfid/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/obfid
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/obfid
[version-img]: https://img.shields.io/github/v/release/cristalhq/obfid
[version-url]: https://github.com/cristalhq/obfid/releases
