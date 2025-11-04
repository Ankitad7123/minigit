package commands

import (
	"fmt"
	"os"
	"time"
)

func Commits(message string) {
	stageData, err := os.ReadFile(".minigit/stage")
	if err != nil {
		fmt.Println("Nothing to commit or staging missing.")
		return
	}

	commitTime := time.Now().Format(time.RFC1123)
	commitContent := fmt.Sprintf("Message: %s\n time: %s\n Files:\n%s\n", message, commitTime, string(stageData))

	commitHash := fmt.Sprintf("%x", time.Now().UnixNano())
	commitPath := ".minigit/commits/" + commitHash

	err = os.WriteFile(commitPath, []byte(commitContent), 0644)
	if err != nil {
		fmt.Println("Commit failed:", err)
		return
	}

	os.Remove(".minigit/stage")

	fmt.Println("Commited :", message)

}
