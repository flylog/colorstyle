package colorStyle

import "strconv"

type Color int

func (c Color) String() *string {
	str := strconv.Itoa(int(c))
	return &str
}

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	DarkBlue
	Purple
	LightBlue
	White
)

func (c *colorStyle) ColorBlack() *colorStyle {
	return c.SetColor(Black)
}

func (c *colorStyle) ColorRed() *colorStyle {
	return c.SetColor(Red)
}

func (c *colorStyle) ColorGreen() *colorStyle {
	return c.SetColor(Green)
}

func (c *colorStyle) ColorYellow() *colorStyle {
	return c.SetColor(Yellow)
}

func (c *colorStyle) ColorDarkBlue() *colorStyle {
	return c.SetColor(DarkBlue)
}

func (c *colorStyle) ColorPurple() *colorStyle {
	return c.SetColor(DarkBlue)
}

func (c *colorStyle) ColorLightBlue() *colorStyle {
	return c.SetColor(LightBlue)
}

func (c *colorStyle) ColorWhite() *colorStyle {
	return c.SetColor(White)
}
