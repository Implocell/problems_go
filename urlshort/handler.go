package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathstoUrls := buildMap(pathUrls)

	return MapHandler(pathstoUrls, fallback), nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJson(jsonBytes)
	if err != nil {
		return nil, err
	}
	pathstoUrls := buildMap(pathUrls)

	return MapHandler(pathstoUrls, fallback), nil
}

func buildMap(pathUrls []pathtUrl) map[string]string {
	pathstoUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathstoUrls[pu.Path] = pu.URL
	}
	return pathstoUrls
}

func parseYaml(data []byte) ([]pathtUrl, error) {
	var pathUrls []pathtUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func parseJson(data []byte) ([]pathtUrl, error) {
	var pathUrls []pathtUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathtUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type pathUrlJson struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
