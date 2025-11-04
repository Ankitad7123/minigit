package commands

import (
	"fmt"
	"os"
)

func Init() {
	if _, err := os.Stat(".minigit"); !os.IsNotExist(err) {
		fmt.Println("Alredy initialized")
		return
	}

	os.Mkdir(".minigit", 0755)
	os.Mkdir(".minigit/objects", 0755)
	os.Mkdir(".minigit/commits", 0755)
	fmt.Println("Initialized empty mini-git repository in .minigit/")

}
