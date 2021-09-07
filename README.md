# colorStyle

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/flylog/colorstyle?color=peru)](https://github.com/flylog/colorstyle/releases/latest)
[![Released API docs](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/flylog/colorstyle)
[![Build](https://img.shields.io/github/workflow/status/flylog/colorstyle/Go.svg)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/flylog/colorstyle?color=pink)](https://goreportcard.com/report/github.com/flylog/colorstyle)
[![Lines of code](https://img.shields.io/tokei/lines/github/flylog/colorstyle.svg?color=beige)](#)
[![Downloads of releases](https://img.shields.io/github/downloads/flylog/colorstyle/total.svg?color=lavender)](https://github.com/flylog/colorstyle/releases/latest)
[![Languages](https://img.shields.io/github/languages/top/flylog/colorstyle.svg?color=yellow)](#)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/flylog/colorstyle)](#)
[![GPL3 licensed](https://img.shields.io/github/license/flylog/colorstyle.svg)](./LICENSE)

ColorStyle is a library of styles for command-line text.
Used to modify the style of text for standard output to the terminal interface, you can change the foreground colour of the text, the background colour, add underline and bold, etc.

ColorStyle 是一个用于命令行文本的样式库。
用于标准输出到终端界面的文本的样式修改，可以修改文本前景色，背景色，增加下划线和加粗显等。

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/flylog/colorstyle)

# Install
    
    go get -u github.com/flylog/colorstyle@latest


# Example

```
package main

import (
	"fmt"

	"github.com/flylog/colorstyle"
)

func main() {
	text := colorstyle.Green("green")
	fmt.Printf("a %s text\n", text)

	text = colorstyle.BrightBlue("BrightBlue")
	fmt.Printf("a %s text\n", text)

	text = colorstyle.New().ColorRed().Sprint("red")
	fmt.Println("a", text, "text")
	colorstyle.New().StyleItalic().ColorRed().BgYellow().Printf("a italic red bgYellow text: %s\n", "Hello 世界!")
	colorstyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	colorstyle.New().StyleItalic().Printf("a italic text: %s\n", "Hello 世界!")
	colorstyle.New().ColorBrightMagenta().Printf("a magenta text: %s\n", "Hello 世界!")
	colorstyle.New().BgCyan().Printf("a bgCyan text: %s\n", "Hello 世界!")
	colorstyle.New().BgCyan().Println("a bgCyan text")

	css := colorstyle.New()
	css.StyleUnderline().Println("下划线文本")
	css.StyleReverse().Println("反显文本")

}

```
output:

![](example/output.jpg)
# Implementation of other languages

- [x] write in Rust: https://github.com/code-translation/colorstyle
