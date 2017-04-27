package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	var cycles = 5
	cyclesStr := r.FormValue("cycles")
	if len(cyclesStr) != 0 {
		var err error
		cycles, err = strconv.Atoi(cyclesStr)
		if err != nil || cycles < 1 {
			fmt.Fprintf(w, "Cycles must be specified by natrual number: %d\n", cycles)
			return
		}
	}
	lissajous(w, cycles)
}
