package colorstyle_test

import (
	"fmt"

	"github.com/flylog/colorstyle"
)

func Example() {
	text := colorstyle.Green("green")
	fmt.Printf("a %s text\n", text)

	text = colorstyle.Blue("Blue")
	fmt.Printf("a %s text\n", text)

	text = colorstyle.New().ColorRed().Sprint("red")
	fmt.Println("a", text, "text")

	colorstyle.New().StyleItalic().ColorRed().BgYellow().Printf("a italic red bgYellow text: %s\n", "Hello 世界!")
	colorstyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	colorstyle.New().StyleItalic().Printf("a italic text: %s\n", "Hello 世界!")
	colorstyle.New().ColorMagenta().Printf("a magenta text: %s\n", "Hello 世界!")
	colorstyle.New().BgCyan().Printf("a  background color cyan text: %s\n", "Hello 世界!")
	colorstyle.New().BgCyan().Println("a background color cyan text")

	css := colorstyle.New()
	css.StyleStrikethrough().Println("删除线文本")
	css.StyleUnderline().Println("下划线文本")
	css.StyleReverse().Println("反显文本")
}

func ExampleGreen() {
	text := colorstyle.Green("green")
	fmt.Printf("a %s text\n", text)
}

func ExampleBlue() {
	text := colorstyle.Blue("Blue")
	fmt.Printf("a %s text\n", text)
}

func ExampleNew() {
	css := colorstyle.New()
	css.StyleItalic().ColorGreen().BgYellow()
	css.Printf("a italic green bgYellow text: %s\n", "Hello 世界!")
}

func ExampleCSS_StyleBold() {
	colorstyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	// or
	css := colorstyle.New().StyleBold()
	css.Printf("a bold text: %s\n", "Hello 世界!")

}

func ExampleCSS_ColorRed() {
	colorstyle.New().ColorRed().Printf("a bold text: %s\n", "Hello 世界!")

}

func ExampleCSS_Printf() {
	colorstyle.New().ColorRed().Printf("a bold text: %s\n", "Hello 世界!")
}

func ExampleCSS_Println() {
	colorstyle.New().BgBlue().Println("a background color blue text")
}
