package main

import (
	"fmt"
)

func main() {
	// colorStyle.New().ColorPrint("Hello 世界!")
	css.New(css.Red).SetStyle(css.Italic).SetBgColor(css.BgWhite).ColorPrint("Hello")
	fmt.Println()
	css.New(css.Red).SetBgColor(css.BgBlue).ColorPrint("Hello")
	css.New(css.Red).ColorPrint("Hello")
	fmt.Println(string([]byte{27, 91, 51, 50, 109}), "heelo", string([]byte{27, 91, 48, 109}))
}
