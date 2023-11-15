package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	log.Println(stdout)
}
