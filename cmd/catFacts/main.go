package main

import (
	"fmt"
	"github.com/bahodurnazarov/CatFacts/internal/Job"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	SayHi()
	go Job.Route()
	handler.Listen()
}

func SayHi(name string) {
	fmt.Printf("Hi, %s", name)
}
