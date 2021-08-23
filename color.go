package colorStyle

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
	FgDarkBlue
	FgPurple
	FgLightBlue
	FgWhite
)

func (c *CSS) ColorBlack() *CSS {
	return c.setColor(FgBlack)
}

func (c *CSS) ColorRed() *CSS {
	return c.setColor(FgRed)
}

func (c *CSS) ColorGreen() *CSS {
	return c.setColor(FgGreen)
}

func (c *CSS) ColorYellow() *CSS {
	return c.setColor(FgYellow)
}

func (c *CSS) ColorDarkBlue() *CSS {
	return c.setColor(FgDarkBlue)
}

func (c *CSS) ColorPurple() *CSS {
	return c.setColor(FgDarkBlue)
}

func (c *CSS) ColorLightBlue() *CSS {
	return c.setColor(FgLightBlue)
}

func (c *CSS) ColorWhite() *CSS {
	return c.setColor(FgWhite)
}
