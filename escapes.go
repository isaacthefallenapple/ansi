package ansi

import "fmt"

// TERM is the ANSI code to reset the console color scheme.
const TERM ESC = "\x1b[m"

// An ESC represents an ANSI escape code.
type ESC string
// Paint wraps s in the escape code.
func (e ESC) Paint(s string) string {
	return Paint(e, s)
}
// String returns the escape code as a string.
func (e ESC) String() string {
	return string(e)
}

// Term prints TERM to stdout, resetting the console color scheme.
func Term() {
	fmt.Print(TERM)
}

// Paint wraps the string s in the escape code escape.
func Paint(escape ESC, s string) string {
	return fmt.Sprintf("%s%s%s", escape, s, TERM)
}

func _esc(s string) ESC {
	return ESC(fmt.Sprintf("\x1b[%sm", s))
}
