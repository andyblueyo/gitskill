package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/andyblueyo/gitskill/server/gateway/handlers"
)

const repoPath = "/repos"

func main() {

	//and use that as the address this server will listen on
	addr := os.Getenv("ADDR")
	//if not set, default to ":443", which means listen for
	//all requests to all hosts on port 443
	if len(addr) == 0 {
		addr = ":80"
	}

	token := os.Getenv("GHTOKEN")
	if len(token) == 0 {
		log.Fatalf("you must have GHTOKEN set in your env")
	}

	mux := http.NewServeMux()

	//create a new handlers.CityHandler struct
	//since that is in a different package, use the
	//package name as a prefix, and import the package above
	repoHandler := &handlers.RepoHandler{token}
	//add the handler to the mux using .Handle() instead
	//of .HandleFunc(). The former is used for structs that
	//implement the http.Handler interface, while the latter
	//is used for simple functions that conform to the
	//http.HandlerFunc type.
	//see https://drstearns.github.io/tutorials/goweb/#sechandlers
	mux.Handle(repoPath, repoHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
