package model

import "time"

// YAML is object of config yaml
type YAML struct{}

// GeneratorYAML will return YAML Object
func GeneratorYAML() *YAML {
	return &YAML{}
}

func (y *YAML) ReadmeMarkdown(version string) string {
	t := time.Now()
	return `# Gitgo (v` + version + `)

  This is a configuration file for gitgo repository with hosting on https://github.com/kamontat/gitgo/tree/version/3.x.x

### Creator

- Kamontat Chantrachirathumrong (https://github.com/kamontat)

### Datetime

Someone create this configuration on '` + t.UTC().Format(time.UnixDate) + `'

### Thank you
Thank you for using this command to manage your project :)
`
}

// GDefaultConfig is global default config.yaml
func (y *YAML) GDefaultConfig() string {
	return `version: 3
log: true
commit:
  message: true
branch:
  iteration:
    require: true
  description:
    require: true
  issue:
    require: false
    hashtag: false
`
}

// GDefaultList is global default list.yaml
func (y *YAML) GDefaultList() string {
	return `version: 3
commits:
  - type: feat
    value: Introducing new features.
  - type: impr
    value: Improving user experience / usability / reliablity.
  - type: fix
    value: Fixing a bug.
  - type: refactor
    value: A code change that neither fixes a bug nor adds a feature.
  - type: chore
    value: Other changes that don't modify src or test files.
branches:
  - type: feat
    value: Introducing new features or project enhancement.
  - type: impr
    value: Improving user experience / usability / performance.
  - type: fix
    value: Fixing a bug.
`
}

// LEmptyList is empty list.yaml
func (y *YAML) LEmptyList() string {
	return `version: 3
commits:
  - type: perf
    value: A code change that improves performance.
  - type: doc
    value: Documenting source code / user manual.
  - type: test
    value: Adding missing tests or correcting existing tests.
  - type: build
    value: Changes that affect the build system or external dependencies.
  - type: custom
    value: this is a custom commit header
branches:
  - type: test
    value: Adding missing tests or correcting existing tests.
  - type: custom
    value: this is a custom branch header
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
        - feat
        - impr
        - perf
        - fix
        - doc
  commit_groups:
    title_maps:
      feat: Feature
      impr: Improving application
      perf: Improving performance
      fix: Fixes Bug
      doc: Documentation
  header:
    pattern: "^(\\w*)(?:\\(([\\w\\$\\.\\-\\*\\s]*)\\))?\\:\\s(.*)$"
    pattern_maps:
      - Type
      - Scope
      - Subject
  issues: 
    prefix: 
      - "#"
  notes:
    keywords:
      - BREAKING CHANGE`
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
