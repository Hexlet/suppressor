#!/usr/bin/env python3.5

"""Usage: suppressor [-h] pass COMMAND
          suppressor [-h] fail COMMAND

Arguments:
  EXPECTED_STATUS        expect pass or fails comand status
  COMMAND                Command to execute
Options:
  -h --help

"""

import sys
from subprocess import run, PIPE
from docopt import docopt

if __name__ == '__main__':
    ARGS = docopt(__doc__)
    result = run(ARGS["COMMAND"], shell=True, stdout=PIPE, stderr=PIPE)

    if result.returncode != 0 and ARGS["pass"]:
        print(result.stderr.decode(), file=sys.stderr)
        exit(1)
    elif result.returncode == 0 and ARGS["fail"]:
        print(result.stdout.decode(), file=sys.stderr)
        exit(1)

