package filefetcher

import (
	"net/http"
	"time"
)

type webClient interface {
	Get(url string) (resp *http.Response, err error)
}

type nativeWebClient struct {
	http.Client
}

func NewWebClient() webClient {
	var timeout = time.Duration(DEFAULT_TIMEOUT_IN_SECS * time.Second)
	return &nativeWebClient{
		Client: http.Client{Timeout: timeout},
	}
}
