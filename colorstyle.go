import (
	"bytes"
	"fmt"
	"strings"
)

type colorStyle struct {
	textStyle *string
	textColor *string
	bgColor   *string
}

func New() *colorStyle {
	return new(colorStyle)
}

func (c *colorStyle) Style() *string {
	return c.textStyle
}

func (c *colorStyle) Color() *string {
	return c.textColor
}
func (c *colorStyle) BgColor() *string {
	return c.bgColor
}

func (c *colorStyle) SetStyle(s Style) *colorStyle {
	c.textStyle = s.String()
	return c
}

func (c *colorStyle) SetColor(s Color) *colorStyle {
	c.textColor = s.String()
	return c

}

func (c *colorStyle) SetBgColor(s BgColor) *colorStyle {
	c.bgColor = s.String()
	return c
}

// [32m heelo [0
// \033[字符的显示方式;字符的颜色;字符的背景颜色m 需要显示的字符 \033[m
func (c *colorStyle) ToString(text string) string {
	var codes []string
	buf := bytes.NewBuffer([]byte("\033["))

	if c.textStyle != nil {
		codes = append(codes, *c.textStyle)

	}
	if c.textColor != nil {
		codes = append(codes, *c.textColor)
	}
	if c.bgColor != nil {
		codes = append(codes, *c.bgColor)
	}
	codeStr := strings.Join(codes, ";")
	buf.WriteString(codeStr)
	buf.WriteString("m")
	buf.WriteString(text)
	buf.WriteString("\033[0m")
	return buf.String()
}

func (c *colorStyle) ColorPrintln(text string) {
	fmt.Println(c.ToString(text))
}
