package utils

import (
	"fmt"
	"regexp"

	"github.com/kamontat/gitgo/utils/phase"
)

// VersionChecker will throw error if config version is not accepted
func VersionChecker(configVersion string, appVersion string) {
	m, err := regexp.MatchString("^v"+configVersion, appVersion)
	phase.Error(err)

	if !m {
		phase.Error(fmt.Errorf("your config is for version '%s' but you run with version %s", configVersion, appVersion))
	}
}
