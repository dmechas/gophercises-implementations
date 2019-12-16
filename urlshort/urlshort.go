package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

// DataHandler will parse the provided data and then return an http.Handler
func DataHandler(data []byte, dataUnmarshal func(data []byte, v interface{}) error, fallback http.Handler) (http.Handler, error) {
	pathUrls, err := parse(data, dataUnmarshal)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	urlMap := make(map[string]string)
	for _, u := range pathUrls {
		urlMap[u.Path] = u.URL
	}

	return urlMap
}

func parse(data []byte, unmarshal func(data []byte, v interface{}) error) ([]pathURL, error) {
	var pathUrls []pathURL
	err := unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
