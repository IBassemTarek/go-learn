package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cmd/docker_build/main.go <image_name> <tag>")
		os.Exit(1)
	}

	imageName := os.Args[1]
	tag := os.Args[2]

	cmd := exec.Command("docker", "build", "--tag", imageName+":"+tag, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error building Docker image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully built Docker image: %s:%s\n", imageName, tag)
}
