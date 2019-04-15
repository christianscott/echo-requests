package main

import (
	"bufio"
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

	requests := make(chan *http.Request)
	go func() {
		for r := range requests {
			printRequest(r)
		}
	}()

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests <- r

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "ok")
	})

	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

func printRequest(r *http.Request) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	fmt.Fprintf(f, "%s %s\n", r.Method, r.URL.Path)

	for name, value := range r.Header {
		fmt.Fprintf(f, "%s: %s\n", name, value)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err == nil && len(b) > 0 {
		fmt.Fprintf(f, "%s\n", string(b[:]))
	}

	fmt.Fprint(f, "\n")
}
