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
	err := logs.UseDefault(os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
	err = logs.Run()
	if err != nil {
		t.Fatal(err)
	}

	err = manager.CreateManager(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = manager.Events(func() bool {
		logs.Debug("The process has stopped running")

		logs.Over()
		logs.ReleaseLogs()

		return true
	}, syscall.SIGINT, syscall.SIGTERM)
	if err != nil {
		t.Fatal(err)
	}

	go Serve()

	manager.Run()
}
