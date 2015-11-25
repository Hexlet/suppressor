#!/usr/bin/env python3

"""Usage: suppressor [-h] pass COMMAND
          suppressor [-h] fail COMMAND

Arguments:
  EXPECTED_STATUS        expect pass or fails comand status
  COMMAND                Command to execute
Options:
  -h --help

"""

import sys
from subprocess import Popen, PIPE
from docopt import docopt

def main():
    docopt_args = docopt(__doc__)
    p = Popen(docopt_args["COMMAND"], shell=True, stdout=PIPE, stderr=PIPE)
    output, err = p.communicate()
    if docopt_args["pass"]:
        if p.returncode != 0:
            print(prepare_out(output), file=sys.stderr)
            print(prepare_out(err), file=sys.stderr)
            print("\033[91m Expected tests to pass, but error occurred. See output above. \033[0m", file=sys.stdout)
            exit(1)

    if docopt_args["fail"]:
        if p.returncode == 0:
            print(prepare_out(output), file=sys.stderr)
            print("\033[91m Expected tests to fail, but they passed. See output above. \033[0m", file=sys.stdout)
            exit(1)

def prepare_out(output):
    splitted_list = output.decode().split("\n")
    joined_str = "\n#\t".join(item for item in splitted_list)
    return "#\t" + joined_str

