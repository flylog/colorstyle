package colorStyle_test

import (
	"fmt"

	"github.com/flylog/colorStyle"
)

func Example() {
	text := colorStyle.Green("green")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.DarkBlue("Blue")
	fmt.Printf("a %s text\n", text)

	text = colorStyle.New().ColorRed().Sprint("red")
	fmt.Println("a", text, "text")

	colorStyle.New().StyleItalic().ColorRed().BgYellow().Printf("a italic red bgYellow text: %s\n", "Hello 世界!")
	colorStyle.New().StyleBold().Printf("a bold text: %s\n", "Hello 世界!")
	colorStyle.New().StyleItalic().Printf("a italic text: %s\n", "Hello 世界!")
	colorStyle.New().ColorPurple().Printf("a purple text: %s\n", "Hello 世界!")
	colorStyle.New().BgLightBlue().Printf("a bgLightBlue text: %s\n", "Hello 世界!")
	colorStyle.New().BgLightBlue().Println("a bgLightBlue text")
}

func ExampleGreen() {
	text := colorStyle.Green("green")
	fmt.Printf("a %s text\n", text)
}

func ExampleDarkBlue() {
	text := colorStyle.DarkBlue("Blue")
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
	colorStyle.New().BgLightBlue().Println("a bgLightBlue text")
}
