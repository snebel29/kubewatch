package cmd

import (
	"github.com/sirupsen/logrus/hooks/test"
	"testing"
)

func TestVersionLdFlags(t *testing.T) {
	if buildDate == "" || gitCommit == "" {
		t.Errorf("Some flag is not set => buildDate: %s, gitCommit: %s", buildDate, gitCommit)
	}
}

func TestVersionPrettyString(t *testing.T) {
	logger, hook := test.NewNullLogger()
	versionPrettyString(logger)
	num := len(hook.Entries)
	if num != 2 {
		t.Errorf("Wrong number [%d] of output messages", num)
	}
}
