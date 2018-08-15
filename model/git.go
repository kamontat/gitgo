package model

import (
	"io"
	"os"
	"os/exec"

	"github.com/kamontat/go-error-manager"
)

// GitCommand is git-cli command with custom stdout and custom stderr
type GitCommand struct {
	out io.Writer
	err io.Writer
}

var gc = &GitCommand{
	out: os.Stdout,
	err: os.Stderr,
}

// Git will return singleton GitCommand
func Git() *GitCommand {
	return gc
}

// SetErrWriter will set err writer
func (g *GitCommand) SetErrWriter(err io.Writer) *GitCommand {
	g.err = err
	return g
}

// SetOutWriter will set out writer
func (g *GitCommand) SetOutWriter(out io.Writer) *GitCommand {
	g.out = out
	return g
}

// Exec will run git cli in command line
func (g *GitCommand) Exec(args ...string) *manager.ErrManager {
	cmd := exec.Command("git", args...)
	if g.out != nil {
		cmd.Stdout = g.out
	}
	if g.err != nil {
		cmd.Stderr = g.err
	}

	e := cmd.Start()
	return manager.StartNewManageError().AddNewError(e)
}
