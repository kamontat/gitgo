#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by create-script-file v4.0.1
# link (https://github.com/Template-generator/create-script-file/tree/v4.0.1)

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

scripts="$(dirname "$0")"
go="$scripts/raw-go.sh"

exec_go() {
  echo "go" "$@"
  $go "$@"
}

command="$1"
shift
args=("$@")

if [[ "$command" == "help" ]]; then
  echo """
# ./scripts/go.sh [command] [...params]

1. run          - called 'go run main.go'
2. start        - called 'go build && ./cli'
3. new \$1 \$2    - called 'go mod init \$2' on \$1 module
4. all \$1       - called 'go \$1' on every modules
5. publishLocal - called 'go build && cp cli /usr/local/bin/gitgo-next'
6. raw \$@       - called 'go \$2' on \$1 module
"""

elif [[ "$command" == "run" ]]; then
  exec_go cli run main.go "${args[@]}" || exit $?
elif [[ "$command" == "start" ]]; then
  exec_go cli build || exit $?
  "$PWD/cli/cli" "${args[@]}"
elif [[ "$command" == "new" ]]; then
  exec_go "$1" mod init "$2"
elif [[ "$command" == "all" ]]; then
  exec_go cli "${args[@]}" &&
    exec_go core "${args[@]}" &&
    exec_go prompt "${args[@]}" &&
    exec_go git "${args[@]}" &&
    exec_go config "${args[@]}" &&
    exec_go utils "${args[@]}" || exit $?
elif [[ "$command" == "publishLocal" ]]; then
  exec_go cli build || exit $?
  cp "$PWD/cli/cli" "/usr/local/bin/gitgo-next"
elif [[ "$command" == "raw" ]]; then
  exec_go "${args[@]}"
else
  echo "command not found $command '${args[*]}'" >&2 || exit 1
fi
