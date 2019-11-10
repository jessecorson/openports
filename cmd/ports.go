package cmd

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	helpURL  = "https://github.com/jessecorson/openports"
	helpHTML = fmt.Sprintf("<a href=%s>%s</a>", helpURL, helpURL)
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; path {
	case "":
		fmt.Fprintf(w, "Hello world!")
	case "help":
		fmt.Fprintf(w, helpHTML)
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

		if r == "all" {
			pRange = makeRange(1, 65535)
			return pRange
		}

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

	sort.Ints(pRange)
	pRange = unique(pRange)
	return pRange
}

func stringPortList(p []int) []string {
	var sPorts []string
	for _, port := range p {
		sPorts = append(sPorts, strconv.Itoa(port))
	}
	return sPorts
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func connect(host string, ports []string) {
	found := false
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("open", net.JoinHostPort(host, port))
			found = true
		}
	}
	if found == false {
		fmt.Println("No connections established")
	}
}

func scan(host string, ports []string) {
	ports = stringPortList(makePortList(ports))
	connect(host, ports)
}
