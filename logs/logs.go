package logs

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"
)

// LogsStyle 用于决定 logs 服务中的日志风格。LogsChannelSize 决定日志写入管道的大小，BEFlag 决定输出 begin 和 end 日志的风格，PreBegin 和 PreEnd 为 begin 和 end 日志的前缀，OutputFlag 决定输出 debug、info、warning 和 error 等日志的风格，PreDebug、PreInfo、PreWarning 和 PreError 分别为 debug、info、warning 和 error 日志的前缀。BEFlag 和 OutputFlag 可以使用 log 包中的 flag 赋值。
type LogsStyle struct {
	LogsChannelSize uint

	BEFlag int

	PreBegin string
	PreEnd   string

	OutputFlag int

	PreDebug   string
	PreInfo    string
	PreWarning string
	PreError   string
}

// 生成一个默认风格的日志风格对象。
func NewLogsStyle() *LogsStyle {
	return &LogsStyle{
		LogsChannelSize: 1024,

		BEFlag: log.Ldate | log.Ltime,

		PreBegin: ">> Log begin: ",
		PreEnd:   "<< Log end: ",

		OutputFlag: log.Ldate | log.Ltime,

		PreDebug:   "[DEBUG] ",
		PreInfo:    "[INFO] ",
		PreWarning: "[WARNING] ",
		PreError:   "[ERROR] ",
	}
}

type aLog struct {
	log_type int
	log      string
}

type logsController struct {
	log_file *os.File

	log_file_writer io.Writer

	begin_logger *log.Logger
	end_logger   *log.Logger

	debug_logger   *log.Logger
	info_logger    *log.Logger
	warning_logger *log.Logger
	error_logger   *log.Logger
}

var (
	default_logs_style = NewLogsStyle()
)

const (
	begin_log = iota
	end_log
	debug_log
	info_log
	warning_log
	error_log
)

var (
	logs_ready      bool
	logs_run        sync.RWMutex
	user_logs_style *LogsStyle
	logs_channel    chan aLog
	logs_controller logsController
)

func cancelConflictFlag(flag int) int {
	return flag & (log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lmsgprefix | log.LstdFlags)
}

// 使用一个已打开的可写文件启动指定风格的 logs 服务。使用时一般会与 ReleaseLogs 方法连用。
func UseLogs(used_log_file *os.File, logs_style *LogsStyle) error {
	if used_log_file == nil {
		return errors.New("logs: Used log file does not exist")
	}
	logs_controller.log_file = used_log_file

	logs_controller.log_file_writer = logs_controller.log_file

	logs_controller.begin_logger = log.New(logs_controller.log_file_writer, logs_style.PreBegin, cancelConflictFlag(logs_style.BEFlag))
	logs_controller.end_logger = log.New(logs_controller.log_file_writer, logs_style.PreEnd, cancelConflictFlag(logs_style.BEFlag))

	logs_controller.debug_logger = log.New(logs_controller.log_file_writer, logs_style.PreDebug, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.info_logger = log.New(logs_controller.log_file_writer, logs_style.PreInfo, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.warning_logger = log.New(logs_controller.log_file_writer, logs_style.PreWarning, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.error_logger = log.New(logs_controller.log_file_writer, logs_style.PreError, cancelConflictFlag(logs_style.OutputFlag))

	user_logs_style = logs_style

	logs_channel = make(chan aLog, logs_style.LogsChannelSize)
	if pan := recover(); pan != nil {
		return errors.New("manager: Failed to create the memory allocation required for the channel of logs")
	}

	logs_ready = true
	begin()

	return nil
}

// 使用一个已打开的可写文件启动默认风格的 logs 服务。使用时一般会与 ReleaseLogs 方法连用。
func UseDefault(used_log_file *os.File) error {
	err := UseLogs(used_log_file, default_logs_style)
	if err != nil {
		return err
	}
	return nil
}

// 释放通过已打开的可写文件启动的 logs 服务的资源。
func ReleaseLogs() {
	logs_ready = false
	end()
	logs_controller.log_file = nil
}

// 重定向通过 UseLogs 或 UseDefault 打开的 logs 服务的输出日志文件。
func ReuseOutput(output_file *os.File) error {
	if output_file == nil {
		return errors.New("logs: Output file does not exist")
	}

	end()

	logs_controller.log_file = output_file

	logs_controller.log_file_writer = logs_controller.log_file

	logs_controller.debug_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.info_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.warning_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.error_logger.SetOutput(logs_controller.log_file_writer)

	logs_ready = true
	begin()

	return nil
}

// 通过日志文件路径启动指定风格的 logs 服务。使用时一般会与 CloseLogs 方法连用。
func OpenLogs(log_file_path string, logs_style *LogsStyle) error {
	var err error
	logs_controller.log_file, err = os.OpenFile(log_file_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	logs_controller.log_file_writer = logs_controller.log_file

	logs_controller.begin_logger = log.New(logs_controller.log_file_writer, logs_style.PreBegin, cancelConflictFlag(logs_style.BEFlag))
	logs_controller.end_logger = log.New(logs_controller.log_file_writer, logs_style.PreEnd, cancelConflictFlag(logs_style.BEFlag))

	logs_controller.debug_logger = log.New(logs_controller.log_file_writer, logs_style.PreDebug, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.info_logger = log.New(logs_controller.log_file_writer, logs_style.PreInfo, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.warning_logger = log.New(logs_controller.log_file_writer, logs_style.PreWarning, cancelConflictFlag(logs_style.OutputFlag))
	logs_controller.error_logger = log.New(logs_controller.log_file_writer, logs_style.PreError, cancelConflictFlag(logs_style.OutputFlag))

	user_logs_style = logs_style

	logs_channel = make(chan aLog, logs_style.LogsChannelSize)
	if pan := recover(); pan != nil {
		return errors.New("logs: Failed to create the memory allocation required for the channel of logs")
	}

	logs_ready = true
	begin()

	return nil
}

// 通过日志文件路径启动默认风格的 logs 服务。使用时一般会与 CloseLogs 方法连用。
func OpenDefault(log_file_path string) error {
	err := OpenLogs(log_file_path, default_logs_style)
	if err != nil {
		return err
	}
	return nil
}

// 释放通过日志文件路径启动的 logs 服务的资源。
func CloseLogs() error {
	logs_ready = false
	end()
	err := logs_controller.log_file.Close()
	if err != nil {
		return err
	}
	return nil
}

// 重定向通过 OpenLogs 或 OpenDefault 打开的 logs 服务的输出日志文件。
func ReopenOutput(output_path string) error {
	err := CloseLogs()
	if err != nil {
		return err
	}

	logs_controller.log_file, err = os.OpenFile(output_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	logs_controller.log_file_writer = logs_controller.log_file

	logs_controller.debug_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.info_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.warning_logger.SetOutput(logs_controller.log_file_writer)
	logs_controller.error_logger.SetOutput(logs_controller.log_file_writer)

	logs_ready = true
	begin()

	return nil
}

// 获取当前启动的 logs 服务相关信息。
func GetLogsInfo() (string, *LogsStyle) {
	return logs_controller.log_file.Name(), user_logs_style
}

// 运行已启动的 logs 服务。使用时一般会与 Over 方法连用。logs 服务可以通过 UseLogs、UseDefault、OpenLogs 或 OpenDefault 等方法启动。
func Run() error {
	if !logs_ready {
		if logs_controller.log_file == nil {
			return errors.New("logs: Log file of logs controller does not exist")
		}

		if logs_channel == nil {
			return errors.New("logs: Logs channel does not exist")
		}

		return errors.New("logs: Logs does not ready")
	}

	logs_run.Lock()
	go func() {
		for {
			if a_log, exist := <-logs_channel; exist {
				switch a_log.log_type {
				case debug_log:
					logs_controller.debug_logger.Println(a_log.log)
				case info_log:
					logs_controller.info_logger.Println(a_log.log)
				case warning_log:
					logs_controller.warning_logger.Println(a_log.log)
				case error_log:
					logs_controller.error_logger.Println(a_log.log)
				}
			} else {
				break
			}
		}
		logs_run.Unlock()
	}()

	return nil
}

// 停止正在运行的 logs 服务。
func Over() {
	close(logs_channel)
	logs_run.RLock()
	logs_run.RUnlock()
}
