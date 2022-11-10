# go-strip-json-comments

#### Note:

This is an exact port of Sindresorhus [strip-json-comments](http://githib.com/sindresorhus/strip-json-comments) to
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
        Whitespace: false,
		TrailingCommas: true,
    }
	
	fmt.Println(stripjsoncomments.StripWithOptions(json, &options))
}
```


## API