package main

import (
	"log"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

func job(name string, fn func() error) {
	defer wg.Done()
	log.Println("[job]", name, "start")
	if err := fn(); err != nil {
		log.Println("[job]", name, "error:", err)
	}
	log.Println("[job]", name, "done")
}

func main() {
	wg.Add(2)

	go job("lint", func() error { return exec.Command("golangci-lint", "run").Run() })
	go job("test", func() error { return exec.Command("go", "test", "./...").Run() })

	wg.Wait()
}
