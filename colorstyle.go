package colorStyle

import (
	"bytes"
	"fmt"
)

type colorStyle struct {
	textStyle string
	textColor string
	bgColor   string
}

func (c *colorStyle) TextStyle() string {
	return c.textStyle
}

func (c *colorStyle) TextColor() string {
	return c.textColor
}
func (c *colorStyle) BgColor() string {
	return c.bgColor
}

func (c *colorStyle) SetTextStyle(s Style) *colorStyle {
	c.textStyle = s.String()
	return c
}

func New(s Color) *colorStyle {
	c := new(colorStyle)
	c.textColor = s.String()
	return c
}

func (c *colorStyle) SetBgColor(s BgColor) *colorStyle {
	c.bgColor = s.String()
	return c
}

// [32m heelo [0
// \033[字符的显示方式;字符的颜色;字符的背景颜色m 需要显示的字符 \033[m
func (c *colorStyle) ColorPrint(text string) {
	buf := bytes.NewBuffer([]byte("\033["))
	if c.TextStyle() != "0" {

		buf.WriteString(c.textStyle)
		buf.WriteString(";")
	}
	if c.TextColor() != "0" {

		buf.WriteString(c.textColor)
		buf.WriteString(";")
	}
	if c.BgColor() != "0" {

		buf.WriteString(c.bgColor)
		buf.WriteString("m")
	}
	b := buf.String()
	fmt.Println(b, text, "\033[0m")
}
