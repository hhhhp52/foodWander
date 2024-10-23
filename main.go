package main

import (
	"fmt"
	"foodWander/src/routes"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	fmt.Println("Starting server on port 8080...")

	router1 := routes.Router()
	router2 := routes.Router()

	s1 := &http.Server{
		Addr:           ":8080",
		Handler:        router1,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s2 := &http.Server{
		Addr:           ":8081",
		Handler:        router2,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	servers := []*http.Server{s1, s2}

	for _, s := range servers {
		g.Go(func() error {
			return s.ListenAndServe()
		})
	}
}
