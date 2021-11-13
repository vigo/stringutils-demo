![Version](https://img.shields.io/badge/version-0.1.0-orange.svg)
![Test](https://github.com/vigo/stringutils-demo/actions/workflows/go.yml/badge.svg)
![Go](https://img.shields.io/badge/go-1.17-black.svg)
[![Documentation](https://godoc.org/github.com/vigo/stringutils-demo?status.svg)](https://pkg.go.dev/github.com/vigo/stringutils-demo)
[![Go Report Card](https://goreportcard.com/badge/github.com/vigo/stringutils-demo)](https://goreportcard.com/report/github.com/vigo/stringutils-demo)

# stringutils

Demo purpose basic golang package. Contains only one function:

`stringutils.Reverse(string)`.

---

## Installation

Go to your project root, where `go.mod` files exists, than;

```bash
$ go install github.com/vigo/stringutils-demo@latest
```

---

## Usage

```go
package main

import (
	"fmt"

	"github.com/vigo/stringutils-demo"
)

func main(){
	reversed := stringutils.Reverse("vigo")
	fmt.Println(reversed)
}
```

---

## Makefile

```bash
$ make help

  make <command>

  commands:

  test    run tests
  bench   run benchmark tests
  doc     run godoc server at 3000 unless PORT env-var is set
```

---

## Contributor(s)

* [Uğur Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/stringutils-demo/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com/vigo/stringutils-demo/blob/main/CODE_OF_CONDUCT.md