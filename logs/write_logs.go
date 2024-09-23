package logs

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
)

// 写入一条 begin 日志。
func begin() {
	logs_controller.begin_logger.Println()
}

// 写入一条 end 日志。
func end() {
	logs_controller.end_logger.Println()
}

// 写入一条 debug 日志。
func Debug(debug string) (uintptr, error) {
	pc, filename, line, ok := runtime.Caller(1)
	if !ok {
		return pc, errors.New("logs: Unknown error")
	}
	if (user_logs_style.OutputFlag & log.Llongfile) != 0 {
		debug = fmt.Sprintf("%s:%d %s", filename, line, debug)
	} else if (user_logs_style.OutputFlag & log.Lshortfile) != 0 {
		index := strings.LastIndexByte(filename, '/')
		debug = fmt.Sprintf("%s:%d %s", filename[index+1:], line, debug)
	}
	logs_channel <- aLog{
		log_type: debug_log,
		log:      debug,
	}
	return pc, nil
}

// 写入一条 info 日志。
func Info(info string) (uintptr, error) {
	pc, filename, line, ok := runtime.Caller(1)
	if !ok {
		return pc, errors.New("logs: Unknown error")
	}
	if (user_logs_style.OutputFlag & log.Llongfile) != 0 {
		info = fmt.Sprintf("%s:%d %s", filename, line, info)
	} else if (user_logs_style.OutputFlag & log.Lshortfile) != 0 {
		index := strings.LastIndexByte(filename, '/')
		info = fmt.Sprintf("%s:%d %s", filename[index+1:], line, info)
	}
	logs_channel <- aLog{
		log_type: info_log,
		log:      info,
	}
	return pc, nil
}

// 写入一条 warning 日志。
func Warning(warning string) (uintptr, error) {
	pc, filename, line, ok := runtime.Caller(1)
	if !ok {
		return pc, errors.New("logs: Unknown error")
	}
	if (user_logs_style.OutputFlag & log.Llongfile) != 0 {
		warning = fmt.Sprintf("%s:%d %s", filename, line, warning)
	} else if (user_logs_style.OutputFlag & log.Lshortfile) != 0 {
		index := strings.LastIndexByte(filename, '/')
		warning = fmt.Sprintf("%s:%d %s", filename[index+1:], line, warning)
	}
	logs_channel <- aLog{
		log_type: warning_log,
		log:      warning,
	}
	return pc, nil
}

// 写入一条 error 日志。
func Error(err error) (uintptr, error) {
	pc, filename, line, ok := runtime.Caller(1)
	if !ok {
		return pc, errors.New("logs: Unknown error")
	}
	err_text := err.Error()
	if (user_logs_style.OutputFlag & log.Llongfile) != 0 {
		err_text = fmt.Sprintf("%s:%d %s", filename, line, err_text)
	} else if (user_logs_style.OutputFlag & log.Lshortfile) != 0 {
		index := strings.LastIndexByte(filename, '/')
		err_text = fmt.Sprintf("%s:%d %s", filename[index+1:], line, err_text)
	}
	logs_channel <- aLog{
		log_type: error_log,
		log:      err_text,
	}
	return pc, nil
}
