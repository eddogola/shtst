package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"

	"github.com/eddogola/shtst"
)

func main() {
	mux := defaultMux()

	// cmdline options
	jsonFile := flag.String("json", "urls.json", "path to json file mapping paths to urls")
	yamlFile := flag.String("yaml", "urls.yml", "path to yaml file mapping paths to urls")
	flag.Parse()

	// map handler
	mapToUrls := map[string]string{
		"/mit": "https://web.mit.edu",
		"/golang": "https://golang.org",
		"/goodreads": "https://goodreads.com",
	}
	mapHandler := shtst.MapHandler(mapToUrls, mux)	


	// json handler
	jsonData, err := ioutil.ReadFile(*jsonFile)
	check(err)
	jsonHandler, err := shtst.JSONHandler(jsonData, mapHandler)
	check(err)

	// yaml handler
	yamlData, err := ioutil.ReadFile(*yamlFile)
	check(err)
	yamlHandler, err := shtst.YAMLHandler(yamlData, jsonHandler)
	check(err)

	fmt.Println("Listening from localhost port 5000")
	http.ListenAndServe(":5000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Love galore")
	})
	return mux	 
}

func check(e error) {
	if e != nil {
		fmt.Println("encountered an unexpected error", e)
	}
}