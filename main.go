package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command:", err)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		cmd := exec.Command("sh", "-c", input)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Command execution failed:", err)
		}
	}
}
