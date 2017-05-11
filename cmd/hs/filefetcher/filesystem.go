package filefetcher

import (
	"io"
	"io/ioutil"
	"os"
)

type filesystem interface {
	Open(name string) (File, error)
	ReadFile(filename string) ([]byte, error)
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

type NativeFS struct{}

func (NativeFS) Open(name string) (File, error) {
	return os.Open(name)
}

func (NativeFS) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
