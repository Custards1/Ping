package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WebGlob struct {
	Name string
	Data *[]byte
}

func NewWebGlob(filename string) (*WebGlob, error) {
	body, err := ioutil.ReadFile(filename)
	return &WebGlob{Name: filename, Data: &body}, err
}
func error_page(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "404: ERR", r.URL.Path[1:])
}

func write_glob(w http.ResponseWriter, r *http.Request, glob *WebGlob) {
	fmt.Fprint(w, string(*glob.Data))
}

func main() {
	glob, err := NewWebGlob("glob.html")
	a, e := NewWebGlob("game.js")
	if err != nil {
		panic(err)
	}
	if e != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write_glob(w, r, glob)
	})
	http.HandleFunc("/source", func(w http.ResponseWriter, r *http.Request) {
		write_glob(w, r, a)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
