package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 4000, "server port number")

	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	db := NewDB()
	server := NewServer(db)

	log.Fatal(http.ListenAndServe(addr, server))
}
