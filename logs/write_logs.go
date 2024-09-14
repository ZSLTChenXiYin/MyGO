package logs

// write log head
func begin() {
	begin_logger.Println()
}

// write log tail
func end() {
	end_logger.Println()
}

func Debug(debug string) {
	debug_logger.Println(debug)
}

func Info(info string) {
	info_logger.Println(info)
}

func Warning(warning string) {
	warning_logger.Println(warning)
}

func Error(err error) {
	error_logger.Println(err)
}

func Debugf(format string, value ...any) {
	debug_logger.Printf(format, value...)
}

func Infof(format string, value ...any) {
	info_logger.Printf(format, value...)
}

func Warningf(format string, value ...any) {
	warning_logger.Printf(format, value...)
}

func Errorf(format string, value ...any) {
	error_logger.Printf(format, value...)
}
