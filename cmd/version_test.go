package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/snebel29/kubewatch/pkg/logging"
	"testing"
)

func TestVersionLdFlags(t *testing.T) {
	if buildDate == "" || gitCommit == "" {
		t.Errorf("Some flag is not set => buildDate: %s, gitCommit: %s", buildDate, gitCommit)
	}
}

func TestVersionPrettyString(t *testing.T) {
	logger, hook := test.NewNullLogger()
	log := &logging.Logger{logrus.NewEntry(logger)}
	versionPrettyString(log)
	num := len(hook.Entries)
	if num != 2 {
		t.Errorf("Wrong number [%d] of output messages", num)
	}
}
