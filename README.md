# Go Easy Style

<p align="center">
	<a href="https://github.com/leo-alvarenga/go-easy-style/actions/workflows/integrate.yml">
		<img src="https://github.com/leo-alvarenga/go-easy-style/actions/workflows/integrate.yml/badge.svg" alt="Integration Tests">
	</a>
	<a href="https://goreportcard.com/report/github.com/leo-alvarenga/go-easy-style">
		<img src="https://goreportcard.com/badge/github.com/leo-alvarenga/go-easy-style" alt="Go Report Card">
	</a>
	<a href="https://github.com/leo-alvarenga/go-easy-style/blob/main/LICENSE.txt">
		<img src="https://img.shields.io/github/license/leo-alvarenga/go-easy-style" alt="License">
	</a>	
	<img src="https://img.shields.io/github/go-mod/go-version/leo-alvarenga/go-easy-style" alt="License">
	<a href="https://gitHub.com/leo-alvarenga/go-easy-style/releases/">
		<img src="https://img.shields.io/github/v/tag/leo-alvarenga/go-easy-style?sort=semver" alt="GitHub release">
	</a>	
	<a href="https://pkg.go.dev/github.com/leo-alvarenga/go-easy-style">
		<img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs">
	</a>
</p>

A dependency free package to help with "styling" `stdout` text with ANSI escape sequences.

## Installation

```bash
    go get github.com/leo-alvarenga/go-easy-style@v1
```

## Usage

### `TextStyle`
`TextStyle` is a struct as (seen bellow) that acts as the "entrypoint" to use the package.

```go
type TextStyle struct {
	Text       string
	Background string
	Format     []string
	ANSI       string
}

func (t *TextStyle) New(txt, background string, styles ...string) {...}
func (t *TextStyle) Style(s string) string {...}
func (t *TextStyle) ShowWithStyle(s string){...}
```

It also exposes the three (3) methods:

- `New`: Initializes a `TextStyle` instance using defined values. Every time you call `New` for the same instance, the style is overwritten. Invalid values are simply ignored
- `Style`: Returns the styled version of the `string` passed as argument, ready to be printed to `stdout` without interfering with the rest of the content
- `ShowWithStyle`: A shorthand to style and print the `string` passed as argument to `stdout` 

### Other
Other exposed functions:

```go
func ShowAsError(title, msg string) {...}
```

- `ShowAsError`: A shorthand to style the `string` passed as argument with a black background, red text and bold letters and print it to `stdout` 

## Example

```go
package main

import (
    "fmt"
    "github.com/leo-alvarenga/go-easy-style"
)

func main() {
    var num int

    style := new(gostyle.TextStyle)

    style.New("black", "white", "bold", "italic", "undescore")
    style.ShowWithStyle("Hello, colorfull world!")
    
    style.New("white", "blue", "blink")

    println(style.Style("Type in a number:"))
    fmt.Scanf("%d", &num)

    if num < 0 {
        gostyle.ShowAsError("Too low!", "The value you typed is negative")
    }
}
```