package main

import (
	"net/http"
	"time"
)

type HelloHander struct{}

func (hh HelloHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

func main() {
	s := http.Server{
		Addr: ":8080",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: HelloHander{},
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	})

	person := http.NewServeMux()
		person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})
	dog := http.NewServeMux()
		dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good puppy!\n"))
	})
	mux.Handle("/person/", http.StripPrefix("/person", person)) // to remove the part of the path that’s already been processed by mux
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))
}