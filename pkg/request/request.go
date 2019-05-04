package request

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"time"
)

type Req struct {
	Url     string
	Method  string
	Timeout time.Duration
	Retry   int
}

type ReqApi interface {
	Request() (*http.Response, error)
}

func NewRequest() *Req {
	return &Req{}
}

func (r *Req) Request() (*http.Response, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Timeout: r.Timeout, Transport: tr}

	req, err := http.NewRequest(r.Method, r.Url, nil)

	if err != nil {
		log.Printf("http request url error:(%v)", err)
		return nil, errors.New("http request error")
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("http response url error:(%v)", err)
		return nil, errors.New("http response error")
	}

	if resp.StatusCode == 200 {
		return resp, nil
	}

	return nil, errors.New(string(resp.StatusCode))
}