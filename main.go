package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {

	// /home
	f, err := os.Open("static" + r.URL.Path)

	if err != nil {
		log.Println(err)
		return
	}

	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Add("Content-Type", "text/css")
	}

	io.Copy(w, f)
}

func main() {

	http.HandleFunc("/", StaticHandler)

	http.ListenAndServe(":3000", nil)
}
