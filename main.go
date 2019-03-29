package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: 	%s [port]", os.Args[0])
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		printRequest(r)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
	fmt.Printf("listening on %s\n", port)
}

func printRequest(r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	for name, value := range r.Header {
		fmt.Printf("%s: %s\n", name, value)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err == nil && len(b) > 0 {
		fmt.Printf("%s\n", string(b[:]))
	}

	fmt.Print("\n")
}
