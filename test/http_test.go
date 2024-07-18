package test

import (
	"io"
	"net/http"
	"testing"
)

const (
	URL = "https://jsonplaceholder.typicode.com"
)

func TestHTTPRequest(t *testing.T) {
	res, err := http.Get(URL + "/todos/1")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
