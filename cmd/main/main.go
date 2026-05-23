package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SadraSamadi/git-proxy/internal/app"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	switch os.Args[1] {
	case "save":
		save()
	case "list":
		list()
	case "remove":
		remove()
	case "use":
		use()
	case "current":
		current()
	case "unset":
		unset()
	default:
		usage()
	}
}

func save() {
	if len(os.Args) != 4 {
		usage()
	}
	key := os.Args[2]
	value := os.Args[3]
	err := app.Save(key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func list() {
	proxies, err := app.List()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range proxies {
		fmt.Printf("%-20s | %s\n", k, v)
	}
}

func remove() {
	if len(os.Args) != 3 {
		usage()
	}
	key := os.Args[2]
	err := app.Remove(key)
	if err != nil {
		log.Fatal(err)
	}
}

func use() {
	if len(os.Args) != 3 {
		usage()
	}
	key := os.Args[2]
	err := app.Use(key)
	if err != nil {
		log.Fatal(err)
	}
}

func current() {
	value, err := app.Current()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(value)
}

func unset() {
	err := app.Unset()
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	msg := `Git Proxy Manager

Usage:
  git-proxy save <key> <value>   Save a proxy  (e.g. save test socks5://127.0.0.1:1080)
  git-proxy list                 List all saved proxies
  git-proxy remove <key>         Remove a saved proxy
  git-proxy use <key>            Configure git to use a saved proxy
  git-proxy current              Show the current git proxy
  git-proxy unset                Unset the current git proxy
`
	fmt.Print(msg)
	os.Exit(0)
}
