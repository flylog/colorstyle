package main

import (
	"fmt"

	"github.com/flylog/colorStyle"
)

func main() {
	text := colorStyle.Green("green")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.BrightBlue("BrightBlue")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.New().ColorRed().Sprint("red")
	fmt.Println("a", text, "text")

	colorStyle.New().StyleItalic().ColorRed().BgYellow().Printf("a italic red bgYellow text: %s\n", "Hello 世界!")
	colorStyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	colorStyle.New().StyleItalic().Printf("a italic text: %s\n", "Hello 世界!")
	colorStyle.New().ColorBrightMagenta().Printf("a magenta text: %s\n", "Hello 世界!")
	colorStyle.New().BgCyan().Printf("a bgLightBlue text: %s\n", "Hello 世界!")
	colorStyle.New().BgCyan().Println("a bgLightBlue text")

	css := colorStyle.New()
	css.StyleUnderline().Println("下划线文本")
	css.StyleReverse().Println("反显文本")

}
