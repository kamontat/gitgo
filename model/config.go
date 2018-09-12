package model

// YAML is object of config yaml
type YAML struct{}

// GeneratorYAML will return YAML Object
func GeneratorYAML() *YAML {
	return &YAML{}
}

// GDefaultConfig is global default config.yaml
func (y *YAML) GDefaultConfig() string {
	return `version: 2
log: true
commit:
  message: false
  format: "%<key>: %<title> \n %<message>"
branch:
  issue:
    require: true
    hashtag: false
  iteration: true
  format: "%<issue>/%<action>/%<desc>/%<optional>"
`
}

// GDefaultList is global default list.yaml
func (y *YAML) GDefaultList() string {
	return `version: 2
commits:
  - key: feature
    value: Introducing new features.
  - key: improve
    value: Improving user experience / usability / performance.
  - key: fix
    value: Fixing a bug.
  - key: refactor
    value: Refactoring code.
  - key: file
    value: Updating file(s) or folder(s).
  - key: doc
    value: Documenting source code / user manual.
  - key: init
    value: Start project or Initial commit.
  - key: release
    value: Release stable version or tags.
branches:
  - key: enhance
    value: Introducing new features or project enhancement.
  - key: improve
    value: Improving user experience / usability / performance.
  - key: fix
    value: Fixing a bug.
`
}

// LEmptyList is empty list.yaml
func (y *YAML) LEmptyList() string {
	return `version: 2
commits:
  - key: empty
    value: Update this commit header
branches:
  - key: empty
    value: Update this branch header
`
}
