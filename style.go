package colorStyle

import "strconv"

type Style int

func (s Style) ptrString() *string {
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

func (c *CSS) StyleDefault() *CSS {
	return c.setStyle(Default)
}

func (c *CSS) StyleBold() *CSS {
	return c.setStyle(Bold)
}

func (c *CSS) StyleGrey() *CSS {
	return c.setStyle(Grey)
}

func (c *CSS) StyleItalic() *CSS {
	return c.setStyle(Italic)
}

func (c *CSS) StyleUnderline() *CSS {
	return c.setStyle(Underline)
}

func (c *CSS) StyleReverse() *CSS {
	return c.setStyle(Reverse)
}

func (c *CSS) StyleStrikethrough() *CSS {
	return c.setStyle(Strikethrough)
}
