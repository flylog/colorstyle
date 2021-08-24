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
