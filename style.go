package colorstyle

import "strconv"

type Style int

func (s Style) ptrString() *string {
	str := strconv.Itoa(int(s))
	return &str
}

const (
	Normal        Style = iota // 默认值
	Bold                       // 加粗
	Grey                       // 灰显
	Italic                     // 斜体
	Underline                  // 下划线
	SlowBlink                  // 缓慢闪烁
	RapidBlink                 // 快速闪烁
	Reverse                    // 反显
	Hide                       // 隐藏
	Strikethrough              // 删除线
)

// 设置字体样式为常规
//
// Set font style to default
func (c *CSS) StyleDefault() *CSS {
	return c.setStyle(Normal)
}

// 设置字体样式为粗体
//
// Set font style to bold
func (c *CSS) StyleBold() *CSS {
	return c.setStyle(Bold)
}

// 设置字体样式为灰显
//
// Set font style to grey
func (c *CSS) StyleGrey() *CSS {
	return c.setStyle(Grey)
}

// 设置字体样式为斜体
//
// Set font style to italic

func (c *CSS) StyleItalic() *CSS {
	return c.setStyle(Italic)
}

// 设置字体样式为下划线
//
// Set font style to underline
func (c *CSS) StyleUnderline() *CSS {
	return c.setStyle(Underline)
}

// 设置字体样式为缓慢闪烁
//
// Set font style t0 rapid blink
func (c *CSS) StyleRapidBlink() *CSS {
	return c.setStyle(RapidBlink)
}

// 设置字体样式为快速闪烁
//
// Set font style t0 slow blink
func (c *CSS) StyleSlowBlink() *CSS {
	return c.setStyle(SlowBlink)
}

// 设置字体样式为反显
//
// Set font style to reverse
func (c *CSS) StyleReverse() *CSS {
	return c.setStyle(Reverse)
}

// 设置字体样式为隐藏
//
// Set font style to hide
func (c *CSS) StyleHide() *CSS {
	return c.setStyle(Hide)
}

// 设置字体样式为删除线,可能不支持
//
// Set font style to strikethrough，may not be supported
func (c *CSS) StyleStrikethrough() *CSS {
	return c.setStyle(Strikethrough)
}
