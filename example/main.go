package main

import (
	"fmt"

	"github.com/flylog/colorStyle"
)

func main() {
	// colorStyle.New().ColorPrint("Hello 世界!")
	colorStyle.New(colorStyle.Red).SetTextStyle(colorStyle.Italic).SetBgColor(colorStyle.BgWhite).ColorPrint("Hello")
	fmt.Println()
	colorStyle.New(colorStyle.Red).SetBgColor(colorStyle.BgBlue).ColorPrint("Hello")
	colorStyle.New(colorStyle.Red).ColorPrint("Hello")
	fmt.Println(string([]byte{27, 91, 51, 50, 109}), "heelo", string([]byte{27, 91, 48, 109}))
}
