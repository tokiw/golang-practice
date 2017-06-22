package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "specify item name\n")
		return
	}
	if p, hasItem := db[item]; hasItem {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Already create item: %q\n", item)
		fmt.Fprintf(w, "%s\n", p)
	} else {
		d, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(d)
		fmt.Fprintf(w, "Created item: %q\n", item)
		fmt.Fprintf(w, "%s\n", price)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, hasItem := db[item]; hasItem {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not find item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, hasItem := db[item]; hasItem {
		d, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(d)
		fmt.Fprintf(w, "Updated item %s: %f\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not find item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, hasItem := db[item]; hasItem {
		delete(db, item)
		fmt.Fprintf(w, "Deleted item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not found item: %q\n", item)
	}
}
