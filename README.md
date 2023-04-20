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
generator, err := NewGenerator(32452867, 123)
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
// 0 => 123 => 0
// 1 => 32452984 => 1
// 2 => 64905853 => 2
// 3 => 97358706 => 3
// 4 => 129811575 => 4
// 5 => 162264436 => 5
// 6 => 194717289 => 6
// 7 => 227170158 => 7
// 8 => 259623011 => 8
// 9 => 292075872 => 9
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
