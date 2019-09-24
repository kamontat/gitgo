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
	// IsInitial is error when first output
	IsInitial ErrorType = 10
	// IsCommit is error while try to commit
	IsCommit ErrorType = 20
	// IsPreCommit is error when setup to commit
	IsPreCommit ErrorType = 19
	// IsPostCommit is error when cleanup after commit completed
	IsPostCommit ErrorType = 21
	// IsBranch is error when branch management
	IsBranch ErrorType = 20
	// IsCheckout is error occurred when checkout code
	IsCheckout ErrorType = 30
	// IsChangelog is error when export changelog
	IsChangelog ErrorType = 40
	// IsUser is error cause by user
	IsUser ErrorType = 50
	// IsLibrary is error cause by external libraries
	IsLibrary ErrorType = 100
)

// ErrorMessage throw by message
func ErrorMessage(t ErrorType, message string) *manager.Throwable {
	return manager.NewE().AddMessage(message).Throw().SCode(t.Code())
}

// Error throw by error
func Error(t ErrorType, e error) *manager.Throwable {
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
