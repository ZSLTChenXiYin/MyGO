package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/ZSLTChenXiYin/MyGO/configure"
)

const (
	color_red    = "\033[91m"
	color_green  = "\033[92m"
	color_yellow = "\033[93m"
	color_blue   = "\033[94m"
	color_white  = "\033[0m"
)

func stdRedString(str string) string {
	return color_red + str + color_white
}

func stdGreenString(str string) string {
	return color_green + str + color_white
}

func stdYellowString(str string) string {
	return color_yellow + str + color_white
}

func stdBlueString(str string) string {
	return color_blue + str + color_white
}

const (
	std_error_prefix_text = "ERROR"
	std_info_prefix_text  = "INFO"
	std_warn_prefix_text  = "WARN"
	std_debug_prefix_text = "DEBUG"
)

func newErrorLogger() *log.Logger {
	return log.New(os.Stderr, fmt.Sprintf("[%s] ", stdRedString(std_error_prefix_text)), log.LstdFlags)
}

func newInfoLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", stdGreenString(std_info_prefix_text)), log.LstdFlags)
}

func newWarnLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", stdYellowString(std_warn_prefix_text)), log.LstdFlags)
}

func newDebugLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", stdBlueString(std_debug_prefix_text)), log.LstdFlags)
}

type StdLogger struct {
	conf         configure.Configuration
	error_logger *log.Logger
	info_logger  *log.Logger
	warn_logger  *log.Logger
	debug_logger *log.Logger
}

func NewStdLogger(conf configure.Configuration) *StdLogger {
	return &StdLogger{
		conf:         conf,
		error_logger: newErrorLogger(),
		info_logger:  newInfoLogger(),
		warn_logger:  newWarnLogger(),
		debug_logger: newDebugLogger(),
	}
}

func (sl *StdLogger) ErrorLogger() *log.Logger {
	return sl.error_logger
}

func (sl *StdLogger) InfoLogger() *log.Logger {
	return sl.info_logger
}

func (sl *StdLogger) WarnLogger() *log.Logger {
	return sl.warn_logger
}

func (sl *StdLogger) DebugLogger() *log.Logger {
	return sl.debug_logger
}

func (sl *StdLogger) Errorf(format string, v ...any) {
	sl.error_logger.Printf(format, v...)
}

func (sl *StdLogger) Infof(format string, v ...any) {
	sl.info_logger.Printf(format, v...)
}

func (sl *StdLogger) Warnf(format string, v ...any) {
	sl.warn_logger.Printf(format, v...)
}

func (sl *StdLogger) Debugf(format string, v ...any) {
	if sl.conf.Server().Debug() {
		sl.debug_logger.Printf(format, v...)
	}
}
