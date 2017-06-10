package term

import (
	"fmt"
	pkgterm "github.com/pkg/term"
)

const defaultTty = "/dev/tty"

var tty = defaultTty

type Term struct {
	term *pkgterm.Term
}

func NewTerm() (Term, error) {
	t, err := pkgterm.Open(tty)
	if err != nil {
		return Term{}, err
	}
	return Term{t}, err
}

func (t *Term) Close() error {
	return t.term.Close()
}

func (t *Term) Restore() {
	// term.Restore() blocks when running tests with pseudo term (termios.Pty())
	if tty != defaultTty {
		return
	}

	err := t.term.Restore()
	if err == nil {
		return
	}

	fmt.Printf("An error occurred while restoring your terminal default values : %s\n", err)
	fmt.Println("Your terminal may behave differently than usual.")
	fmt.Println("If it is the case, you can close and start it again.")
	fmt.Println("Please file a bug report at https://github.com/julienmoumne/hotshell/issues/new")
}

func (t *Term) ReadUserChoice() (string, error) {
	err := pkgterm.CBreakMode(t.term)
	defer t.Restore()
	if err != nil {
		return "", err
	}
	bs := make([]byte, 4)
	n, err := t.term.Read(bs)
	if err != nil {
		return "", err
	}
	// hack required for testing (test inputs need to be sent in fixed byte count, ie. 4, as they are sent immediately)
	for i, b := range bs {
		if b != 0 {
			return string(bs[i:n]), nil
		}
	}
	// during tests, sending 4 to the fake tty reads 0..
	return string(4), nil
}
