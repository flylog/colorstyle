package colorStyle

import (
	"bytes"
	"fmt"
	"strings"
)

type CSS struct {
	textStyle *string
	textColor *string
	bgColor   *string
}

func New() *CSS {
	return new(CSS)
}

func (c *CSS) setStyle(s Style) *CSS {
	c.textStyle = s.ptrString()
	return c
}
func (c *CSS) setColor(s Color) *CSS {
	c.textColor = s.ptrString()
	return c
}

func (c *CSS) setBgColor(s BgColor) *CSS {
	c.bgColor = s.ptrString()
	return c
}

// \033[字符的显示方式;字符的颜色;字符的背景颜色m 需要显示的字符 \033[m
func (c *CSS) toString() (start, end string) {
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
	return buf.String(), "\033[0m"
}

func (c *CSS) Println(text ...interface{}) {
	start, end := c.toString()
	fmt.Println(start + fmt.Sprint(text...) + end)
}

func (c *CSS) Printf(format string, text ...interface{}) {
	start, end := c.toString()
	fmt.Printf(format, start+fmt.Sprint(text...)+end)
}

func (c *CSS) Sprint(text ...interface{}) string {
	start, end := c.toString()
	return fmt.Sprint(start + fmt.Sprint(text...) + end)

}

func (c *CSS) Sprintf(format string, text ...interface{}) string {
	start, end := c.toString()
	return fmt.Sprintf(format, start+fmt.Sprint(text...)+end)
}

func Black(text ...interface{}) string {
	c := New()
	return c.setColor(FgBlack).Sprint(text...)
}

func Red(text ...interface{}) string {
	c := New()
	return c.setColor(FgRed).Sprint(text...)
}

func Green(text ...interface{}) string {
	c := New()
	return c.setColor(FgGreen).Sprint(text...)
}

func Yellow(text ...interface{}) string {
	c := New()
	return c.setColor(FgYellow).Sprint(text...)
}

func DarkBlue(text ...interface{}) string {
	c := New()
	return c.setColor(FgDarkBlue).Sprint(text...)
}

func Purple(text ...interface{}) string {
	c := New()
	return c.setColor(FgDarkBlue).Sprint(text...)
}

func LightBlue(text ...interface{}) string {
	c := New()
	return c.setColor(FgLightBlue).Sprint(text...)
}

func White(text ...interface{}) string {
	c := New()
	return c.setColor(FgWhite).Sprint(text...)
}
