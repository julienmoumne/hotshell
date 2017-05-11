package filefetcher_test

import (
	"errors"
	"github.com/julienmoumne/hotshell/cmd/hs/filefetcher"
	"os"
	"time"
)

type fileSystemStub struct {
	registeredFiles []fileStub
}

func (f fileSystemStub) Open(filename string) (filefetcher.File, error) {
	for _, testFile := range f.registeredFiles {
		if filename == testFile.path {
			return testFile, nil
		}
	}

	return nil, errors.New("")
}

func (f fileSystemStub) ReadFile(filename string) ([]byte, error) {
	if file, err := f.Open(filename); err == nil {
		tFile := file.(fileStub)
		if tFile.fileInfo.fileMode.IsRegular() {
			return file.(fileStub).content, nil
		}
	}

	return nil, errors.New("")
}

type fileStub struct {
	path     string
	content  []byte
	fileInfo fileInfoStub
}

func (f fileStub) Close() error                                  { return nil }
func (f fileStub) Read(p []byte) (n int, err error)              { panic("not implemented") }
func (f fileStub) ReadAt(p []byte, off int64) (n int, err error) { panic("not implemented") }
func (f fileStub) Seek(offset int64, whence int) (int64, error)  { panic("not implemented") }
func (f fileStub) Stat() (os.FileInfo, error)                    { return f.fileInfo, nil }

type fileInfoStub struct {
	fileMode os.FileMode
}

func (f fileInfoStub) Name() string       { panic("not implemented") }
func (f fileInfoStub) Size() int64        { panic("not implemented") }
func (f fileInfoStub) Mode() os.FileMode  { return f.fileMode }
func (f fileInfoStub) ModTime() time.Time { panic("not implemented") }
func (f fileInfoStub) IsDir() bool        { panic("not implemented") }
func (f fileInfoStub) Sys() interface{}   { panic("not implemented") }
