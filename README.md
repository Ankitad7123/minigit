MiniGit (Made in Go)
A very simple version of Git, built from scratch using Golang, to understand how version control really works under the hood.
This project helps you learn how Git stores, stages, and commits files — step by step.


# Step 1: Initialize MiniGit
go run main.go init

# Step 2: Add a file
go run main.go add file.txt

# Step 3: Commit changes
go run main.go commit -m "Added first file"

# Step 4: See all commits
go run main.go log


Why This Project Is Useful
Teaches how Git actually works instead of just using it
You’ll understand hashing, staging, and commit history clearly
Great resume project to show backend + CLI + file handling skills
100% local — no network, no databases, just pure Go logic
