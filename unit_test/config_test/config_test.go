package unit_test

import (
	"syscall"
	"testing"

	"github.com/ZSLTChenXiYin/MyGO/config"
	"github.com/ZSLTChenXiYin/MyGO/logs"
	"github.com/ZSLTChenXiYin/MyGO/manager"
)

func TestImportLogsConfig(t *testing.T) {
	err := config.ImportLogsConfig("logs_config.json")
	if err != nil {
		t.Fatal(err)
	}

	err = config.ImportManagerConfig("manager_config.json")
	if err != nil {
		t.Fatal(err)
	}
	err = manager.Events(func() bool {
		logs.Debug("The process has stopped running")
		logs.Over()
		logs.CloseLogs()
		return true
	}, syscall.SIGINT, syscall.SIGTERM)
	if err != nil {
		t.Fatal(err)
	}

	_, err = logs.Debug("The log service has been enabled")
	if err != nil {
		t.Error(err)
	}

	manager.Run()
}
