package ansi

import (
	"fmt"
	"strings"
)

type Code uint8

const (
	Default     Code = 0
	Bold        Code = 1
	Underline   Code = 4
	Negative    Code = 7
	NoUnderline Code = 24
)

const (
	FG_Black Code = iota + 30
	FG_Red
	FG_Green
	FG_Yellow
	FG_Blue
	FG_Magenta
	FG_Cyan
	FG_White
	FG_Extended
	FG_Default
)

const (
	BG_Black Code = iota + 40
	BG_Red
	BG_Green
	BG_Yellow
	BG_Blue
	BG_Magenta
	BG_Cyan
	BG_White
	BG_Extended
	BG_Default
)

func (c Code) Paint(s string) string {
	return Paint(esc(c), s)
}
func (c Code) String() string {
	return fmt.Sprintf("%d", c)
}
func esc(c Code) ESC {
	return _esc(c.String())
}
func Compose(effect, fg, bg Code) ESC {
	return _esc(fmt.Sprintf("%d;%d;%d", effect, fg, bg))
}
func Chain(codes ...Code) ESC {
	escapes := make([]string, len(codes))
	for i, code := range codes {
		escapes[i] = esc(code).String()
	}
	return ESC(strings.Join(escapes, ""))
}
