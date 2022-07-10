package main

import (
	"L2/develop/dev11/pkg/server"
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "number of port")
	flag.Parse()
	mux := http.NewServeMux()
	handler := server.NewHandler()
	handler.Routing(mux)
	loggedRouter := server.Logging(mux)
	log.Printf("Server is running on %s port\n", *port)
	err := http.ListenAndServe("localhost:"+*port, loggedRouter)
	if err != nil {
		log.Fatalln(err)
	}
}
