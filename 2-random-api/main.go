package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, rand.Intn(6)+1)
	})

	server := &http.Server{
		Addr:    ":8091",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
