package main

import (
	"log"
	"os"
	"os/exec"
)

func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("[exec] %s %v\n", name, args)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run("docker", "ps")
	run("git", "status")
}
