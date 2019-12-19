package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 8080, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	fmt.Println(*filename)
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	handler := cyoa.NewHander(story, getPathFromRequest)
	fmt.Printf("Starting tht server on :%v\n", *port)
	http.ListenAndServe(fmt.Sprintf("localhost:%v", *port), handler)
}

func getPathFromRequest(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/" {
		path = "/intro"
	}
	return path[len("/"):]
}
