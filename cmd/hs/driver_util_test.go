package main

import (
	"bytes"
	"github.com/julienmoumne/hotshell"
	"github.com/pkg/term/termios"
	"io"
	"os"
)

type driver struct {
	menuDefinition string
	input          []byte
	osCwd          string
	osStdin        *os.File
	osStdout       *os.File
	osStderr       *os.File
	readOut        *os.File
	writeOut       *os.File
	readErr        *os.File
	writeErr       *os.File
	ptm            *os.File
	pts            *os.File
}

func (d *driver) run() (string, string, error) {
	d.backupStds()
	defer d.restoreStds()

	if err := d.setupPty(); err != nil {
		return "", "", err
	}
	defer d.closePty()

	if err := d.setupPipes(); err != nil {
		return "", "", err
	}
	defer d.closePipes()

	if err := d.backupCwd(); err != nil {
		return "", "", err
	}
	defer d.restoreCwd()

	d.execMain()

	return d.getStds()
}

func (d *driver) restoreCwd() {
	os.Chdir(d.osCwd) // ignored returned Error
}

func (d *driver) backupCwd() error {
	var err error
	d.osCwd, err = os.Getwd()
	return err
}

func (d *driver) execMain() {
	os.Args = []string{"", "--chdir", "-f", d.menuDefinition}

	// todo test exitCode
	var exitCode int
	exit = func(code int) {
		exitCode = code
	}

	main()
}

func (d *driver) setupPty() error {
	var err error
	d.ptm, d.pts, err = termios.Pty()
	if err != nil {
		return err
	}

	if _, err = d.ptm.Write(d.input); err != nil {
		return err
	}

	hotshell.Tty = d.pts.Name()
	os.Stdin = d.pts

	return err
}

func (d *driver) closePty() {
	d.ptm.Close() // ignored returned Error
	d.pts.Close() // ignored returned Error
}

func (d *driver) setupPipes() error {
	if err := d.setupOutPipe(); err != nil {
		return err
	}
	if err := d.setupErrPipe(); err != nil {
		return err
	}
	return nil
}

func (d *driver) closePipes() {
	d.readOut.Close()  // ignored returned Error
	d.writeOut.Close() // ignored returned Error
	d.readErr.Close()  // ignored returned Error
	d.writeErr.Close() // ignored returned Error
}

func (d *driver) setupOutPipe() error {
	var err error
	d.readOut, d.writeOut, err = os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = d.writeOut
	return nil
}

func (d *driver) setupErrPipe() error {
	var err error
	d.readErr, d.writeErr, err = os.Pipe()
	if err != nil {
		return err
	}
	os.Stderr = d.writeErr
	return nil
}

func (d *driver) backupStds() {
	d.osStdin = os.Stdin
	d.osStdout = os.Stdout
	d.osStderr = os.Stderr
}

func (d *driver) restoreStds() {
	os.Stdin = d.osStdin
	os.Stdout = d.osStdout
	os.Stderr = d.osStderr
}

func (d *driver) getStds() (string, string, error) {
	stdout, err := fileToBuf(d.readOut, d.writeOut)
	if err != nil {
		return "", "", err
	}
	stderr, err := fileToBuf(d.readErr, d.writeErr)
	if err != nil {
		return "", "", err
	}
	return stdout, stderr, nil
}

func fileToBuf(read *os.File, write *os.File) (string, error) {
	write.Close() // failing to close makes io.Copy hangs
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, read); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}
