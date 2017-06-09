package term

import (
	"bytes"
	"github.com/pkg/term/termios"
	"io"
	"os"
)

type TestDriver struct {
	Input    []byte
	Main     func()
	Cwd      string
	osCwd    string
	osStdin  *os.File
	osStdout *os.File
	osStderr *os.File
	readOut  *os.File
	writeOut *os.File
	readErr  *os.File
	writeErr *os.File
	ptm      *os.File
	pts      *os.File
}

func (d *TestDriver) Run() (string, string) {
	d.backupStds()
	defer d.restoreStds()

	d.setupPty()
	defer d.closePty()

	d.setupPipes()
	defer d.closePipes()

	d.backupCwd()

	if d.Cwd != "" {
		if err := os.Chdir(d.Cwd); err != nil {
			panic(err)
		}
	}
	defer d.restoreCwd()

	d.Main()

	return fileToBuf(d.readOut, d.writeOut), fileToBuf(d.readErr, d.writeErr)
}

func (d *TestDriver) restoreCwd() {
	if err := os.Chdir(d.osCwd); err != nil {
		panic(err)
	}
}

func (d *TestDriver) backupCwd() {
	var err error
	d.osCwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func (d *TestDriver) setupPty() {
	var err error
	d.ptm, d.pts, err = termios.Pty()
	if err != nil {
		panic(err)
	}

	if _, err = d.ptm.Write(d.Input); err != nil {
		panic(err)
	}

	tty = d.pts.Name()
	os.Stdin = d.pts
}

func (d *TestDriver) closePty() {
	if err := d.ptm.Close(); err != nil {
		panic(err)
	}
	if err := d.pts.Close(); err != nil {
		panic(err)
	}
}

func (d *TestDriver) setupPipes() {
	d.setupOutPipe()
	d.setupErrPipe()
}

func (d *TestDriver) closePipes() {
	if err := d.readOut.Close(); err != nil {
		panic(err)
	}
	// panics with "invalid argument"
	//if err := d.writeOut.Close(); err != nil {
	//	panic(err)
	//}
	if err := d.readErr.Close(); err != nil {
		panic(err)
	}
	// panics with "invalid argument"
	//if err := d.writeErr.Close(); err != nil {
	//	panic(err)
	//}
}

func (d *TestDriver) setupOutPipe() {
	var err error
	d.readOut, d.writeOut, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = d.writeOut
}

func (d *TestDriver) setupErrPipe() {
	var err error
	d.readErr, d.writeErr, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = d.writeErr
}

func (d *TestDriver) backupStds() {
	d.osStdin = os.Stdin
	d.osStdout = os.Stdout
	d.osStderr = os.Stderr
}

func (d *TestDriver) restoreStds() {
	os.Stdin = d.osStdin
	os.Stdout = d.osStdout
	os.Stderr = d.osStderr
}

func fileToBuf(read *os.File, write *os.File) string {
	write.Close() // failing to close makes io.Copy hangs
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, read); err != nil {
		panic(err)

	}
	return string(buf.Bytes())
}
