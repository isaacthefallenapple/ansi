package ansi

import "fmt"

type RGB [3]uint8

func (rgb RGB) Foreground() ESC {
	return _esc("38;2;" + rgb.String())
}
func (rgb RGB) FG() ESC {
	return rgb.Foreground()
}
func (rgb RGB) Background() ESC {
	return _esc("48;2;" + rgb.String())
}
func (rgb RGB) BG() ESC {
	return rgb.Background()
}
func (rgb RGB) String() string {
	return fmt.Sprintf("%d;%d;%d", rgb[0], rgb[1], rgb[2])
}
