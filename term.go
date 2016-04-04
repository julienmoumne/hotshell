package hotshell

import (
	"fmt"
	"github.com/julienmoumne/hotshell/item"
	pkgterm "github.com/pkg/term"
)

const DEFAULT_TTY = "/dev/tty"

var Tty = DEFAULT_TTY

type term struct {
	term *pkgterm.Term
}

func NewTerm() (*term, error) {

	t, err := pkgterm.Open(Tty)

	if err != nil {
		return nil, err
	}

	return &term{t}, err
}

func (t *term) close() error {

	return t.term.Close()
}

func (t *term) restore() {

	// term.Restore() blocks when running tests with pseudo term (termios.Pty())
	if Tty != DEFAULT_TTY {
		return
	}

	err := t.term.Restore()
	if err == nil {
		return
	}

	fmt.Printf("An error occurred while restoring your terminal default values : %s\n", err)
	fmt.Println("Your terminal may behave differently than usual.")
	fmt.Println("If it is the case, you can close and start it again.")
	fmt.Println("Please file a bug report at https://github.com/julienmoumne/hotshell/issues")
}

func (t *term) readUserChoice() (item.Key, error) {

	err := pkgterm.CBreakMode(t.term)
	defer t.restore()

	if err != nil {
		return item.NUL_KEY, err
	}

	bytes := make([]byte, 1)
	_, err = t.term.Read(bytes)

	return item.MakeKey(string(bytes)), err
}
