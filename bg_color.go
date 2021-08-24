package colorstyle

import "strconv"

type BgColor int

func (b BgColor) ptrString() *string {
	str := strconv.Itoa(int(b))
	return &str
}

const (
	BgBlack BgColor = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgGray BgColor = iota + 92
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWhite
)

// 设置背景颜色为黑色
//
// Set the background colour to black
func (c *CSS) BgBlack() *CSS {
	return c.setBgColor(BgBlack)
}

// 设置背景颜色为红色
//
// Set the background colour to red
func (c *CSS) BgRed() *CSS {
	return c.setBgColor(BgRed)
}

// 设置背景颜色为绿色
//
// Set the background colour to green
func (c *CSS) BgGreen() *CSS {
	return c.setBgColor(BgGreen)
}

// 设置背景颜色为黄色
//
// Set the background colour to yellow
func (c *CSS) BgYellow() *CSS {
	return c.setBgColor(BgYellow)
}

// 设置背景颜色为 蓝色
//
// Set the background colour to blue
func (c *CSS) BgBlue() *CSS {
	return c.setBgColor(BgBlue)
}

// 设置背景颜色为品红
//
// Set the background colour to magenta
func (c *CSS) BgMagenta() *CSS {
	return c.setBgColor(BgMagenta)
}

// 设置背景颜色为青色
//
// Set the background colour to cyan
func (c *CSS) BgCyan() *CSS {
	return c.setBgColor(BgCyan)
}

// 设置背景颜色为白色
//
// Set the background colour to white
func (c *CSS) BgWhite() *CSS {
	return c.setBgColor(BgWhite)
}

// 设置背景颜色为灰色
//
// Set the background colour to gray
func (c *CSS) BgGray() *CSS {
	return c.setBgColor(BgGray)
}

// 设置背景颜色为亮红色
//
// Set the background colour to bright red
func (c *CSS) BgBrightRed() *CSS {
	return c.setBgColor(BgBrightRed)
}

// 设置背景颜色为亮绿色
//
// Set the background colour to bright green
func (c *CSS) BgBrightGreen() *CSS {
	return c.setBgColor(BgBrightGreen)
}

// 设置背景颜色为亮黄色
//
// Set the background colour to bright yellow
func (c *CSS) BgBrightYellow() *CSS {
	return c.setBgColor(BgBrightYellow)
}

// 设置背景颜色为亮蓝色
//
// Set the background colour to bright blue
func (c *CSS) BgBrightBlue() *CSS {
	return c.setBgColor(BgBrightBlue)
}

// 设置背景颜色为亮品红色
//
// Set the background colour to  bright magenta
func (c *CSS) BgBrightMagenta() *CSS {
	return c.setBgColor(BgBrightMagenta)
}

// 设置背景颜色为亮青色
//
// Set the background colour to bright cyan
func (c *CSS) BgBrightCyan() *CSS {
	return c.setBgColor(BgBrightCyan)

}

// 设置背景颜色为亮白色
//
// Set the background colour to bright white
func (c *CSS) BgBrightWhite() *CSS {
	return c.setBgColor(BgBrightWhite)
}
