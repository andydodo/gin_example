package request

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Req struct {
	Url     string
	Method  string
	Timeout time.Duration
	Retry   int
	// add max worker nums
	MaxNums int
}

type ReqApi interface {
	Request() (*http.Response, error)
}

func NewRequest() *Req {
	return &Req{}
}

func (r *Req) Request() ([]byte, error) {
	tr := &http.Transport{
		MaxIdleConnsPerHost: r.MaxNums,
		MaxIdleConns:        20 * r.MaxNums,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Timeout: r.Timeout, Transport: tr}

	req, err := http.NewRequest(r.Method, r.Url, nil)

	if err != nil {
		log.Printf("http request url error:(%v)", err)
		return nil, errors.New("http request error")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("http response url error:(%v)", err)
		return nil, errors.New("http response error")
	}

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New(string(resp.StatusCode))
}
