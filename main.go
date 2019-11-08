package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	helpUrl = "https://golang.org/doc/articles/wiki/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; path {
	case "":
		fmt.Fprintf(w, "Hello world!")
	case "help":
		fmt.Fprintf(w, helpUrl)
	default:
		fmt.Fprintf(w, "There is nothing at %s", r.URL.Path[1:])
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {

	a := makeRange(0, 255)
	http.HandleFunc("/", handler)
	fmt.Println("here")
	for _, p := range a {
		// fmt.Println(p)
		ps := strconv.Itoa(p)
		go http.ListenAndServe(":"+ps, nil)
		fmt.Println("listening on " + ps)
	}

	// Keep process running
	for {
	}
}
