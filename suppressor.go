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

  args := strings.Split(os.Args[2], " ")
  cmd := exec.Command(args[0], args[1:]...)

  stdout, err := cmd.CombinedOutput()


  if (mode == "pass") {
    if ( err != nil ) {
      fmt.Fprintf(os.Stdout, prepare_output(string(stdout)))
      fmt.Fprintf(os.Stdout,"%s \n", prepare_output(fmt.Sprint(err)))
      fmt.Fprintf(os.Stdout,"\033[91m Expected tests to pass, but error occurred. See output above. \033[0m \n")
      os.Exit(1)
    } else {
      fmt.Fprintf(os.Stdout,"\u001b[32m Expected tests to pass, recieved tests passed\033[0m \n")
    }
  }

  if (mode == "fail") {
    if ( err == nil ) {
      fmt.Fprintf(os.Stdout, prepare_output(string(stdout)))
      fmt.Fprintf(os.Stdout, "\033[91m Expected tests to fail, but they passed. See output above. \033[0m \n")
      os.Exit(1)
    } else {
      fmt.Fprintf(os.Stdout,"\u001b[32m Expected tests to fail, recieved tests failed\033[0m \n")
    }
  }
}
