package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	data := "helloworld"
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/?data=%v", data), nil) 
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.HandlerFunc(echoHandler).ServeHTTP(rr, req)
	body, err := io.ReadAll(rr.Body)
	code := rr.Result().StatusCode
	if err != nil {
		t.Fatal(err)
	}
	if code != http.StatusOK {
		t.Fatalf("StatusCode: %v\nBody:\n%v", code, body)
	}
	if string(body) != data {
		t.Fatalf("StatusCode: %v\nBody:\n%v", code, body)
	}
}
