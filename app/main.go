package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := &PlayerServer{NewInMemoryPlayerStore()}
	fmt.Print("Server started at :5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
