package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	xmin := req.Form.Get("xmin")
	limit, err := strconv.ParseFloat(xmin, 64)
	if err != nil {
		log.Printf("Error parsing limit: %s", err)
		return
	}

	ymin := req.Form.Get("ymin")
	dryRun, _ := strconv.ParseFloat(ymin, 64)
	fmt.Fprintf(w, "hello world. xmin: %f, ymin: %f\n", limit, dryRun)
}
