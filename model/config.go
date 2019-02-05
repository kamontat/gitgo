package model

// YAML is object of config yaml
type YAML struct{}

// GeneratorYAML will return YAML Object
func GeneratorYAML() *YAML {
	return &YAML{}
}

// GDefaultConfig is global default config.yaml
func (y *YAML) GDefaultConfig() string {
	return `version: 3
log: true
commit:
  message: true
`
}

// GDefaultList is global default list.yaml
func (y *YAML) GDefaultList() string {
	return `version: 3
list:
  - type: feature
    value: Introducing new features.
  - type: improve
    value: Improving user experience / usability / performance.
  - type: fix
    value: Fixing a bug.
  - type: refactor
    value: Refactoring code.
  - type: file
    value: Updating file(s) or folder(s).
  - type: doc
    value: Documenting source code / user manual.
  - type: init
    value: Start project or Initial commit.
  - type: release
    value: Release stable version or tags.
`
}

// LEmptyList is empty list.yaml
func (y *YAML) LEmptyList() string {
	return `version: 3
list:
  - type: empty
    value: Update this commit header
`
}
