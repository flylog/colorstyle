package colorstyle

import "strconv"

type Color int

func (c Color) ptrString() *string {
	str := strconv.Itoa(int(c))
	return &str
}

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	FgGray Color = iota + 82
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWhite
)

// 设置前景色为黑色
//
// Set foreground colour to black
func (c *CSS) ColorBlack() *CSS {
	return c.setColor(FgBlack)
}

// 设置前景色为红色
//
// Set foreground colour to red
func (c *CSS) ColorRed() *CSS {
	return c.setColor(FgRed)
}

// 设置前景色为绿色
//
// Set foreground colour to green
func (c *CSS) ColorGreen() *CSS {
	return c.setColor(FgGreen)
}

// 设置前景色为黄色
//
// Set foreground colour to yellow
func (c *CSS) ColorYellow() *CSS {
	return c.setColor(FgYellow)
}

// 设置前景色为蓝色
//
// Set foreground colour to blue
func (c *CSS) ColorBule() *CSS {
	return c.setColor(FgBlue)
}

// 设置前景色为品红色
//
// Set foreground colour to magenta
func (c *CSS) ColorMagenta() *CSS {
	return c.setColor(FgMagenta)
}

// 设置前景为青色
//
// Set foreground colour to cyan
func (c *CSS) ColorCyan() *CSS {
	return c.setColor(FgCyan)
}

// 设置前景色为白色
//
// Set foreground colour to whilte
func (c *CSS) ColorWhite() *CSS {
	return c.setColor(FgWhite)
}

// 设置前景色为灰色
//
// Set foreground colour to cray
func (c *CSS) ColorCray() *CSS {
	return c.setColor(FgGray)
}

// 设置前景色为亮红色
//
// Set foreground colour to bright red
func (c *CSS) ColorBrightRed() *CSS {
	return c.setColor(FgBrightRed)
}

// 设置前景色为亮绿色
//
// Set foreground colour to bright green
func (c *CSS) ColorBrightGreen() *CSS {
	return c.setColor(FgBrightGreen)
}

// 设置前景色为亮黄色
//
// Set foreground colour to bright yellow
func (c *CSS) ColorBrightYellow() *CSS {
	return c.setColor(FgBrightYellow)
}

// 设置前景色为亮蓝色
//
// Set foreground colour to bright blue
func (c *CSS) ColorBrightBule() *CSS {
	return c.setColor(FgBrightBlue)
}

// 设置前景色为亮品红色
//
// Set foreground colour to bright magenta
func (c *CSS) ColorBrightMagenta() *CSS {
	return c.setColor(FgBrightMagenta)
}

// 设置前景为亮青色
//
// Set foreground colour to bright cyan
func (c *CSS) ColorBrightCyan() *CSS {
	return c.setColor(FgBrightCyan)
}

// 设置前景色为亮白色
//
// Set foreground colour to bright whilte
func (c *CSS) ColorBrightWhite() *CSS {
	return c.setColor(FgBrightWhite)
}
