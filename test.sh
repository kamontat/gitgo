#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by 2.3.2
# link (https://github.com/Template-generator/script-genrating/tree/2.3.2)

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

if [[ "$1" == "why" ]]; then
	echo " Exit code meaning...
  1.  - unhandle error
  2.  - command not found
  3.  - 'version' error
  4.  - 'initial' error
  5.  - 'destroy' error
  "
	exit 0
fi

run_setup() {
	echo "------- SETUP TEST -------" >&2
	"$@" || exit 1
}

run_test() {
	local header="$1"
	local error_code="$2"
	shift 2

	echo "------- $header -------"
	"$@" || exit "$error_code"
	echo "--------     END     --------"
	echo
}

run_error_test() {
	local header="$1"
	local error_code="$2"
	shift 2

	echo "------- $header -------"
	"$@" && exit "$error_code"
	echo "------- END (code=$?) -------"
	echo
}

run_test_with_input() {
	local header="$1"
	local error_code="$2"
	local input="$3"
	shift 3

	echo "------- $header -------"
	echo "$input" | "$@" || exit "$error_code"
	echo "--------     END     --------"
	echo
}

run_error_test_with_input() {
	local header="$1"
	local error_code="$2"
	local input="$3"
	shift 3

	echo "------- $header -------"
	echo "$input" | "$@" && exit "$error_code"
	echo "------- END (code=$?) -------"
	echo
}

test_location="/tmp/test"

run_test "INSTALL BIN" 1 go install

run_test "SHOW LOCATION" 2 command -v 'gitgo'

run_test "SHOW VERSION" 3 gitgo version

rm -rf "$test_location" 2>/dev/null
mkdir "$test_location" 2>/dev/null
run_test "GLOBAL SETUP" 2 cd "$test_location"

run_test "INITIAL GIT (initial)" 4 gitgo init

run_test_with_input "DESTROY GIT (destroy)" 5 "y" gitgo destroy

run_test "INITIAL GIT (i)" 4 gitgo i

run_test_with_input "DESTROY GIT (d)" 5 "y" gitgo d

run_setup gitgo i

run_test "FORCE DESTROY GIT" 5 gitgo d --force

run_error_test_with_input "ERROR DESTROY" 5 "y" gitgo d
