package colorStyle_test

import (
	"fmt"

	"github.com/flylog/colorStyle"
)

func Example() {
	text := colorStyle.Green("green")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.Blue("Blue")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.New().ColorRed().Sprint("red")
	fmt.Println("a", text, "text")

	colorStyle.New().StyleItalic().ColorRed().BgYellow().Printf("a italic red bgYellow text: %s\n", "Hello 世界!")
	colorStyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	colorStyle.New().StyleItalic().Printf("a italic text: %s\n", "Hello 世界!")
	colorStyle.New().ColorMagenta().Printf("a magenta text: %s\n", "Hello 世界!")
	colorStyle.New().BgCyan().Printf("a  background color cyan text: %s\n", "Hello 世界!")
	colorStyle.New().BgCyan().Println("a background color cyan text")

	css := colorStyle.New()
	css.StyleStrikethrough().Println("删除线文本")
	css.StyleUnderline().Println("下划线文本")
	css.StyleReverse().Println("反显文本")
}

func ExampleGreen() {
	text := colorStyle.Green("green")
	fmt.Printf("a %s text\n", text)
}

func ExampleBlue() {
	text := colorStyle.Blue("Blue")
	fmt.Printf("a %s text\n", text)
}

func ExampleNew() {
	css := colorStyle.New()
	css.StyleItalic().ColorGreen().BgYellow()
	css.Printf("a italic green bgYellow text: %s\n", "Hello 世界!")
}

func ExampleCSS_StyleBold() {
	colorStyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	// or
	css := colorStyle.New().StyleBold()
	css.Printf("a bold text: %s\n", "Hello 世界!")

}

func ExampleCSS_ColorRed() {
	colorStyle.New().ColorRed().Printf("a bold text: %s\n", "Hello 世界!")

}

func ExampleCSS_Printf() {
	colorStyle.New().ColorRed().Printf("a bold text: %s\n", "Hello 世界!")
}

func ExampleCSS_Println() {
	colorStyle.New().BgBlue().Println("a background color blue text")
}
