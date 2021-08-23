package main

import (
	"fmt"

	"github.com/flylog/colorStyle"
)

func main() {
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
