package filefetcher_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type webClientStub struct {
	registeredFile    string
	registeredContent []byte
}

func (w webClientStub) Get(url string) (resp *http.Response, err error) {

	if url != w.registeredFile {
		return nil, errors.New("")
	}

	return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(w.registeredContent))}, nil
}
