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

## Commit (1.0.0-beta.1)

<!-- open commitment prompt -->
git commit|c [--add|-a] [--all|-A]
<!-- commit with default inital commit -->
git commit|c [emoji|moji|e|m|text|t] initial|i
<!-- change commit with emoji -->
<!-- 
[key]: title
message
-->
git commit|c emoji|moji|e|m [--add|-a] [--all|-A] [--key|k <key>] [--title|t <title>] [<message>]
<!-- change commit with text -->
<!-- 
[key]: title
message
-->
git commit|c text|t [--add|-a] [--all|-A] [--key|k <key>] [--title|t <title>] [<message>]

## Push (0.0.1-alpha.6)

<!-- push code and create upstream -->
git push|p set|s [--repo <repository>] [--branch <branch>] <link>
<!-- push code -->
git push|p [--force|-f] [--repo <repository>] [<branch>...] 

## Pull (0.0.1-alpha.7)

<!-- pull code -->
git pull|P [--force|-f] [<repository>] [<branch>...]