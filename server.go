package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", path.Base(r.URL.EscapedPath()))
	var html, err = loadFile(path.Base(r.URL.EscapedPath()))
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404, html file doesn't exist!")
	}
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
