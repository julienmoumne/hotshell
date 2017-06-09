package term

import (
	"fmt"
	"github.com/julienmoumne/hotshell/cmd/hs/item"
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

func (t *Term) ReadUserChoice() (item.Key, error) {
	err := pkgterm.CBreakMode(t.term)
	defer t.Restore()

	if err != nil {
		return item.Key{}, err
	}

	bytes := make([]byte, 1)
	_, err = t.term.Read(bytes)

	// during tests, writing 4 to the fake pty results in 0 being read here..
	if bytes[0] == 0 {
		bytes[0] = 4
	}

	return item.MakeKey(string(bytes)), err
}
