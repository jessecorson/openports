package cmd

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	helpURL = "https://golang.org/doc/articles/wiki/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; path {
	case "":
		fmt.Fprintf(w, "Hello world!")
	case "help":
		fmt.Fprintf(w, helpURL)
	default:
		fmt.Fprintf(w, "There is nothing at %s", r.URL.Path[1:])
	}
}

func makeRange(min, max int) []int {

	if min > max {
		fmt.Printf("High port %v was less than low port %v, only using low port %v\n", max, min, min)
		return []int{min}
	}
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func openports(r []int) {

	f := r[0]
	l := r[1]

	a := makeRange(f, l)
	http.HandleFunc("/", handler)
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
