package main

import (
	run "github.com/bahodurnazarov/CatFacts/internal/bot"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	go run.Route()
	handler.Listen()
}
