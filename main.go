// Launch microservice server- main.go
package main

import (
	"log"

	"github.com/elsaimo/github-action/tree/main/microservice"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
