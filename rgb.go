package ansi

import "fmt"

// RGB represents a rgb color value. Each channel has to be an integer between 0 and 255.
type RGB [3]uint8

// Foreground returns an ANSI escape sequence representing a foreground of the rgb value.
func (rgb RGB) Foreground() ESC {
	return esc("38;2;" + rgb.String())
}

// FG is a convenience function for Foreground.
func (rgb RGB) FG() ESC {
	return rgb.Foreground()
}

// Background returns an ANSI escape sequence representing a background of the rgb value.
func (rgb RGB) Background() ESC {
	return esc("48;2;" + rgb.String())
}

// BG is a convenience function for Background.
func (rgb RGB) BG() ESC {
	return rgb.Background()
}

// String returns a semi-colon-separated list of the red, green, and blue channel.
func (rgb RGB) String() string {
	return fmt.Sprintf("%d;%d;%d", rgb[0], rgb[1], rgb[2])
}
