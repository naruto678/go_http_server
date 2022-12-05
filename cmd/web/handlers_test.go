package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Got error while creating the goddamn request")
	}
	app := &application{}
	app.PingHandler(recorder, req)

	res := recorder.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("want [%d] Got [%d]", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "OK" {
		t.Errorf("want body to be equal to %s", "OK")
	}

}
