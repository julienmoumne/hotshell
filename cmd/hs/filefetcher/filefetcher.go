package filefetcher

import (
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"time"
)

const DEFAULT_TIMEOUT_IN_SECS time.Duration = 5

type Filefetcher struct {
	DefaultHSFilename string
	Fs                filesystem
	WebClient         webClient
	path              string
	url               *url.URL
	pathToLoad        string
	fileContent       []byte
	fileMode          os.FileMode
}

func (f *Filefetcher) FetchFile(path string) ([]byte, string, error) {
	if path == "" {
		return nil, "", errors.New("invalid path : empty string")
	}

	f.path = path
	if err := f.parsePath(); err != nil {
		return nil, "", err
	}

	if f.isLocal() {
		if err := f.loadLocalFile(); err != nil {
			return nil, "", err
		}
	} else {
		if err := f.fetchRemoteFile(); err != nil {
			return nil, "", err
		}
	}

	return f.fileContent, f.pathToLoad, nil
}

func (f *Filefetcher) parsePath() error {

	url, err := url.Parse(f.path)
	if err != nil {
		return err
	}

	f.url = url

	return nil
}

func (f *Filefetcher) isLocal() bool {
	return f.url.Scheme == "" || f.url.Scheme == "file"
}

func (f *Filefetcher) loadLocalFile() error {
	if err := f.loadFileMode(); err != nil {
		return err
	}

	f.pathToLoad = f.url.Path
	if f.fileMode.IsDir() {
		f.pathToLoad += "/" + f.DefaultHSFilename
	}

	if err := f.loadLocalFileContent(); err != nil {
		return err
	}

	return nil
}

func (f *Filefetcher) loadLocalFileContent() error {
	buf, err := f.Fs.ReadFile(f.pathToLoad)
	f.fileContent = buf
	return err
}

func (f *Filefetcher) loadFileMode() error {

	file, err := f.Fs.Open(f.url.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	f.fileMode = fileInfo.Mode()

	return nil
}

func (f *Filefetcher) fetchRemoteFile() error {
	f.pathToLoad = f.path

	response, err := f.WebClient.Get(f.pathToLoad)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	f.fileContent = contents

	return nil
}
