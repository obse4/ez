package net

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, contentType, data string) (body []byte, code int, err error) {
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	code = resp.StatusCode
	if err != nil {
		return nil, code, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, code, err
	}
	return body, code, err
}

func HttpGet(url string) (body []byte, code int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, err
}
