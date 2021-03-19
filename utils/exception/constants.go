package exception

import "github.com/kamontat/gitgo/utils/logger"

// InitialPhaseCode when initial error
const InitialPhaseCode = 10

// OnInitialPhase will check and throw error if initial error
func OnInitialPhase(err error) {
	When(err).LogExit(logger.Get(), InitialPhaseCode)
}
