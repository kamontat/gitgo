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

1. `Type` (configable) is the word (usually contain only 1 word) that represent the commit
2. `Scope` (configable) is a scope of commit type
3. `Title` (configable) is the important part, that show what the commit for (should less that 50 word)
4. `Message` (configable) is the long description about the commit that might/might not relate to the commit, e.g. add sign text, add long description for more detail

### Branch

```text
iter/key/title/description
```

1. `iter` (configable) is the **iteration** number for agile project, to seperate the branch to each of iteration and make easy to review overall of each iteration
2. `key` (configable) is the key part of the branch, it should be only 1 word to represent the action of this branch (e.g. update, add, refactor). Notes that this should be *verb*
3. `title` (configable) is the title of branch mostly we call 'action'. This is the action or subtype of the key and should stay on 1-2 word only. Notes that this should be *sentence* and seperate each word by `-` (dash)
4. `description` (configable) is the description of branch but it should more and 2-3 word and seperate each word by `-` (dash)

### Help

```sh

Gitgo: git commit for organize user.

This command create by golang with cobra cli.

Motivation by gitmoji and GitFlow,
Force everyone to create the exect same templates of commit and branch

4.0.0-beta.4 - update README on config folder to be version 4 instead of 3
4.0.0-beta.3 - fix commit always call git commit even prompt was exited
4.0.0-beta.2 - refactor how configuration will be loaded
               and add more key to configable
4.0.0-beta.1 - remove global configuration settings;
               force every settings should place in project

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

## Development

### How to

1. Install all dependencies via `go get ./...`
2. Install Gitgo to GOBIN via `go install`
