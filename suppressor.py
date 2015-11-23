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

def main(docopt_args):
    p = Popen(ARGS["COMMAND"], shell=True, stdout=PIPE, stderr=PIPE)
    output, err = p.communicate()
    if docopt_args["pass"]:
        if p.returncode != 0:
            print(err.decode(), file=sys.stderr)
            exit(1)

    if docopt_args["fail"]:
        if p.returncode == 0:
            print(output.decode(), file=sys.stderr)
            exit(1)

if __name__ == '__main__':
    ARGS = docopt(__doc__)
    main(ARGS)

