package shtst

import (
	"net/http"
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(urlsMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := urlsMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func JSONHandler(JsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse json
	mappings, err := parseJson(JsonBytes)
	if err != nil {
		return nil, err
	}
	// Convert to map
	pathsToUrls := toMap(mappings)
	//return from MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse yaml
	mappings, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	// Convert to map
	pathsToUrls := toMap(mappings)
	// return from MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

func parseJson(data []byte) ([]PathUrl, error) {
	var mappings []PathUrl
	err := json.Unmarshal(data, &mappings)
	if err != nil {
		return nil, err
	}
	return mappings, err
}

func parseYaml(data []byte) ([]PathUrl, error) {
	var mappings []PathUrl
	err := yaml.Unmarshal(data, &mappings)
	if err != nil {
		return nil, err
	}
	return mappings, err
}

func toMap(mappings []PathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, mapping := range mappings {
		pathsToUrls[mapping.Path] = mapping.Url
	}
	return pathsToUrls
}

type PathUrl struct {
	Path string `yaml:"path" json:"path"`
	Url string `yaml:"url" json:"url"`
}