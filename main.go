package main

import (
	"fmt"
	"os"

	"github.com/Ankitd7123/minigit/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [init|add|commit|log]")
		return
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Specify file to add.")
			return
		}
		commands.Add(os.Args[2])
	case "commit":
		if len(os.Args) < 4 || os.Args[2] != "-m" {
			fmt.Println("Usage: go run main.go commit -m \"message\"")
			return
		}
		commands.Commits(os.Args[3])
	case "log":
		commands.Log()
	default:
		fmt.Println("Unknown command.")
	}

}
