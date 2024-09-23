package unit_test

import (
	"fmt"
	"testing"

	"github.com/ZSLTChenXiYin/MyGO/logs"
)

const (
	log_file_path = "log"
)

func TestLogs(t *testing.T) {
	err := logs.OpenDefault(log_file_path)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = logs.CloseLogs()
		if err != nil {
			t.Fatal(err)
		}
	}()
	logs.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer logs.Over()

	_, err = logs.Debug("The log service has been enabled")
	if err != nil {
		t.Error(err)
	}
	_, err = logs.Debug(fmt.Sprintf("The %s debug", "second"))
	if err != nil {
		t.Error(err)
	}
}

func TestLogsReader(t *testing.T) {
	logs_reader := logs.NewLogsReader(log_file_path, logs.NewLogsStyle())
	if logs_reader == nil {
		t.Fatal("logs_test: No object obtained")
	}
	defer func() {
		err := logs_reader.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	all_begin, err := logs_reader.FindAllBegin()
	if err != nil {
		t.Fatal(err)
	}

	all_end, err := logs_reader.FindAllEnd()
	if err != nil {
		t.Fatal(err)
	}

	all_log, err := logs_reader.GetDebug(all_begin[0].Line, all_end[0].Line)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Line\tLog")
	for _, a_log := range all_log {
		fmt.Printf("%d\t%s", a_log.Line, a_log.Log)
	}
}
