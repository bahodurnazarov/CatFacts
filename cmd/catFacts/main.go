package main

import (
	"fmt"
	"github.com/bahodurnazarov/CatFacts/internal/Job"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	Ping()
	go Job.Route()
	handler.Listen()
}

func Ping() {
	fmt.Println("ping-pong")
}
