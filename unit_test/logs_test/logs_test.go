package unit_test

import (
	"testing"

	"github.com/ZSLTChenXiYin/MyGO/logs"
)

func TestLogs(t *testing.T) {
	logs.Default("log")
	defer logs.CloseLogs()

	logs.Debug("The log service has been enabled")
	logs.Debugf("The %s debug\n", "second")
}
