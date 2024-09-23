package logs

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

const (
	logs_read_size = 1024
)

// 存放读取的日志的结构。Line 为该条日志存在日志文件的行号，Log 为日志内容。
type ALog struct {
	Line int64
	Log  string
}

// LogsReader 实现了读取日志的一系列方法。
type LogsReader struct {
	log_file        *os.File
	log_file_reader *bufio.Reader
	logs_style      *LogsStyle

	line int64
}

// 生成一个 LogsReader 实例。停止使用前注意调用 Close 方法释放占用的资源。
func NewLogsReader(log_file_path string, logs_style *LogsStyle) *LogsReader {
	log_file, err := os.OpenFile(log_file_path, os.O_RDONLY, 0666)
	if err != nil {
		return nil
	}
	return &LogsReader{
		log_file:        log_file,
		log_file_reader: bufio.NewReader(log_file),
		logs_style:      logs_style,
		line:            0,
	}
}

// 关闭生成的 LogsReader 实例的资源。
func (tho *LogsReader) Close() error {
	err := tho.log_file.Close()
	tho.log_file_reader = nil
	tho.line = 0
	if err != nil {
		return err
	}
	return nil
}

// 重定向目标日志文件和重设日志文件风格。
func (tho *LogsReader) ResetLogsReader(log_file_path string, logs_style *LogsStyle) error {
	err := tho.Close()
	if err != nil {
		return err
	}
	log_file, err := os.OpenFile(log_file_path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	tho.log_file_reader = bufio.NewReader(log_file)
	tho.logs_style = logs_style
	return nil
}

// SeekLine 设置下一次读取日志的行号。
func (tho *LogsReader) SeekLine(offset int64, whence int) (ret int64, err error) {
	ret = 0
	if offset < 0 {
		return -1, errors.New("logs: The value of offset is invalid")
	}

	if whence == 0 {
		_, err = tho.log_file.Seek(0, io.SeekStart)
		if err != nil {
			return -1, err
		}
		tho.log_file_reader = bufio.NewReader(tho.log_file)
		tho.line = 0
	}
	if whence > 1 {
		return -1, errors.New("logs: The value of whence is invalid")
	}

	for {
		if ret == offset {
			break
		}

		char, err := tho.log_file_reader.ReadByte()
		if err != nil {
			return -1, errors.New("logs: The target location does not exist")
		}

		if char == '\n' {
			ret++
		}
	}

	tho.line += ret

	return ret, nil
}

// 获取下次读取日志的行号。
func (tho *LogsReader) CurrentLine() int64 {
	return tho.line
}

// 从日志文件读取一条日志。
func (tho *LogsReader) GetALog() (a_log ALog, err error) {
	a_log_read_buffer, err := tho.log_file_reader.ReadString('\n')
	if err != nil {
		return ALog{}, err
	}
	a_log.Line = tho.line
	a_log.Log = a_log_read_buffer
	tho.line++
	return a_log, err
}

func (tho *LogsReader) findAllLogs(log_type int) ([]ALog, error) {
	all_log := make([]ALog, logs_read_size)
	if pan := recover(); pan != nil {
		return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
	}
	read_count := 0
	tho.SeekLine(0, 0)

	for {
		if read_count == len(all_log) {
			all_log_mirror := make([]ALog, len(all_log)*2)
			if pan := recover(); pan != nil {
				return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
			}
			copy(all_log_mirror, all_log)
			all_log = all_log_mirror
		}

		a_log, err := tho.GetALog()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return all_log[:read_count], err
			}
		}

		var log_prefix string
		switch log_type {
		case begin_log:
			log_prefix = tho.logs_style.PreBegin
		case end_log:
			log_prefix = tho.logs_style.PreEnd
		case debug_log:
			log_prefix = tho.logs_style.PreDebug
		case info_log:
			log_prefix = tho.logs_style.PreInfo
		case warning_log:
			log_prefix = tho.logs_style.PreWarning
		case error_log:
			log_prefix = tho.logs_style.PreError
		}

		has_prefix := false
		if log_type <= end_log {
			if (tho.logs_style.BEFlag & log.Lmsgprefix) != 0 {
				has_prefix = strings.HasSuffix(a_log.Log, log_prefix)
			} else {
				has_prefix = strings.HasPrefix(a_log.Log, log_prefix)
			}
		} else {
			if (tho.logs_style.OutputFlag & log.Lmsgprefix) != 0 {
				has_prefix = strings.HasSuffix(a_log.Log, log_prefix)
			} else {
				has_prefix = strings.HasPrefix(a_log.Log, log_prefix)
			}
		}

		if has_prefix {
			all_log[read_count] = a_log
			read_count++
		}
	}

	return all_log[:read_count], nil
}

// 查询日志文件的所有 begin 日志。
func (tho *LogsReader) FindAllBegin() ([]ALog, error) {
	return tho.findAllLogs(begin_log)
}

// 查询日志文件的所有 end 日志。
func (tho *LogsReader) FindAllEnd() ([]ALog, error) {
	return tho.findAllLogs(end_log)
}

// 查询日志文件的所有 debug 日志。
func (tho *LogsReader) FindAllDebug() ([]ALog, error) {
	return tho.findAllLogs(debug_log)
}

// 查询日志文件的所有 info 日志。
func (tho *LogsReader) FindAllInfo() ([]ALog, error) {
	return tho.findAllLogs(info_log)
}

// 查询日志文件的所有 warning 日志。
func (tho *LogsReader) FindAllWarning() ([]ALog, error) {
	return tho.findAllLogs(warning_log)
}

// 查询日志文件的所有 error 日志。
func (tho *LogsReader) FindAllError() ([]ALog, error) {
	return tho.findAllLogs(error_log)
}

// 获取指定行区间的所有日志。
func (tho *LogsReader) GetLogs(start_line int64, end_line int64) ([]ALog, error) {
	all_log := make([]ALog, logs_read_size)
	if pan := recover(); pan != nil {
		return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
	}
	read_count := 0
	tho.SeekLine(start_line, 0)

	for {
		if read_count == len(all_log) {
			all_log_mirror := make([]ALog, len(all_log)*2)
			if pan := recover(); pan != nil {
				return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
			}
			copy(all_log_mirror, all_log)
			all_log = all_log_mirror
		}

		if tho.line > end_line {
			break
		}

		a_log, err := tho.GetALog()
		if err != nil {
			if err == io.EOF {
				return all_log[:read_count], err
			}
		}

		all_log[read_count] = a_log
		read_count++
	}

	return all_log[:read_count], nil
}

func (tho *LogsReader) getLogs(log_type int, start_line int64, end_line int64) ([]ALog, error) {
	all_log := make([]ALog, logs_read_size)
	if pan := recover(); pan != nil {
		return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
	}
	read_count := 0
	tho.SeekLine(start_line, 0)

	for {
		if read_count == len(all_log) {
			all_log_mirror := make([]ALog, len(all_log)*2)
			if pan := recover(); pan != nil {
				return nil, errors.New("logs: Failed to create the memory allocation required for the buffer of all log")
			}
			copy(all_log_mirror, all_log)
			all_log = all_log_mirror
		}

		if tho.line > end_line {
			break
		}

		a_log, err := tho.GetALog()
		if err != nil {
			if err == io.EOF {
				return all_log[:read_count], err
			}
		}

		var log_prefix string
		switch log_type {
		case begin_log:
			log_prefix = tho.logs_style.PreBegin
		case end_log:
			log_prefix = tho.logs_style.PreEnd
		case debug_log:
			log_prefix = tho.logs_style.PreDebug
		case info_log:
			log_prefix = tho.logs_style.PreInfo
		case warning_log:
			log_prefix = tho.logs_style.PreWarning
		case error_log:
			log_prefix = tho.logs_style.PreError
		}

		has_prefix := false
		if log_type <= end_log {
			if (tho.logs_style.BEFlag & log.Lmsgprefix) != 0 {
				has_prefix = strings.HasSuffix(a_log.Log, log_prefix)
			} else {
				has_prefix = strings.HasPrefix(a_log.Log, log_prefix)
			}
		} else {
			if (tho.logs_style.OutputFlag & log.Lmsgprefix) != 0 {
				has_prefix = strings.HasSuffix(a_log.Log, log_prefix)
			} else {
				has_prefix = strings.HasPrefix(a_log.Log, log_prefix)
			}
		}

		if has_prefix {
			all_log[read_count] = a_log
			read_count++
		}
	}

	return all_log[:read_count], nil
}

// 获取指定行区间的所有 debug 日志。
func (tho *LogsReader) GetDebug(start_line int64, end_line int64) ([]ALog, error) {
	return tho.getLogs(debug_log, start_line, end_line)
}

// 获取指定行区间的所有 info 日志。
func (tho *LogsReader) GetInfo(start_line int64, end_line int64) ([]ALog, error) {
	return tho.getLogs(info_log, start_line, end_line)
}

// 获取指定行区间的所有 warning 日志。
func (tho *LogsReader) GetWarning(start_line int64, end_line int64) ([]ALog, error) {
	return tho.getLogs(warning_log, start_line, end_line)
}

// 获取指定行区间的所有 error 日志。
func (tho *LogsReader) GetError(start_line int64, end_line int64) ([]ALog, error) {
	return tho.getLogs(error_log, start_line, end_line)
}
