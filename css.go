package colorStyle

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	ANSI_SET   = "\033["   // ANSI转义码，设置文本样式的开头
	ANSI_END   = "m"       // 设置结束字符
	ANSI_RESET = "\033[0m" // ANSI转义码,重置设置到常规
)

type CSS struct {
	textStyle *string
	textColor *string
	bgColor   *string
}

// 新建一个文本样式对象
//
// Create a new text style object
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
	buf := bytes.NewBuffer([]byte(ANSI_SET))

	if c.textStyle != nil {
		codes = append(codes, *c.textStyle)

	}
	if c.textColor != nil {
		codes = append(codes, *c.textColor)
	}
	if c.bgColor != nil {
		codes = append(codes, *c.bgColor)
	}
	codeStr := strings.Join(codes, ";") //多种设置用分号连接
	buf.WriteString(codeStr)
	buf.WriteString(ANSI_END)
	return buf.String(), ANSI_RESET
}

// like fmt.Println
func (c *CSS) Println(text ...interface{}) {
	start, end := c.toString()
	fmt.Println(start + fmt.Sprint(text...) + end)
}

// like fmt.Printf
func (c *CSS) Printf(format string, text ...interface{}) {
	start, end := c.toString()
	fmt.Printf(format, start+fmt.Sprint(text...)+end)
}

// like fmt.Sprint
func (c *CSS) Sprint(text ...interface{}) string {
	start, end := c.toString()
	return fmt.Sprint(start + fmt.Sprint(text...) + end)

}

// like fmt.Sprintf
func (c *CSS) Sprintf(format string, text ...interface{}) string {
	start, end := c.toString()
	return fmt.Sprintf(format, start+fmt.Sprint(text...)+end)
}

// 生成黑色的文本
//
// Generate black text
func Black(text ...interface{}) string {
	c := New()
	return c.setColor(FgBlack).Sprint(text...)
}

// 生成红色的文本
//
// Generate red text
func Red(text ...interface{}) string {
	c := New()
	return c.setColor(FgRed).Sprint(text...)
}

// 生成绿色的文本
//
// Generate green text
func Green(text ...interface{}) string {
	c := New()
	return c.setColor(FgGreen).Sprint(text...)
}

// 生成黄色的文本
//
// Generate yellow text
func Yellow(text ...interface{}) string {
	c := New()
	return c.setColor(FgYellow).Sprint(text...)
}

// 生成蓝色的文本
//
// Generate blue text
func Blue(text ...interface{}) string {
	c := New()
	return c.setColor(FgBlue).Sprint(text...)
}

// 生成品红色的文本
//
// Generate magenta text
func Magenta(text ...interface{}) string {
	c := New()
	return c.setColor(FgMagenta).Sprint(text...)
}

// 生成青色的文本
//
// Generate cyan text
func Cyan(text ...interface{}) string {
	c := New()
	return c.setColor(FgCyan).Sprint(text...)
}

// 生成白色的文本
//
// Generate white text
func White(text ...interface{}) string {
	c := New()
	return c.setColor(FgWhite).Sprint(text...)
}

// 生成灰色的文本
//
// Generate gray text
func Gray(text ...interface{}) string {
	c := New()
	return c.setColor(FgGray).Sprint(text...)
}

// 生成亮红色的文本
//
// Generate bright red text
func BrightRed(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightRed).Sprint(text...)
}

// 生成亮黄色的文本
//
// Generate bright yellow text
func BrightYellow(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightYellow).Sprint(text...)
}

// 生成亮蓝的文本
//
// Generate bright blue text
func BrightBlue(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightBlue).Sprint(text...)
}

// 生成亮品红的文本
//
// Generate bright magenta text
func BrightMagenta(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightMagenta).Sprint(text...)
}

// 生成亮青色的文本
//
// Generate bright cyan text
func BrightCyan(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightCyan).Sprint(text...)
}

// 生成亮色的文本
//
// Generate bright white text
func BrightWhite(text ...interface{}) string {
	c := New()
	return c.setColor(FgBrightWhite).Sprint(text...)
}
