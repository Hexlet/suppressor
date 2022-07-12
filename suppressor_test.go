package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuppressorPassPassed(t *testing.T) {
	expected, _ := ioutil.ReadFile("fixtures/pass_passed.txt")
	out, code := checkCommand("pass", "echo", []string{"true"})

	require.Equal(t, code, 0)
	require.Equal(t, string(expected), out.String())
}

func TestSuppressorPassFailed(t *testing.T) {
	expected, _ := ioutil.ReadFile("fixtures/pass_failed.txt")
	out, code := checkCommand("pass", "false", nil)

	require.Equal(t, code, 1)
	require.Equal(t, string(expected), out.String())
}

func TestSuppressorFailFailed(t *testing.T) {
	expected, _ := ioutil.ReadFile("fixtures/fail_failed.txt")
	out, code := checkCommand("fail", "false", nil)

	require.Equal(t, code, 0)
	require.Equal(t, string(expected), out.String())
}

func TestSuppressorFailPassed(t *testing.T) {
	expected, _ := ioutil.ReadFile("fixtures/fail_passed.txt")
	out, code := checkCommand("fail", "echo", []string{"true"})

	require.Equal(t, code, 1)
	require.Equal(t, string(expected), out.String())
}
