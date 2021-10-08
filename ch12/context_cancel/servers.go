package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)

func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		rw.Write([]byte("Slow response"))
	}))
	return s
}

func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			rw.Write([]byte("error"))
			return
		}
		rw.Write([]byte("ok"))
	}))
	return s
}