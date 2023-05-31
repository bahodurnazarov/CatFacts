package main

import (
	"fmt"
	"github.com/bahodurnazarov/CatFacts/internal/Job"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	PingPong()
	go Job.Route()
	handler.Listen()
}

func PingPong() {
	fmt.Println("ping-pong")
}
