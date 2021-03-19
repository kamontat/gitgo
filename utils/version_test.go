package utils_test

import (
	"fmt"
	"testing"

	"github.com/kamontat/gitgo/utils"
)

func TestMajorVersion(t *testing.T) {
	current := "v5.4.0"

	successCases := []string{"5", "5.4", "5.4.0"}
	failureCases := []string{"1", "2", "3", "4", "55", "v5", "5v"}

	for _, tcase := range successCases {
		t.Run(fmt.Sprintf("checking %s with %s", tcase, current), func(t *testing.T) {
			err := utils.VersionChecker(tcase, current)
			if err != nil {
				t.Error("expected not error")
			}
		})
	}

	for _, tcase := range failureCases {
		t.Run(fmt.Sprintf("checking %s with %s", tcase, current), func(t *testing.T) {
			err := utils.VersionChecker(tcase, current)
			if err == nil {
				t.Error("expected error occurred")
			}
		})
	}
}
