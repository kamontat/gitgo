#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by v3.0.2
# link (https://github.com/Template-generator/script-genrating/tree/v3.0.2)

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

#/ -----------------------------------
#/ Description:  ...
#/ How to:       ...
#/               ...
#/ Option:       --help | -h | -? | help | h | ?
#/                   > show this message
#/               --version | -v | version | v
#/                   > show command version
#/ -----------------------------------
#/ Create by:    Kamontat Chantrachirathunrong <kamontat.c@hotmail.com>
#/ Since:        16/08/2018
#/ -----------------------------------
#/ Error code    1      -- error
#/ -----------------------------------
#/ Known bug:    ...
#/ -----------------------------------
#// Version:      0.0.1   -- description
#//               0.0.2b1 -- beta-format
#//               0.0.2a1 -- alpha-format

export APPNAME="gitgo"

mkdir ./build/ &>/dev/null

process() {
	local cmd="$1"
	local title="$2"
	shift 2
	IFS=" " read -r -a array <<<"$@"
	duration=$#

	curr_bar=0
	for ((elapsed = 1; elapsed <= duration; elapsed++)); do
		barsize=$(($(tput cols) - 17))
		unity=$((barsize / duration))
		increment=$((barsize % duration))
		skip=$((duration / (duration - increment)))
		# Elapsed
		((curr_bar += unity))

		if [[ $increment -ne 0 ]]; then
			if [[ $skip -eq 1 ]]; then
				[[ $((elapsed % (duration / increment))) -eq 0 ]] && ((curr_bar++))
			else
				[[ $((elapsed % skip)) -ne 0 ]] && ((curr_bar++))
			fi
		fi

		[[ $elapsed -eq 1 && $increment -eq 1 && $skip -ne 1 ]] && ((curr_bar++))
		[[ $((barsize - curr_bar)) -eq 1 ]] && ((curr_bar++))
		[[ $curr_bar -lt $barsize ]] || curr_bar=$barsize

		printf "%-7s |" "$title"

		# Exection
		"$cmd" "${array[elapsed - 1]}"

		# Progress
		for ((filled = 0; filled <= curr_bar; filled++)); do
			printf "#"
		done

		# Remaining
		for ((remain = curr_bar; remain < barsize; remain++)); do
			printf " "
		done

		# Percentage
		printf "| %s%%" $(((elapsed * 100) / duration))

		# Return
		printf '\r'
	done
	echo
}

build() {
	local input="${1}"
	local os="${input%%,*}"
	local arch="${input##*,}"
	local ext=".sh"

	[[ "$os" == "windows" ]] && ext=".exe"
	[[ "$os" == "darwin" ]] && ext=".sh"

	local filename="./build/${APPNAME}.${os}.${arch}${ext}"
	env GOOS="$os" GOARCH="$arch" go build -o "$filename" &>/dev/null
}

process build "MacOS" "darwin,386" "darwin,amd64"

process build "Window" "windows,386" "windows,amd64"

process build "Linux" "linux,386" "linux,amd64" "linux,arm" "linux,arm64" "solaris,amd64"

process build "bsd" "openbsd,386" "openbsd,amd64" "openbsd,arm" "freebsd,386" "freebsd,amd64" "freebsd,arm" "netbsd,386" "netbsd,amd64" "netbsd,arm"
