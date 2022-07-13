package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var (
	red   = color.New(color.FgHiRed)
	green = color.New(color.FgGreen)
)

func init() {
	red.EnableColor()
	green.EnableColor()
}

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
			b.WriteString(" \n")
			b.WriteString(red.Sprint(" Expected tests to pass, but error occurred. See output above. "))
			b.WriteString(" \n")
			return &b, 1
		}

		b.WriteString(green.Sprint(" Expected tests to pass, recieved tests passed. "))
		b.WriteString(" \n")
	}

	if mode == "fail" {
		writeOutput(string(stdout), &b)

		if err == nil {
			b.WriteString(red.Sprint(" Expected tests to fail, but they passed. See output above. "))
			b.WriteString(" \n")
			return &b, 1
		}

		writeOutput(err.Error(), &b)
		b.WriteString(" \n")
		b.WriteString(green.Sprint(" Expected tests to fail, recieved tests failed. "))
		b.WriteString(" \n")
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
