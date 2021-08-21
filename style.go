package colorStyle

import "strconv"

type Style int

func (s Style) String() *string {
	str := strconv.Itoa(int(s))
	return &str
}

const (
	Default       Style = iota // 默认值
	Bold                       // 加粗
	Grey                       // 灰显
	Italic                     // 斜体
	Underline                  // 下划线
	Reverse                    // 反显
	Strikethrough              // 删除线
)

func (c *colorStyle) StyleDefault() *colorStyle {
	return c.SetStyle(Default)
}

func (c *colorStyle) StyleBold() *colorStyle {
	return c.SetStyle(Bold)
}

func (c *colorStyle) StyleGrey() *colorStyle {
	return c.SetStyle(Grey)
}

func (c *colorStyle) StyleItalic() *colorStyle {
	return c.SetStyle(Italic)
}

func (c *colorStyle) StyleUnderline() *colorStyle {
	return c.SetStyle(Underline)
}

func (c *colorStyle) StyleReverse() *colorStyle {
	return c.SetStyle(Reverse)
}

func (c *colorStyle) StyleStrikethrough() *colorStyle {
	return c.SetStyle(Strikethrough)
}
