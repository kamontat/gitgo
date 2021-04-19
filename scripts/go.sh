#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by create-script-file v4.0.1
# link (https://github.com/Template-generator/create-script-file/tree/v4.0.1)

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

scripts="$(dirname "$0")"
go="$scripts/raw-go.sh"
cmd="$scripts/raw-cmd.sh"

modules=(
  "utils"
  "config"
  "git"
  "prompt"
  "core"
  "cli"
)

exec_go() {
  echo "go" "$@"
  $go "$@"
}

exec_cmd() {
  echo "$@"
  $cmd "$@"
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
elif [[ "$command" == "coverage" ]]; then
  module_filename="coverage.out"

  for module in "${modules[@]}"; do
    if ! exec_go "$module" test -cover "-coverprofile=${module_filename}" -covermode=atomic; then
      exit 1
    fi
  done
elif [[ "$command" == "html" ]]; then
  module_filename="coverage.out"

  for module in "${modules[@]}"; do
    if ! exec_go "$module" tool cover -html="${module_filename}"; then
      exit 1
    fi
  done
elif [[ "$command" == "all" ]]; then
  for module in "${modules[@]}"; do
    if ! exec_go "$module" "${args[@]}"; then
      exit 1
    fi
  done
elif [[ "$command" == "publishLocal" ]]; then
  exec_go cli build || exit $?
  cp "$PWD/cli/cli" "/usr/local/bin/gitgo-next"
elif [[ "$command" == "raw" ]]; then
  exec_go "${args[@]}"
else
  echo "command not found $command '${args[*]}'" >&2 || exit 1
fi
