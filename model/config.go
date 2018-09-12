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
branch:
  iteration:
    require: true
  description:
    require: true
  issue:
    require: true
    hashtag: false
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
    value: A code change that neither fixes a bug nor adds a feature.
  - key: file
    value: Add or remove file(s) or folder(s).
  - key: doc
    value: Documenting source code / user manual.
  - key: test
    value: Adding missing tests or correcting existing tests.
  - key: release
    value: Release stable version or tags.
  - key: BREAKING CHANGE
    value: introduce break change code.
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
