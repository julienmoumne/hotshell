package fileloader

import (
	"github.com/blang/vfs"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const timeoutInSeconds = 5

// todo because this file is short and not reused it could be merged into definitionloader
type Loader struct {
	fs          vfs.Filesystem
	path        string
	fileContent []byte
}

func (f *Loader) Load(fs vfs.Filesystem, path string) ([]byte, error) {
	f.fs = fs
	f.path = path
	var isLocal, err = f.isLocal()
	if err != nil {
		return nil, err
	}
	if isLocal {
		err = f.loadLocalFile()
	} else {
		err = f.fetchRemoteFile()
	}
	return f.fileContent, err
}

func (f *Loader) isLocal() (bool, error) {
	u, err := url.Parse(f.path)
	if err != nil {
		return true, err
	}
	return u.Scheme == "" || u.Scheme == "file", nil
}

func (f *Loader) loadLocalFile() error {
	var err error
	f.fileContent, err = vfs.ReadFile(f.fs, f.path)
	return err
}

func (f *Loader) fetchRemoteFile() error {
	c := http.Client{
		Timeout: time.Duration(timeoutInSeconds * time.Second),
	}
	response, err := c.Get(f.path)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	f.fileContent, err = ioutil.ReadAll(response.Body)
	return err
}