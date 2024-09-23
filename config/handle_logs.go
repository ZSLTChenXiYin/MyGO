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

	LogsChannelSize uint

	BEFlags  flags
	PreBegin string
	PreEnd   string

	OutputFlags flags
	PreDebug    string
	PreInfo     string
	PreWarning  string
	PreError    string
}

func (tho *logsConfig) getLogsStyle() *logs.LogsStyle {
	return &logs.LogsStyle{
		LogsChannelSize: tho.LogsChannelSize,

		BEFlag:   getFlagsNum(tho.BEFlags),
		PreBegin: tho.PreBegin,
		PreEnd:   tho.PreEnd,

		OutputFlag: getFlagsNum(tho.OutputFlags),
		PreDebug:   tho.PreDebug,
		PreInfo:    tho.PreInfo,
		PreWarning: tho.PreWarning,
		PreError:   tho.PreError,
	}
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

// 将指定 json 文件配置导入 logs 框架。
func ImportLogsConfig(config_path string) error {
	var logs_config logsConfig
	err := ImportConfig(&logs_config, config_path)
	if err != nil {
		return err
	}

	err = logs.OpenLogs(logs_config.LogFilePath, logs_config.getLogsStyle())
	if err != nil {
		return err
	}

	err = logs.Run()
	if err != nil {
		return err
	}

	return nil
}

func getLogsConfig() *logsConfig {
	log_file_path, logs_style := logs.GetLogsInfo()
	return &logsConfig{
		LogFilePath:     log_file_path,
		LogsChannelSize: logs_style.LogsChannelSize,
		BEFlags:         getFlags(logs_style.BEFlag),
		PreBegin:        logs_style.PreBegin,
		PreEnd:          logs_style.PreEnd,
		OutputFlags:     getFlags(logs_style.OutputFlag),
		PreDebug:        logs_style.PreDebug,
		PreInfo:         logs_style.PreInfo,
		PreWarning:      logs_style.PreWarning,
		PreError:        logs_style.PreError,
	}
}

// 将 logs 框架配置导出至指定 json 文件。
func ExportLogsConfig(config_path string) error {
	err := ExportConfig(getLogsConfig(), config_path)
	if err != nil {
		return err
	}

	return nil
}

// 将指定 json 文件配置导入并生成 logs reader。
func ImportLogsReaderConfig(config_path string) (logs_reader *logs.LogsReader, err error) {
	var logs_config logsConfig
	err = ImportConfig(&logs_config, config_path)
	if err != nil {
		return nil, err
	}
	return logs.NewLogsReader(logs_config.LogFilePath, logs_config.getLogsStyle()), nil
}
