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

go install

gitgo --version

printf "Check version; press <enter> to next"
# shellcheck disable=SC2034
read -r ans

git status
printf "Check status; press <enter> to next"
read -r ans

tag="v$(gitgo --version | sed -e 's/gitgo version //g')"
git add -A
git commit -m "chore(release): version: $tag"

printf "create tag %s; press <enter> to next" "$tag"
# shellcheck disable=SC2034
read -r ans

git tag "$tag"

printf "build project; press <enter> to next"
# shellcheck disable=SC2034
read -r ans

./build.sh
./docs.sh

git add -A
git commit -m "[doc] Update documents (docs)"

git push &&
  git push --tag

printf "create release version; press <enter> to next"
# shellcheck disable=SC2034
read -r ans

# command -v "hub" &>/dev/null &&
#   hub pull-request -m "create pull request to release latest version $tag" -b master

command -v "hub" &>/dev/null &&
  hub release create -m "Release newest version" -a ./build/gitgo.darwin.386.sh -a ./build/gitgo.darwin.amd64.sh -a ./build/gitgo.freebsd.386.sh -a ./build/gitgo.freebsd.amd64.sh -a ./build/gitgo.freebsd.arm.sh -a ./build/gitgo.linux.386.sh -a ./build/gitgo.linux.amd64.sh -a ./build/gitgo.linux.arm.sh -a ./build/gitgo.linux.arm64.sh -a ./build/gitgo.netbsd.386.sh -a ./build/gitgo.netbsd.amd64.sh -a ./build/gitgo.netbsd.arm.sh -a ./build/gitgo.windows.386.exe -a ./build/gitgo.windows.amd64.exe "$tag"
