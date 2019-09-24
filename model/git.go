package model

import (
	"io"
	"os"
	"os/exec"

	"github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"
)

// GitCommand is git-cli command with custom stdout and custom stderr
type GitCommand struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

var gc = &GitCommand{}

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

// SetReader will set reader
func (g *GitCommand) SetReader(in io.Reader) *GitCommand {
	g.in = in
	return g
}

// Exec will run git cli in command line
func (g *GitCommand) Exec(args ...string) *manager.ErrManager {
	cmd := exec.Command("git", args...)

	if g.out != nil {
		om.Log.ToVerbose("setting", "custom command output")
		cmd.Stdout = g.out
	} else {
		cmd.Stdout = os.Stdout
	}

	if g.err != nil {
		om.Log.ToVerbose("setting", "custom command error")
		cmd.Stderr = g.err
	} else {
		cmd.Stderr = os.Stderr
	}

	if g.in != nil {
		om.Log.ToVerbose("setting", "custom command input")
		cmd.Stdin = g.in
	} else {
		cmd.Stdin = os.Stdin
	}

	e := cmd.Run()
	return manager.NewE().Add(e)
}
