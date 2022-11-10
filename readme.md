# go-strip-json-comments

[github](https://github.com/trapcodeio/go-strip-json-comments) | [pkg.go.dev](https://pkg.go.dev/github.com/trapcodeio/go-strip-json-comments)

#### Note:

This is an exact port of Sindresorhus Nodejs [strip-json-comments](https://github.com/sindresorhus/strip-json-comments) to
GoLang.



> Strip comments from JSON. Lets you use comments in your JSON files!

This is now possible:

```js
{
    // Rainbows
    "unicorn"
: /* ❤ */
    "cake"
}
```

It will replace single-line comments `//` and multi-line comments `/**/` with whitespace. This allows JSON error
positions to remain as close as possible to the original source.

Also available as a [Gulp](https://github.com/sindresorhus/gulp-strip-json-comments)
/[Grunt](https://github.com/sindresorhus/grunt-strip-json-comments)
/[Broccoli](https://github.com/sindresorhus/broccoli-strip-json-comments) plugin.

## Install

```sh
go get github.com/trapcodeio/go-strip-json-comments
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/trapcodeio/go-strip-json-comments"
)

func main() {
	json := `{
        // Rainbows
        "unicorn": /* ❤ */ "cake"
    }`

	fmt.Println(stripjsoncomments.Strip(json))
	//=>{ "unicorn": "cake" }

	// with options
	options := stripjsoncomments.Options{
		Whitespace:     false,
		TrailingCommas: true,
	}

	fmt.Println(stripjsoncomments.StripWithOptions(json, &options))
}
```

### Test
Clone repo and run
```sh
go test -v
```
you can also view them in this repo actions tab.

### Benchmark

```sh
go test -bench=. -count 5 -run=^#
```

```
goos: darwin
goarch: amd64
pkg: github.com/trapcodeio/go-strip-json-comments
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkWithFiles/strip_JSON_comments-16                    332           3304452 ns/op
BenchmarkWithFiles/strip_JSON_comments-16                    408           3759578 ns/op
BenchmarkWithFiles/strip_JSON_comments-16                    289           3634032 ns/op
BenchmarkWithFiles/strip_JSON_comments-16                    297           3826366 ns/op
BenchmarkWithFiles/strip_JSON_comments-16                    338           3652408 ns/op
BenchmarkWithFiles/strip_JSON_comments_without_whitespace-16                 621           1747448 ns/op
BenchmarkWithFiles/strip_JSON_comments_without_whitespace-16                 639           1899337 ns/op
BenchmarkWithFiles/strip_JSON_comments_without_whitespace-16                 652           1891341 ns/op
BenchmarkWithFiles/strip_JSON_comments_without_whitespace-16                 766           1799857 ns/op
BenchmarkWithFiles/strip_JSON_comments_without_whitespace-16                 699           1867128 ns/op
BenchmarkWithFiles/strip_Big_JSON_comments-16                                  9         131861237 ns/op
BenchmarkWithFiles/strip_Big_JSON_comments-16                                  8         138128030 ns/op
BenchmarkWithFiles/strip_Big_JSON_comments-16                                  9         140734501 ns/op
BenchmarkWithFiles/strip_Big_JSON_comments-16                                  8         126000649 ns/op
BenchmarkWithFiles/strip_Big_JSON_comments-16                                  8         133652308 ns/op
PASS
ok      github.com/trapcodeio/go-strip-json-comments    24.843s
```

