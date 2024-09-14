package logs

import (
	"errors"
	"io"
	"log"
	"os"
)

const (
	default_beflag = log.Ldate | log.Ltime

	default_pre_begin = ">> Log begin: "
	default_pre_end   = "<< Log end: "

	default_flag = log.Ldate | log.Ltime

	default_pre_debug   = "[DEBUG] "
	default_pre_info    = "[INFO] "
	default_pre_warning = "[WARNING] "
	default_pre_error   = "[ERROR] "
)

var (
	user_log_file_path string

	user_beflag int

	user_pre_begin string
	user_pre_end   string

	user_flag int

	user_pre_debug   string
	user_pre_info    string
	user_pre_warning string
	user_pre_error   string
)

var (
	log_file *os.File

	// log_file_reader io.Reader
	log_file_writer io.Writer

	begin_logger *log.Logger
	end_logger   *log.Logger

	debug_logger   *log.Logger
	info_logger    *log.Logger
	warning_logger *log.Logger
	error_logger   *log.Logger
)

func setUserInfo(log_file_path string, beflag int, flag int, pre_begin string, pre_end string, pre_debug string, pre_info string, pre_warning string, pre_error string) {
	user_log_file_path = log_file_path
	user_beflag = beflag
	user_flag = flag
	user_pre_begin = pre_begin
	user_pre_end = pre_end
	user_pre_debug = pre_debug
	user_pre_info = pre_info
	user_pre_warning = pre_warning
	user_pre_error = pre_error
}

func UseLogs(used_log_file *os.File, beflag int, flag int, pre_begin string, pre_end string, pre_debug string, pre_info string, pre_warning string, pre_error string) error {
	if used_log_file == nil {
		return errors.New("logs: Used log file does not exist")
	}
	log_file = used_log_file

	// log_file_reader = log_file
	log_file_writer = log_file

	begin_logger = log.New(log_file_writer, pre_begin, beflag)
	end_logger = log.New(log_file_writer, pre_end, beflag)

	debug_logger = log.New(log_file_writer, pre_debug, flag)
	info_logger = log.New(log_file_writer, pre_info, flag)
	warning_logger = log.New(log_file_writer, pre_warning, flag)
	error_logger = log.New(log_file_writer, pre_error, flag)

	setUserInfo("", beflag, flag, pre_begin, pre_end, pre_debug, pre_info, pre_warning, pre_error)

	begin()

	return nil
}

func UseDefault(used_log_file *os.File) error {
	err := UseLogs(used_log_file, default_beflag, default_flag, default_pre_begin, default_pre_end, default_pre_debug, default_pre_info, default_pre_warning, default_pre_error)
	if err != nil {
		return err
	}
	return nil
}

func OpenLogs(log_file_path string, beflag int, flag int, pre_begin string, pre_end string, pre_debug string, pre_info string, pre_warning string, pre_error string) error {
	var err error
	log_file, err = os.OpenFile(log_file_path, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log_file, err = os.OpenFile(log_file_path, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
	}

	// log_file_reader = log_file
	log_file_writer = log_file

	begin_logger = log.New(log_file_writer, pre_begin, beflag)
	end_logger = log.New(log_file_writer, pre_end, beflag)

	debug_logger = log.New(log_file_writer, pre_debug, flag)
	info_logger = log.New(log_file_writer, pre_info, flag)
	warning_logger = log.New(log_file_writer, pre_warning, flag)
	error_logger = log.New(log_file_writer, pre_error, flag)

	setUserInfo(log_file_path, beflag, flag, pre_begin, pre_end, pre_debug, pre_info, pre_warning, pre_error)

	begin()

	return nil
}

func Default(log_file_path string) error {
	err := OpenLogs(log_file_path, default_beflag, default_flag, default_pre_begin, default_pre_end, default_pre_debug, default_pre_info, default_pre_warning, default_pre_error)
	if err != nil {
		return err
	}
	return nil
}

func CloseLogs() {
	end()
	log_file.Close()
}

func ReuseOutput(output_file *os.File) error {
	if output_file == nil {
		return errors.New("logs: Output file does not exist")
	}

	end()

	log_file = output_file

	// log_file_reader = log_file
	log_file_writer = log_file

	debug_logger.SetOutput(log_file_writer)
	info_logger.SetOutput(log_file_writer)
	warning_logger.SetOutput(log_file_writer)
	error_logger.SetOutput(log_file_writer)

	setUserInfo("", user_beflag, user_flag, user_pre_begin, user_pre_end, user_pre_debug, user_pre_info, user_pre_warning, user_pre_error)

	begin()

	return nil
}

func ResetOutput(output_path string) error {
	CloseLogs()

	var err error
	log_file, err = os.OpenFile(output_path, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log_file, err = os.OpenFile(output_path, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
	}

	// log_file_reader = log_file
	log_file_writer = log_file

	debug_logger.SetOutput(log_file_writer)
	info_logger.SetOutput(log_file_writer)
	warning_logger.SetOutput(log_file_writer)
	error_logger.SetOutput(log_file_writer)

	setUserInfo(output_path, user_beflag, user_flag, user_pre_begin, user_pre_end, user_pre_debug, user_pre_info, user_pre_warning, user_pre_error)

	begin()

	return nil
}

func GetUserInfo() (log_file_path string, beflag int, pre_begin string, pre_end string, flag int, pre_debug string, pre_info string, pre_warning string, pre_error string) {
	return user_log_file_path, user_beflag, user_pre_begin, user_pre_end, user_flag, user_pre_debug, user_pre_info, user_pre_warning, user_pre_error
}
