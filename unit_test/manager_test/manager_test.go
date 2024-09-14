package unit_test

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/ZSLTChenXiYin/MyGO/logs"
	"github.com/ZSLTChenXiYin/MyGO/manager"
)

func Serve() {
	for {
		logs.Info("Serving users")
		time.Sleep(time.Second * 10)
	}
}

func TestManager(t *testing.T) {
	logs.UseDefault(os.Stdout)

	manager.CreateManager(1, 1)
	manager.Events(func() bool {
		logs.Debug("The process has stopped running")
		logs.CloseLogs()
		return true
	}, syscall.SIGINT, syscall.SIGTERM)

	go Serve()

	manager.Run()
}
