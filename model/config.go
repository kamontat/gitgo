package model

type YAML struct{}

func GeneratorYAML() *YAML {
	return &YAML{}
}

func (y *YAML) GDefaultConfig() string {
	return `version: 2
log: true
commit:
  message: false
`
}

func (y *YAML) GDefaultList() string {
	return `version: 2
list:
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
`
}

func (y *YAML) LEmptyList() string {
	return `version: 2
list:
  - key: empty
    value: Update this commit header
`
}
