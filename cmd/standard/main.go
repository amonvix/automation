// cmd/tool/main.go
package main

import (
	"flag"
	"log"
)

func main() {
	path := flag.String("path", ".", "Path to scan")          //declaration of the path to be modified or changed by the automation
	dry := flag.Bool("dry-run", false, "Run without changes") // declaration of the dry test that can be done before the use of the automation
	flag.Parse()                                              // explicit declaration about the use of the flags in CLI

	log.Println("[tool] start")
	log.Printf("[tool] path=%s dry=%v\n", *path, *dry)

}
