#!/usr/bin/env bash

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

#/ -------------------------------------------------
#/ Description:  ...
#/ Create by:    ...
#/ Since:        ...
#/ -------------------------------------------------
#/ Version:      0.0.1  -- description
#/               0.0.2b -- beta-format
#/ -------------------------------------------------
#/ Error code    1      -- error
#/ -------------------------------------------------
#/ Bug:          ...
#/ -------------------------------------------------

cd "$(dirname "$0")" || exit 1
# cd "$(dirname "$(realpath "$0")")"

help() {
	cat "./go-build.sh" | grep "^#/" | tr -d "#/ "
}

_build() {
	local name="$1"
	local os="$2"
	local ext=""
	[[ "$os" == "windows" ]] && ext=".exe"
	shift 2
	for arch in "$@"; do
		printf 'Currently build: %-8s - %s\n' "$os" "$arch"
		local filename="${name}.${os}.${arch}${ext}"
		env GOOS="$os" GOARCH="$arch" go build -o "$filename"
	done
}

build_for_window() {
	local os="windows"
	local arches=("386" "amd64")
	_build "$1" "$os" "${arches[@]}"
}

build_for_mac() {
	local os="darwin"
	local arches=("386" "amd64") #  "arm" "arm64"
	_build "$1" "$os" "${arches[@]}"
}

build_for_linux() {
	local arches4=("386" "amd64" "arm" "arm64")
	local arches3=("386" "amd64" "arm")
	_build "$1" "linux" "${arches4[@]}"

	_build "$1" "openbsd" "${arches3[@]}"
	_build "$1" "freebsd" "${arches3[@]}"
	_build "$1" "netbsd" "${arches3[@]}"

	_build "$1" "solaris" "amd64"
}

APP="gitgo"

build_for_window "$APP"

build_for_mac "$APP"

build_for_linux "$APP"

mkdir bin &>/dev/null
mv gitgo* bin/
