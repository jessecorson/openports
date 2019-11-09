package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func openports(p []string) {

	a := makePortList(p)

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

func makePortList(rSlice []string) []int {

	var pRange []int
	// Checking if port is an int
	if len(rSlice) < 2 {
		r := rSlice[0]
		p, err := strconv.Atoi(r)
		if err != nil {
			// Checking if range
			rSplit := strings.Split(r, "-")
			if !(len(rSplit) > 2) {
				for _, a := range rSplit {
					if i, err := strconv.Atoi(a); err == nil {
						pRange = append(pRange, i)
					}
				}
				f := pRange[0]
				l := pRange[1]
				pRange := makeRange(f, l)
				return pRange
			}
		}
		pRange = []int{p}
		return pRange
	}
	for _, a := range rSlice {
		if i, err := strconv.Atoi(a); err == nil {
			pRange = append(pRange, i)
		}
	}
	return pRange
}
