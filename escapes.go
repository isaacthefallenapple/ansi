package ansi

import "fmt"

const TERM ESC = "\x1b[m"

type ESC string
func (e ESC) Paint(s string) string {
	return Paint(e, s)
}
func (e ESC) String() string {
	return string(e)
}

func Term() {
	fmt.Print(TERM)
}

func Paint(escape ESC, s string) string {
	return fmt.Sprintf("%s%s%s", escape, s, TERM)
}

func _esc(s string) ESC {
	return ESC(fmt.Sprintf("\x1b[%sm", s))
}
