package main

import (
	"fmt"
	"github.com/bahodurnazarov/CatFacts/internal/Job"
	"github.com/bahodurnazarov/CatFacts/internal/handler"
)

func main() {
	TestGit()
	go Job.Route()
	handler.Listen()
}

func TestGit() {
	fmt.Println("hello from git")
}
