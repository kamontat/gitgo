package e

import (
	"github.com/kamontat/go-error-manager"
)

// ErrorType for throw
type ErrorType int

// Code is getter error code
func (e ErrorType) Code() int {
	return int(e)
}

const (
	// InitialError is error when first output
	InitialError ErrorType = 10
	// CommitError is error while try to commit
	CommitError ErrorType = 11
	// PreCommitError is error when setup to commit
	PreCommitError ErrorType = 12
	// BranchError is error when branch management
	BranchError ErrorType = 20
	// ChangelogError is error when export changelog
	ChangelogError ErrorType = 30
	// UserError is error cause by user
	UserError ErrorType = 50
)

// Throw throw by message
func Throw(t ErrorType, message string) *manager.Throwable {
	return manager.NewE().AddMessage(message).Throw().SCode(t.Code())
}

// ThrowE throw by error
func ThrowE(t ErrorType, e error) *manager.Throwable {
	return manager.NewE().Add(e).Throw().SCode(t.Code())
}

// Update will update throwable as error type
func Update(throw *manager.Throwable, t ErrorType) *manager.Throwable {
	return throw.SCode(t.Code())
}

// Show is util method for show error message
func Show(t *manager.Throwable) {
	t.ShowMessage()
}

// ShowAndExit is util method for show error message and instant exit program
func ShowAndExit(t *manager.Throwable) {
	t.ShowMessage().Exit()
}
