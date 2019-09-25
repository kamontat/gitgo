package util

import "time"

// YAML is object of config yaml
type YAML struct{}

// GeneratorYAML will return YAML Object
func GeneratorYAML() *YAML {
	return &YAML{}
}

// ReadmeMarkdown will return readme file template
func (y *YAML) ReadmeMarkdown(version string) string {
	t := time.Now()
	return `# Gitgo (v` + version + `)

This is a configuration file for gitgo repository with hosting on https://github.com/kamontat/gitgo/tree/version/4.x.x

### Creator

- Kamontat Chantrachirathumrong (https://github.com/kamontat)

### Datetime

Someone create this configuration on '` + t.UTC().Format(time.UnixDate) + `'

### Thank you

Thank you for using this command to manage your project :)
`
}

// Config is a default config.yaml
func (y *YAML) Config() string {
	return `# All **type** on each configuration can be 3 value only
#   1. list      -> in order to use this key the list.yaml must be presented and list of root name must be presented
#                   the prompt will show as choice input
#   2. input     -> the prompt will show as string input
#   3. mulitline -> the prompt will show as multiple line input
#   4. mix       -> this will mix between list and custom

# enable and require is configuration root level
# if enable is true, mean command will prompt user to input something
# and if require is true, you will cannot add empty string to that prompt

version: 4
settings:
  log: info
commit:
  key:
    enable: true
    require: true
    type: list
    page: 5
  scope:
    enable: true
    require: false
    type: list
    page: 5
  title:
    enable: true
    require: true
    type: input
    size: 75
  message:
    enable: false
    require: false
    type: multiline
    size: 200
branch:
  iteration:
    enable: true
    require: true
  key:
    type: list
    size: 15
    page: 5
  title:
    type: mix
    size: 20
    page: 5
  description:
    enable: true
    require: false
`
}

// ListConfig is default list.yaml
func (y *YAML) ListConfig() string {
	return `version: 4
commit:
  keys:
    - type: feat
      value: Introducing new features.
    - type: impr
      value: Improving user experience / usability / reliablity.
    - type: fix
      value: Fixing a bug.
    - type: refac
      value: A code change that neither fixes a bug nor adds a feature.
    - type: chore
      value: Other changes that don't modify src or test files.
    - type: perf
      value: A code change that improves performance.
  scopes:
    - type: model
      value: Model folders
    - type: api
      value: APIs folders
    - type: controller
      value: Controller folders
  titles: 
    - type: Start new project
      value: Start new project
branch:
  keys:
    - type: feat
      value: Introducing new features or project enhancement.
    - type: impr
      value: Improving user experience / usability / performance.
    - type: fix
      value: Fixing a bug.
  titles:
    - type: config
      value: Configuration
`
}

// ChgLogConfig will return chglog configuration template
func (y *YAML) ChgLogConfig(style, repoURL string) string {
	return `style: ` + style + `
template: CHANGELOG.tpl.md
info:
  title: CHANGELOG
  repository_url: ` + repoURL + `
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

// ChgLogTpl will return chglog release-notes templates
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
