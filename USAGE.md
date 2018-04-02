# Gitgo

## Setting

<!-- open config with $EDITOR -->
git config|g
<!-- open config location -->
git config|g location
<!-- set config by key and value -->
git config|g --key <key> --value <value>
git config|g <key> <value>
<!-- get config by key -->
git config|g --key <key>
git config|g <key>

## Initial (0.0.1-alpha.2)

<!-- create .git folder -->
git init|i [--force|-f]

## Destroy (0.0.1-alpha.3)

<!-- delete .git folder -->
git destroy|d [--force|-f]

## GitAdd (0.0.1-alpha.2)

<!-- add every changes -->
git add|a all|--all
<!-- add only input folders|files -->
git add|a <folder|file>...

## Commit

<!-- open commitment prompt -->
git commit|c
<!-- commit with default inital commit -->
git commit|c [emoji|moji|e|m|text|t] initial|i
<!-- change commit with emoji -->
git commit|c emoji|moji|e|m [<message>]
<!-- change commit with text -->
git commit|c text|t [<title>] [<message>]

## Push

<!-- push code and create upstream -->
git push|p set|s <repository> <link> [<branch>]
<!-- push code -->
git push|p [--force|-f] [<repository>] [<branch>]

## Pull

<!-- pull code -->
git pull|P [--force|-f] [<repository>] [<branch>]