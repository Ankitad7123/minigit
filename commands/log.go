package commands

import (
	"fmt"
	"os"
)

func Log() {

	files, _ := os.ReadDir(".minigit/commits")
	if len(files) == 0 {
		fmt.Println("No commits yet.")
		return
	}
	for _, file := range files {
		data, _ := os.ReadFile(".minigit/commits/" + file.Name())
		fmt.Println("----")
		fmt.Println(string(data))

	}
}
