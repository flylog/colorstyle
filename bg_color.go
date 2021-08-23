package colorStyle

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
	BgPurple
	BgLightBlue
	BgWhite
)

func (c *CSS) BgBlack() *CSS {
	return c.setBgColor(BgBlack)
}

func (c *CSS) BgRed() *CSS {
	return c.setBgColor(BgRed)
}

func (c *CSS) BgGreen() *CSS {
	return c.setBgColor(BgGreen)
}

func (c *CSS) BgYellow() *CSS {
	return c.setBgColor(BgYellow)
}

func (c *CSS) BgBlue() *CSS {
	return c.setBgColor(BgBlue)
}

func (c *CSS) BgPurple() *CSS {
	return c.setBgColor(BgPurple)
}

func (c *CSS) BgLightBlue() *CSS {
	return c.setBgColor(BgLightBlue)
}
