package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gophercises/urlshort"
	"gopkg.in/yaml.v2"
)

func main() {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	yamlData := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	jsonData := `
[
	{
		"path": "/json",
		"url": "https://www.json.org/"
	},
	{
		"path": "/jsonlint",
		"url": "https://jsonlint.com/"
	}
]
`

	defaultHandler := buildDefaultHandler()
	mapHandler := urlshort.MapHandler(pathsToUrls, defaultHandler)
	yamlHandler, err := urlshort.DataHandler([]byte(yamlData), yaml.Unmarshal, mapHandler)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.DataHandler([]byte(jsonData), json.Unmarshal, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe("localhost:8080", jsonHandler)
}

func buildDefaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
}
