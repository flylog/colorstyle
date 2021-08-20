package colorStyle

import "strconv"

type Style int

func (s Style) String() string {
	return strconv.Itoa(int(s))
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

type Color int

func (c Color) String() string {
	return strconv.Itoa(int(c))
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

type BgColor int

func (b BgColor) String() string {
	return strconv.Itoa(int(b))
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
