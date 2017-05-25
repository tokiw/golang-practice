package main

import (
	"image/png"
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
	w.Header().Set("Content-Type", "image/png")
	x, errX := strconv.ParseFloat(r.FormValue("x"), 64)
	y, errY := strconv.ParseFloat(r.FormValue("y"), 64)
	z, errZ := strconv.ParseFloat(r.FormValue("z"), 64)

	if errX != nil {
		x = 0.0
	} else if errY != nil {
		y = 0.0
	} else if errZ != nil {
		z = 2.0
	}

	png.Encode(w, img(x, y, z))
}
