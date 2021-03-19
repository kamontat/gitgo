package utils

import (
	"fmt"
	"regexp"
)

// VersionChecker will throw error if config version is not accepted
func VersionChecker(configVersion string, appVersion string) error {
	m, err := regexp.MatchString("^v"+configVersion, appVersion)
	if err != nil {
		return err
	}

	if !m {
		return fmt.Errorf("your config is for version '%s' but you run with version %s", configVersion, appVersion)
	}

	return nil
}
