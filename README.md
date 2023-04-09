# natsort

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Natural sorting in Go, see [Wikipedia](https://en.wikipedia.org/wiki/Natural_sort_order).

## Features

* Fast.
* Simple API.
* Dependency-free.

## Install

Go version 1.17+

```
go get github.com/cristalhq/natsort
```

## Example

```go
files := []string{"img12.png", "img10.png", "img2.png", "img1.png"}

fmt.Println("Lexicographically:")

sort.Strings(files)
for _, f := range files {
	fmt.Println(f)
}

fmt.Println("\nNaturally:")

natsort.Sort(files)
for _, f := range files {
	fmt.Println(f)
}

// Output:
// Lexicographically:
// img1.png
// img10.png
// img12.png
// img2.png
//
// Naturally:
// img1.png
// img2.png
// img10.png
// img12.png
```

See examples: [example_test.go](example_test.go).

## Documentation

See [these docs][pkg-url] for more details.

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/natsort/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/natsort/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/natsort
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/natsort
[reportcard-img]: https://goreportcard.com/badge/cristalhq/natsort
[reportcard-url]: https://goreportcard.com/report/cristalhq/natsort
[coverage-img]: https://codecov.io/gh/cristalhq/natsort/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/natsort
[version-img]: https://img.shields.io/github/v/release/cristalhq/natsort
[version-url]: https://github.com/cristalhq/natsort/releases
