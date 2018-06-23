# go-moji
[![Build Status](https://circleci.com/gh/ktnyt/go-moji.svg?style=shield&circle-token==7da9cb901d095995e930651e7298c8ab233a0c85)](https://circleci.com/gh/ktnyt/go-moji)
[![Go Report Card](https://goreportcard.com/badge/github.com/ktnyt/go-moji)](https://goreportcard.com/report/github.com/ktnyt/go-moji)
[![GoDoc](http://godoc.org/github.com/ktnyt/go-moji?status.svg)](http://godoc.org/github.com/ktnyt/go-moji)

This package provides a Go interface for converting between Zenkaku (全角 i.e. full-width) and Hankaku (半角 i.e. half-width) characters (mostly for Japanese). The library has been largely influenced by [niwaringo/moji](https://github.com/niwaringo/moji) the JavaScript implementation.

For detailed information of the API, see the [documents](https://godoc.org/github.com/ktnyt/go-moji).

## Installation
Use `go get`:
```sh
$ go get github.com/ktnyt/go-moji
```

## Requirements
This package has only been tested on Go >= 1.8. Beware when using lower versions.

## Example
```go
package main

import (
	"fmt"

	"github.com/ktnyt/moji"
)

func main() {
	s := "ＡＢＣ ABC　あがぱ　アガパ　ｱｶﾞﾊﾟ"

	// Convert Zenkaku Eisuu to Hankaku Eisuu
	fmt.Println(moji.Convert(s, moji.ZE, moji.HE))

	// Convert Hankaku Eisuu to Zenkaku Eisuu
	fmt.Println(moji.Convert(s, moji.HE, moji.ZE))

	// Convert HiraGana to KataKana
	fmt.Println(moji.Convert(s, moji.HG, moji.KK))

	// Convert KataKana to HiraGana
	fmt.Println(moji.Convert(s, moji.KK, moji.HG))

	// Convert Zenkaku Katakana to Hankaku Katakana
	fmt.Println(moji.Convert(s, moji.ZK, moji.HK))

	// Convert Hankaku Katakana to Zenkaku Katakana
	fmt.Println(moji.Convert(s, moji.HK, moji.ZK))

	// Convert Zenkaku Space to Hankaku Space
	fmt.Println(moji.Convert(s, moji.ZS, moji.HS))

	// Convert Hankaku Space to Zenkaku Space
	fmt.Println(moji.Convert(s, moji.HS, moji.ZS))
}
```

## Copyright
Copyright (C) 2018 by Kotone Itaya <kotone@sfc.keio.ac.jp>

go-moji is released under the terms of the MIT License.
See [LICENSE](https://github.com/ktnyt/go-moji/blob/master/LICENSE) for details.
