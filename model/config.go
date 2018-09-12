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

func (y *YAML) ChgLogConfig(style, repoUrl string) string {
	return `style: ` + style + `
template: CHANGELOG.tpl.md
info:
  title: CHANGELOG
  repository_url: ` + repoUrl + `
options:
  commits:
    filters:
      Type:
        - feature
        - fix
        - improve
        - refactor
        - doc
  commit_groups:
    title_maps:
      feature: Features
      improve: Improving User Experience
      refactor: Code Refactoring
      fix: Fixes Bug
      doc: Documentation
  header:
    pattern: "^\\[(\\w*)\\]\\s(.*)$"
    pattern_maps:
      - Type
      - Subject
  issues: 
    prefix: 
      - "#"
  notes:
    keywords:
      - BREAKING CHANGE
`
}

func (y *YAML) ChgLogTpl() string {
	return `{{ if .Versions -}}
<a name="unreleased"></a>
## [Unreleased]

{{ if .Unreleased.CommitGroups -}}
{{ range .Unreleased.CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ if .Tag.Previous }}[{{ .Tag.Name }}]{{ else }}{{ .Tag.Name }}{{ end }} - {{ datetime "2006-01-02" .Tag.Date }}
{{ range .CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}

{{- if .MergeCommits -}}
### Pull Requests
{{ range .MergeCommits -}}
- {{ .Header }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### {{ .Title }}
{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{- if .Versions }}
[Unreleased]: {{ .Info.RepositoryURL }}/compare/{{ $latest := index .Versions 0 }}{{ $latest.Tag.Name }}...HEAD
{{ range .Versions -}}
{{ if .Tag.Previous -}}
[{{ .Tag.Name }}]: {{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}
{{ end -}}
{{ end -}}
{{ end -}}`
}
