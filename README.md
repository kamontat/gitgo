<<<<<<< HEAD
# Gitgo - agile process management

## What is this ?

Basically, this is a `git commit` with `interactive prompt` that help you to create formatting commit message that able to use to generate `CHANGELOG` file and easy to read the same commit format.

## Version

I seperate each individual version into each branch (current 6 Feb 2019 is version 3.0.0)

1. version 1 on [version/1.x.x](https://github.com/kamontat/gitgo/tree/version/1.x.x)
2. version 2 on [version/2.x.x](https://github.com/kamontat/gitgo/tree/version/2.x.x)
3. version 3 on [version/3.x.x](https://github.com/kamontat/gitgo/tree/version/3.x.x)
=======
# gitgo

git commit and branch creator for organize developer

## Installation

1. Download bundle version with latest version at [release](https://github.com/kamontat/gitgo/releases/latest)
2. Choose file that matches your computer OS
3. change file name to `gitgo` and move to bin folder
    1. MacOS bin folder: `/usr/local/bin/` or `/usr/bin`

### Relative libraries

You might need to create changelog. This libraries is made for you

- Git generate changelog: https://github.com/git-chglog/git-chglog

## Format

This project create to make every commit and branch name be same format for all developer and contributor. Which key hsa mark as **optional** mean you can turn it off by setting config files

### Commit

```text
type(scope): title
message
```

1. `Type` is the word (usually contain only 1 word) that represent the commit
2. `Scope` is a scope of commit type
3. `Title` is the important part, that show what the commit for (should less that 50 word)
4. `Message` (optional) is the long description about the commit that might/might not relate to the commit, e.g. add sign text, add long description for more detail

### Branch

```text
iter/key/title/desc/issue
```

1. `Iter` (Optional) is the **iteration** number for agile project, to seperate the branch to each of iteration and make easy to review overall of each iteration
2. `Key` is the main part of the branch, it should be only 1 word to represent the action of this branch (e.g. update, add, refactor). Notes that this should be *verb*
3. `Title` is the title of branch mostly we call 'action'. This is the action or subtype of the key and should stay on 1-2 word only. Notes that this should be *noun*
4. `Desc` (Optional) is the description of branch but it should more and 2-3 word and seperate each word by `-` (dash)
5. `Issue` (Optional) is the requirement of Waffle.io for make automate workflow in github. This represent github issue with or without `#` sign (Optional).

### Help

```sh 

Gitgo: git commit for organize user.

This command create by golang with cobra cli.

Motivation by gitmoji,
I used to like gitmoji but emoji isn't made for none developer.
And the problem I got is I forget which emoji is represent what.
And hard to generate changelog file.
So I think 'short key text' is the solution of situation.

3.1.1 -> Change default and local configuration list
3.1.0 -> Add --tag to changelog generator
3.0.1 -> Add README file to local config
3.0.0 -> Change commit format and refactor code
2.4.0 -> Add --empty to allow empty changes to commit code
2.3.2 -> Issue hash tag will be always add if setting is true
2.3.1 -> Fix branch creator error, and improve logger
2.3.0 -> Add changelog command with initial changelog
2.2.1 -> Improve branch creator and commit creator

Usage:
  gitgo [command]

Available Commands:
  branch      git branch management, get current branch
  changelog   Create changelog
  commit      Git commit with format string
  config      Gitgo configuration
  help        Help about any command

Flags:
  -D, --debug     add debug output
  -h, --help      help for gitgo
  -V, --verbose   add verbose output
      --version   version for gitgo

Use "gitgo [command] --help" for more information about a command.

```
>>>>>>> version/3.x.x
