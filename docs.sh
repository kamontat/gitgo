#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by v3.0.2
# link (https://github.com/Template-generator/script-genrating/tree/v3.0.2)

# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

#/ -----------------------------------
#/ Description:  generate docs of gitgo command to several OS
#/ -----------------------------------
#/ Create by:    Kamontat Chantrachirathunrong <kamontat.c@hotmail.com>
#/ Since:        16/08/2018
#/ -----------------------------------

printf "Generate godoc     | "
godoc -html . >docs/godoc.html && echo "Completed" || echo "Failure"

printf "Generate changelog | "
git changelog --all --prune-old --stdout >docs/changelog.md && echo "Completed" || echo "Failure"

printf "Generate summary   | "
echo "# Summary" >docs/summary.txt &&
	git summary >>docs/summary.txt &&
	echo "## Lines report" >>docs/summary.txt &&
	echo >>docs/summary.txt &&
	git summary --line | grep -A 100 "lines" >>docs/summary.txt &&
	echo "Completed" ||
	echo "Failure"
