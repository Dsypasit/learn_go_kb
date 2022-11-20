package http

import (
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

var get = http.Get

func MakeHTTPCall(url string) (*Response, error) {
	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}
