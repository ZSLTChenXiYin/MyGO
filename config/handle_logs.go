package config

import (
	"log"

	"github.com/ZSLTChenXiYin/MyGO/logs"
)

type flags struct {
	Date         bool
	Time         bool
	Microseconds bool
	Longfile     bool
	Shortfile    bool
	UTC          bool
	MsgPrefix    bool
	STDFlags     bool
}

type logsConfig struct {
	LogFilePath string

	BEFlags  flags
	PreBegin string
	PreEnd   string

	Flags      flags
	PreDebug   string
	PreInfo    string
	PreWarning string
	PreError   string
}

func getFlagsNum(flags flags) int {
	var flags_num int
	if flags.Date {
		flags_num |= log.Ldate
	}
	if flags.Time {
		flags_num |= log.Ltime
	}
	if flags.Microseconds {
		flags_num |= log.Lmicroseconds
	}
	if flags.Longfile {
		flags_num |= log.Llongfile
	}
	if flags.Shortfile {
		flags_num |= log.Lshortfile
	}
	if flags.UTC {
		flags_num |= log.LUTC
	}
	if flags.MsgPrefix {
		flags_num |= log.Lmsgprefix
	}
	if flags.STDFlags {
		flags_num |= log.LstdFlags
	}

	return flags_num
}

func getFlags(flags_num int) flags {
	var flags flags
	flags.Date = (flags_num & log.Ldate) != 0
	flags.Time = (flags_num & log.Ltime) != 0
	flags.Microseconds = (flags_num & log.Lmicroseconds) != 0
	flags.Longfile = (flags_num & log.Llongfile) != 0
	flags.Shortfile = (flags_num & log.Lshortfile) != 0
	flags.UTC = (flags_num & log.LUTC) != 0
	flags.MsgPrefix = (flags_num & log.Lmsgprefix) != 0
	flags.STDFlags = (flags_num & log.LstdFlags) != 0

	return flags
}

func getLogsConfig() *logsConfig {
	log_file_path, beflag, pre_begin, pre_end, flag, pre_debug, pre_info, pre_warning, pre_error := logs.GetUserInfo()
	return &logsConfig{
		LogFilePath: log_file_path,
		BEFlags:     getFlags(beflag),
		PreBegin:    pre_begin,
		PreEnd:      pre_end,
		Flags:       getFlags(flag),
		PreDebug:    pre_debug,
		PreInfo:     pre_info,
		PreWarning:  pre_warning,
		PreError:    pre_error,
	}
}

// auto import logs config
func ImportLogsConfig(config_path string) error {
	var logs_config logsConfig
	err := ImportConfig(&logs_config, config_path)
	if err != nil {
		return err
	}

	err = logs.OpenLogs(logs_config.LogFilePath, getFlagsNum(logs_config.BEFlags), getFlagsNum(logs_config.Flags), logs_config.PreBegin, logs_config.PreEnd, logs_config.PreDebug, logs_config.PreInfo, logs_config.PreWarning, logs_config.PreError)
	if err != nil {
		return err
	}

	return nil
}

// auto export logs config
func ExportLogsConfig(config_path string) error {
	err := ExportConfig(getLogsConfig(), config_path)
	if err != nil {
		return err
	}

	return nil
}
