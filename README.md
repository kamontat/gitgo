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
[key] title
message
```

1. `Key` is the word (usually contain only 1 word) that represent the commit
2. `Title` is the important part, that show what the commit for (should less that 50 word)
3. `Message` (optional) is the long description about the commit that might/might not relate to the commit, e.g. add sign text, add long description for more detail

### Branch

```text
iter/key/title/desc/issue
```

1. `Iter` (Optional) is the **iteration** number for agile project, to seperate the branch to each of iteration and make easy to review overall of each iteration
2. `Key` is the main part of the branch, it should be only 1 word to represent the action of this branch (e.g. update, add, refactor). Notes that this should be *verb*
3. `Title` is the title of branch mostly we call 'action'. This is the action or subtype of the key and should stay on 1-2 word only. Notes that this should be *noun*
4. `Desc` (Optional) is the description of branch but it should more and 2-3 word and seperate each word by `-` (dash)
5. `Issue` (Optional) is the requirement of Waffle.io for make automate workflow in github. This represent github issue with or without `#` sign (Optional).
