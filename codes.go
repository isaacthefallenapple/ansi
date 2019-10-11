// Package ansi defines a simple ANSI api to format console in- and output.
package ansi

import (
	"fmt"
	"strings"
)

// Code represents an ANSI code
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

// A Painter can "paint" a string.
type Painter interface {
	// Paint takes a string and returns it wrapped in the caller's ANSI code
	// and terminated with TERM.
	Paint(string) string
	// Start colors the console in the Painter's color until Stop is called.
	Start()
	// Stop resets the console color scheme.
	Stop()
}

// Paint will wrap s in the ANSI escape corresponding to c.
func (c Code) Paint(s string) string {
	return Paint(c.ESC(), s)
}

// String returns the code as a string.
func (c Code) String() string {
	return fmt.Sprintf("%d", c)
}

// Compose returns an escape code combining given codes.
func Compose(effect, fg, bg Code) ESC {
	return esc(fmt.Sprintf("%d;%d;%d", effect, fg, bg))
}

// Chain returns an escape code chaining together the given codes left to right.
// The right-most code takes precedence, should two codes be incompatible.
func Chain(codes ...Code) ESC {
	escapes := make([]string, len(codes))
	for i, code := range codes {
		escapes[i] = code.ESC().String()
	}
	return ESC(strings.Join(escapes, ""))
}

// ESC returns the escape sequence corresponding to c.
func (c Code) ESC() ESC {
	return esc(c.String())
}

// Start will apply the format corresponding to c to the console until Stop is called.
func (c Code) Start() {
	fmt.Print(c.ESC())
}

// Stop trivially calls Term() to reset the console color scheme.
func (c Code) Stop() {
	Term()
}
