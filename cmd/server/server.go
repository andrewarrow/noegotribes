package main

import (
	"flag"
	"log"
	"net/http"
	"noegotribes/internal/handlers"
	"noegotribes/internal/memstore"
	"path/filepath"
)

func main() {
	baseDir := flag.String("base", "", "Base directory to look for files, default uses current directory")
	fullBaseDir, _ := filepath.Abs(*baseDir)

	mux := http.NewServeMux()
	frhl := handlers.NewFrontendHandler(fullBaseDir)
	mem := memstore.NewMemStore()

	tchl := handlers.NewTacoStoreAPIHandler(mem)
	cahl := handlers.NewCartAPIHandler(mem)
	//fs := http.FileServer(http.Dir("/Users/aa/noegotribes.com/assets"))
	//mux.Handle("/assets/", fs)
	mux.Handle("/", frhl)
	mux.Handle("/api/taco-list", tchl)
	mux.Handle("/api/cart", cahl)

	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	log.Fatal(http.ListenAndServe(l, mux))
}
