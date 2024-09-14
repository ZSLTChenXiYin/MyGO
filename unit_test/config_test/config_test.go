package unit_test

import (
	"syscall"
	"testing"

	"github.com/ZSLTChenXiYin/MyGO/config"
	"github.com/ZSLTChenXiYin/MyGO/logs"
	"github.com/ZSLTChenXiYin/MyGO/manager"
)

func TestImportLogsConfig(t *testing.T) {
	config.ImportLogsConfig("logs_config.json")

	config.ImportManagerConfig("manager_config.json")
	manager.Events(func() bool {
		logs.Debug("The process has stopped running")
		logs.CloseLogs()
		return true
	}, syscall.SIGINT, syscall.SIGTERM)

	logs.Debug("The log service has been enabled")

	manager.Run()
}
