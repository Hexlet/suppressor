package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	originalStdout := os.Stdout

	r, w, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	var buf bytes.Buffer
	done := make(chan error, 1)

	go func() {
		_, copyErr := io.Copy(&buf, r)
		r.Close()
		done <- copyErr
	}()

	fn()

	require.NoError(t, w.Close())
	require.NoError(t, <-done)

	return buf.String()
}

func runCommand(t *testing.T, mode, name string, args []string) (string, int) {
	t.Helper()

	var code int

	out := captureOutput(t, func() {
		code = checkCommand(mode, name, args)
	})

	return out, code
}

func readFixture(t *testing.T, path string) string {
	t.Helper()

	contents, err := os.ReadFile(path)
	require.NoError(t, err)

	return string(contents)
}

func TestSuppressorPassPassed(t *testing.T) {
	expected := readFixture(t, "fixtures/pass_passed.txt")
	out, code := runCommand(t, "pass", "echo", []string{"true"})

	require.Equal(t, code, 0)
	require.Equal(t, expected, out)
}

func TestSuppressorPassFailed(t *testing.T) {
	expected := readFixture(t, "fixtures/pass_failed.txt")
	out, code := runCommand(t, "pass", "false", nil)

	require.Equal(t, code, 1)
	require.Equal(t, expected, out)
}

func TestSuppressorFailFailed(t *testing.T) {
	expected := readFixture(t, "fixtures/fail_failed.txt")
	out, code := runCommand(t, "fail", "false", nil)

	require.Equal(t, code, 0)
	require.Equal(t, expected, out)
}

func TestSuppressorFailPassed(t *testing.T) {
	expected := readFixture(t, "fixtures/fail_passed.txt")
	out, code := runCommand(t, "fail", "echo", []string{"true"})

	require.Equal(t, code, 1)
	require.Equal(t, expected, out)
}
