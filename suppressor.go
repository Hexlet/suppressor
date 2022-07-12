package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func writeOutput(output string, b *strings.Builder) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	scanner.Split(bufio.ScanRunes)

	b.WriteString("#\t")

	for scanner.Scan() {
		if scanner.Text() == "\n" {
			b.WriteString("\n#\t")
			continue
		}

		b.Write(scanner.Bytes())
	}

	b.WriteString("\n")
}

func checkCommand(mode, name string, args []string) (*strings.Builder, int) {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.CombinedOutput()

	var b strings.Builder

	if mode == "pass" {
		writeOutput(string(stdout), &b)

		if err != nil {
			writeOutput(err.Error(), &b)
			b.WriteString(" \n\033[91m Expected tests to pass, but error occurred. See output above. \033[0m \n")
			return &b, 1
		}

		b.WriteString("\u001b[32m Expected tests to pass, recieved tests passed. \033[0m \n")
	}

	if mode == "fail" {
		writeOutput(string(stdout), &b)

		if err == nil {
			b.WriteString("\033[91m Expected tests to fail, but they passed. See output above. \033[0m \n")
			return &b, 1
		}

		writeOutput(err.Error(), &b)
		b.WriteString(" \n\u001b[32m Expected tests to fail, recieved tests failed. \033[0m \n")
	}

	return &b, 0
}

func main() {
	mode := os.Args[1]
	args := strings.Split(os.Args[2], " ")

	out, code := checkCommand(mode, args[0], args[1:])

	fmt.Print(out)
	os.Exit(code)
}
