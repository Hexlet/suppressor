package main

import (
	"fmt"
        "strings"
	"os"
	"os/exec"
)

func prepare_output(output string) string {
  splitted_list := strings.Split(output, "\n")
  joined_str := strings.Join(splitted_list, "\n#\t")
  return "#\t" + joined_str + "\n";
}

func main() {

  mode := os.Args[1]

  cmd := exec.Command("/bin/bash", "-c", os.Args[2])

  stdout, err := cmd.Output()

  fmt.Fprintf(os.Stdout, prepare_output(string(stdout)))

  if (mode == "pass") {
    if ( err != nil ) {
      fmt.Fprintf(os.Stdout,"%s \n", prepare_output(err.Error()))
      fmt.Fprintf(os.Stdout,"\033[91m Expected tests to pass, but error occurred. See output above. \033[0m \n")
      os.Exit(1)
    }
  }

  if (mode == "fail") {
    if ( err == nil ) {
      fmt.Fprintf(os.Stdout, "\033[91m Expected tests to fail, but they passed. See output above. \033[0m \n")
      os.Exit(1)
    }
  }
}
