package main

import (
	"fmt"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	handler := &PlayerServer{&InMemoryPlayerStore{}}
	fmt.Print("Server started at :5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
