package colorStyle

import "strconv"

type BgColor int

func (b BgColor) String() *string {
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

func (c *colorStyle) BgBlack() *colorStyle {
	return c.SetBgColor(BgBlack)
}

func (c *colorStyle) BgRed() *colorStyle {
	return c.SetBgColor(BgRed)
}

func (c *colorStyle) BgGreen() *colorStyle {
	return c.SetBgColor(BgGreen)
}

func (c *colorStyle) BgYellow() *colorStyle {
	return c.SetBgColor(BgYellow)
}

func (c *colorStyle) BgBlue() *colorStyle {
	return c.SetBgColor(BgBlue)
}

func (c *colorStyle) BgPurple() *colorStyle {
	return c.SetBgColor(BgPurple)
}

func (c *colorStyle) BgLightBlue() *colorStyle {
	return c.SetBgColor(BgLightBlue)
}
