package main

import (
	"fmt"
	"github.com/bahodurnazarov/CatFacts/internal/Job"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	Test("Github")

	go Job.Route()
	handler.Listen()
}

func Test(name string) {
	fmt.Printf("Hello, %s", name)
}
