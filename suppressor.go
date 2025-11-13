package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

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

type prefixedWriter struct {
	mu           sync.Mutex
	out          io.Writer
	prefixActive bool
}

func newPrefixedWriter(out io.Writer) *prefixedWriter {
	return &prefixedWriter{out: out}
}

func (w *prefixedWriter) ensurePrefixLocked() error {
	if !w.prefixActive {
		if _, err := w.out.Write([]byte("#\t")); err != nil {
			return err
		}
		w.prefixActive = true
	}
	return nil
}

func (w *prefixedWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, b := range p {
		if err := w.ensurePrefixLocked(); err != nil {
			return 0, err
		}

		if _, err := w.out.Write([]byte{b}); err != nil {
			return 0, err
		}

		if b == '\n' {
			w.prefixActive = false
		}
	}

	return len(p), nil
}

func (w *prefixedWriter) finalize() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.ensurePrefixLocked(); err != nil {
		return
	}

	if _, err := w.out.Write([]byte("\n")); err != nil {
		return
	}

	w.prefixActive = false
}

func (w *prefixedWriter) printString(text string) {
	if text == "" {
		w.finalize()
		return
	}

	if _, err := w.Write([]byte(text)); err != nil {
		return
	}

	w.finalize()
}

func checkCommand(mode, name string, args []string) int {
	cmd := exec.Command(name, args...)

	stream := newPrefixedWriter(os.Stdout)
	cmd.Stdout = stream
	cmd.Stderr = stream

	err := cmd.Run()
	stream.finalize()

	hasError := err != nil

	if hasError {
		stream.printString(err.Error())
	}

	switch mode {
	case "pass":
		if hasError {
			fmt.Print(" \n")
			fmt.Print(red.Sprint(" Expected tests to pass, but error occurred. See output above. "))
			fmt.Print(" \n")
			return 1
		}
		fmt.Print(green.Sprint(" Expected tests to pass, recieved tests passed. "))
		fmt.Print(" \n")
		return 0
	case "fail":
		if err == nil {
			fmt.Print(red.Sprint(" Expected tests to fail, but they passed. See output above. "))
			fmt.Print(" \n")
			return 1
		}
		fmt.Print(" \n")
		fmt.Print(green.Sprint(" Expected tests to fail, recieved tests failed. "))
		fmt.Print(" \n")
		return 0
	default:
		fmt.Println("Unknown mode:", mode)
		return 1
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <mode> <command>")
		os.Exit(1)
	}

	mode := os.Args[1]
	args := strings.Split(os.Args[2], " ")

	code := checkCommand(mode, args[0], args[1:])
	os.Exit(code)
}
