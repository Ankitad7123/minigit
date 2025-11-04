package commands

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func Add(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	hash := fmt.Sprintf("%x", sha1.Sum(data))
	err = os.WriteFile(".minigit/objects/"+hash, data, 0644)
	if err != nil {
		fmt.Println("Error saving snapshot:", err)
		return
	}
	os.WriteFile(".minigit/stage", []byte(hash+" "+filename+"\n"), 0644)
	fmt.Println("Added file:", filename)

}
