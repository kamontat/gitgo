# gitgo
git management, format commit message and customization

## Commands 
This CLI will usage syntax as `gitgo [<global option>] [command] [subcommand] [<option>] arguments...`

### Initial
<details>
  <summary>Details</summary>

#### Description

For initial git, same as `git init`

#### Example usage

- `gitgo init`
- `gitgo i`

| Options       | Description         |
|---------------|---------------------|
| --force \| -f | Force reinitial git |

</details>

### Destory
<details>
  <summary>Details</summary>

#### Description

Delete git repo, remove **.git** folder

#### Example usage

- `gitgo destroy`
- `gitgo d`

| Options       | Description                 |
|---------------|-----------------------------|
| --force \| -f | Force delete without prompt |

</details>

### Status
<details>
  <summary>Details</summary>

#### Description

Show git status, same as git command `git status`

#### Example usage

- `gitgo status`
- `gitgo s`

| Options       | Description                 |
|---------------|-----------------------------|
| --force \| -f | Force delete without prompt |

</details>


### Configuration
<details>
  <summary>Details</summary>

#### Description

Manage cli configuration

#### Example usage

- `gitgo configuration`
- `gitgo config`
- `gitgo g`

#### Actions

- `gitgo config` - open configuration file by default text-editor (use environment variable call `$EDITOR`)
- `gitgo config location|l` - show current location of configuration file
- `gitgo config --key <key>` - get value in config file by key
- `gitgo config --key <key> --value <value>` - set value into config file


| Options               | Description                                         |
|-----------------------|-----------------------------------------------------|
| --key \| -k <key>     | Input config key, separate layer by `.`             |
| --value \| -v <value> | Input config value, use only want to save new value |

#### Example Configuration

By default configuration folder will be on 
- Default `~/.config/github.com/kamontat/gitgo/config/`
- Also able to get from `$GOPATH/src/github.com/kamontat/gitgo/config/` 

```yaml
config:
    commit:
        type: text # text | emoji
        emoji: string # string | emoji
        key:
            require: true # true | false 
        title:
            require: true # true | false 
            auto: false # true | false
            size: 50 # maximum charecter size
        message:
            require: true # true | false 
        showsize: 8 # list when show the 'commit list'
    editor: '' # vim | nano | other cli..
```

#### Extra

This syntax of key is separate by dot, e.g. `config.editor`, `config.commit.key` etc.

</details>

### Add
<details>
  <summary>Details</summary>

#### Description

Add file/folder to git, similar with `git add <args>`

#### Example usage

- `gitgo add`
- `gitgo a`

#### Actions

- `gitgo add <args>` - add <args> (files or folder) into git
- `gitgo add all|a` - add **every files and folders** into git
- `gitgo add --all|-A` - same as `gitgo add all`

| Options       | Description                        |
|---------------|------------------------------------|
| --all \| -A   | Add every files and folders to git |

</details>

### Commit

<details>
  <summary>Details</summary>

#### Description

Git commit with default format and custom syntax. Next plan this will able to custom git message format.

#### Example usage

- `gitgo commit`
- `gitgo c`

#### Actions

- `gitgo commit` - commit with [config](#configuration) `type` and this will prompt information to user, for generating commit message
- `gitgo commit emoji|moji|e` - commit message with `emoji type`
- `gitgo commit emoji|moji|e initial|init|i` - initial commit message with `emoji type`
- `gitgo commit text|t` - commit message with `text type`
- `gitgo commit text|t initial|init|i` - initial commit message with `text type`
- `gitgo commit initial|init|i` - generate initial commit with default message and [config](#configuration) `type`

| Options               | Description                                          |
|-----------------------|------------------------------------------------------|
| --add \| -a           | Include add option into commit (git -am "<message>") |
| --all \| -A           | Run `git add --all` command, before commit code      |
| --key \| -k <key>     | Add commit [key](#commit-key) to commit message      |
| --title \| -t <title> | Add commit [title](#commit-title) to commit message  |

#### Commit type

On commitment, I create 2 type of them
1. Emoji type, emoji type will split commit purpose by emoji. Those emoji you able to custom by yourself without modify anything in the code (by [config](#configuration))
2. Text type, this type will use text to split purpose commit by text and also customizable (by [config](#configuration))

#### Commit message

Default commit message will follow this format. <br>
For **text** type
```
[key]: title
message
```
For **emoji** type
```
key: title
message
```

The concept of this format is easy to `read` and reproduce to `changelog`. This split messages to 3 sessions **key**, **title** and **message**

#### Commit key

Commit key should be **short**, **easy to understand**, **singular**, and **1 word**. 
This parameter will use for easy to *reverse* or *check* the result of a commit and *understand* what is commit duty.

For example: `test`, `improve`, `fix`, `feature`

#### Commit title

Commit title should **short**, **clear** and **not** longer than 50 words.
This parameter will use for *create changelog*, and *tell more* information about commit.

#### Commit message

This parameter will use for deeply information about the commit, This should tell everything of the commit, 
in case later developer need to reverse to this commit, all known bug, error, information, etc. 

Basically this commit will use only release version (include *alpha*, *beta*) or tag version

</details>

### Push
<details>
  <summary>Details</summary>

#### Description

Push local git repository update to server (Github, Bitbucket, etc.)

#### Example usage

- `gitgo push`
- `gitgo p`

#### Actions

- `gitgo push [<branch>...]` - push local code to input branch or `master` (default)
- `gitgo push --repo <repo> [<branch>...]` - push local code as input branch or `master` (default) to input remote repository or `origin` (default)
- `gitgo push set|s <link>` - initial/set push server and remote, this command will *create remote*, *set upstream* to current branch, and *push code* changes

| Options                             | Description                                                  |
|-------------------------------------|--------------------------------------------------------------|
| --force \| -f                       | Force to push local code to server code                      |
| --repository \| --repo \| -r <repo> | Change repository remote, default is `origin`                |
| --branch \| -b <branch>             | Change server branch, default is `master` (for **SET** only) |

</details>

### Pull
<details>
  <summary>Details</summary>

#### Description

Pull a repository/code from a server git to local git

#### Example usage

- `gitgo pull`
- `gitgo P` (capital P)

#### Actions

- `gitgo pull [<branch>...]` - pull code from server by input branch (default is `master`) to current branch

| Options                             | Description                                                  |
|-------------------------------------|--------------------------------------------------------------|
| --force \| -f                       | Force to push local code to server code                      |
| --repository \| --repo \| -r <repo> | Change repository remote, default is `origin`                |

</details>

## Creator

Kamontat Chantrachirathumrong

## LICENSE

[MIT](https://opensource.org/licenses/MIT)

<details>
  <summary>Details</summary>

Copyright 2018 Kamontat Chantrachirathumrong

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

</details>
