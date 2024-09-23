# MyGO logs #
```go
import "github.com/ZSLTChenXiYin/MyGO/logs"
```
logs 实现了异步日志框架和日志查询组件。
## 索引 ##
* [type LogsStyle](#type-logsstyle)
> * [func NewLogsStyle() *LogsStyle](#func-newlogsstyle)
* [func UseLogs(used_log_file *os.File, logs_style *LogsStyle) error](#func-uselogs)
* [func UseDefault(used_log_file *os.File) error](#func-usedefault)
* [func ReleaseLogs()](#func-releaselogs)
* [func ReuseOutput(output_file *os.File) error](#func-reuseoutput)
* [func OpenLogs(log_file_path string, logs_style *LogsStyle) error](#func-openlogs)
* [func OpenDefault(log_file_path string) error](#func-opendefault)
* [func CloseLogs() error](#func-closelogs)
* [func ReopenOutput(output_path string) error](#func-reopenoutput)
* [func GetLogsInfo() (string, *LogsStyle)](#func-getlogsinfo)
* [func Run() error](#func-run)
* [func Over()](#func-over)
* [func Debug(debug string) (uintptr, error)](#func-debug)
* [func Info(info string) (uintptr, error)](#func-info)
* [func Warning(warning string) (uintptr, error)](#func-warning)
* [func Error(err error) (uintptr, error)](#func-error)
* [type ALog](#type-alog)
* [type LogsReader](#type-logsreader)
> * [func NewLogsReader(log_file_path string, logs_style *LogsStyle) *LogsReader](#func-newlogsreader)
> * [func (tho *LogsReader) Close() error](#func-tho-logsreader-close)
> * [func (tho *LogsReader) ResetLogsReader(log_file_path string, logs_style *LogsStyle) error](#func-tho-logsreader-resetlogsreader)
> * [func (tho *LogsReader) SeekLine(offset int64, whence int) (ret int64, err error)](#func-tho-logsreader-seekline)
> * [func (tho *LogsReader) CurrentLine() int64](#func-tho-logsreader-currentline)
> * [func (tho *LogsReader) GetALog() (a_log ALog, err error)](#func-tho-logsreader-getalog)
> * [func (tho *LogsReader) FindAllBegin() ([]ALog, error)](#func-tho-logsreader-findallbegin)
> * [func (tho *LogsReader) FindAllEnd() ([]ALog, error)](#func-tho-logsreader-findallend)
> * [func (tho *LogsReader) FindAllDebug() ([]ALog, error)](#func-tho-logsreader-findalldebug)
> * [func (tho *LogsReader) FindAllInfo() ([]ALog, error)](#func-tho-logsreader-findallinfo)
> * [func (tho *LogsReader) FindAllWarning() ([]ALog, error)](#func-tho-logsreader-findallwarning)
> * [func (tho *LogsReader) FindAllError() ([]ALog, error)](#func-tho-logsreader-findallerror)
> * [func (tho *LogsReader) GetLogs(start_line int64, end_line int64) ([]ALog, error)](#func-tho-logsreader-getlogs)
> * [func (tho *LogsReader) GetDebug(start_line int64, end_line int64) ([]ALog, error)](#func-tho-logsreader-getdebug)
> * [func (tho *LogsReader) GetInfo(start_line int64, end_line int64) ([]ALog, error)](#func-tho-logsreader-getinfo)
> * [func (tho *LogsReader) GetWarning(start_line int64, end_line int64) ([]ALog, error)](#func-tho-logsreader-getwarning)
> * [func (tho *LogsReader) GetError(start_line int64, end_line int64) ([]ALog, error)](#func-tho-logsreader-geterror)
## API ##
#### type LogsStyle
```go
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
```
LogsStyle 用于决定 logs 服务中的日志风格。LogsChannelSize 决定日志写入管道的大小，BEFlag 决定输出 begin 和 end 日志的风格，PreBegin 和 PreEnd 为 begin 和 end 日志的前缀，OutputFlag 决定输出 debug、info、warning 和 error 等日志的风格，PreDebug、PreInfo、PreWarning 和 PreError 分别为 debug、info、warning 和 error 日志的前缀。BEFlag 和 OutputFlag 可以使用 log 包中的 flag 赋值。
#### func NewLogsStyle
```go
func NewLogsStyle() *LogsStyle
```
生成一个默认风格的日志风格对象。
#### func UseLogs
```go
func UseLogs(used_log_file *os.File, logs_style *LogsStyle) error
```
使用一个已打开的可写文件启动指定风格的 logs 服务。使用时一般会与 ReleaseLogs 方法连用。
#### func UseDefault
```go
func UseDefault(used_log_file *os.File) error
```
使用一个已打开的可写文件启动默认风格的 logs 服务。使用时一般会与 ReleaseLogs 方法连用。
#### func ReleaseLogs
```go
func ReleaseLogs()
```
释放通过已打开的可写文件启动的 logs 服务的资源。
#### func ReuseOutput
```go
func ReuseOutput(output_file *os.File) error
```
重定向通过 UseLogs 或 UseDefault 打开的 logs 服务的输出日志文件。
#### func OpenLogs
```go
func OpenLogs(log_file_path string, logs_style *LogsStyle) error
```
通过日志文件路径启动指定风格的 logs 服务。使用时一般会与 CloseLogs 方法连用。
#### func OpenDefault
```go
func OpenDefault(log_file_path string) error
```
通过日志文件路径启动默认风格的 logs 服务。使用时一般会与 CloseLogs 方法连用。
#### func CloseLogs
```go
func CloseLogs() error
```
释放通过日志文件路径启动的 logs 服务的资源。
#### func ReopenOutput
```go
func ReopenOutput(output_path string) error
```
重定向通过 OpenLogs 或 OpenDefault 打开的 logs 服务的输出日志文件。
#### func GetLogsInfo
```go
func GetLogsInfo() (string, *LogsStyle)
```
获取当前启动的 logs 服务相关信息。
#### func Run
```go
func Run() error
```
运行已启动的 logs 服务。使用时一般会与 Over 方法连用。logs 服务可以通过 UseLogs、UseDefault、OpenLogs 或 OpenDefault 等方法启动。
#### func Over
```go
func Over()
```
停止正在运行的 logs 服务。
#### func Debug
```go
func Debug(debug string) (uintptr, error)
```
写入一条 debug 日志。
#### func Info
```go
func Info(info string) (uintptr, error)
```
写入一条 info 日志。
#### func Warning
```go
func Warning(warning string) (uintptr, error)
```
写入一条 warning 日志。
#### func Error
```go
func Error(err error) (uintptr, error)
```
写入一条 error 日志。
#### type ALog
```go
type ALog struct {
  Line int64
  Log  string
}
```
存放读取的日志的结构。Line 为该条日志存在日志文件的行号，Log 为日志内容。
#### type LogsReader
```go
type LogsReader struct {
  // 内部隐含字段，详情可见于 read_logs.go
}
```
LogsReader 实现了读取日志的一系列方法。
#### func NewLogsReader
```go
func NewLogsReader(log_file_path string, logs_style *LogsStyle) *LogsReader
```
生成一个 LogsReader 实例。停止使用前注意调用 Close 方法释放占用的资源。
#### func (tho *LogsReader) Close
```go
func (tho *LogsReader) Close() error
```
关闭生成的 LogsReader 实例的资源。
#### func (tho *LogsReader) ResetLogsReader
```go
func (tho *LogsReader) ResetLogsReader(log_file_path string, logs_style *LogsStyle) error
```
重定向目标日志文件和重设日志文件风格。
#### func (tho *LogsReader) SeekLine
```go
func (tho *LogsReader) SeekLine(offset int64, whence int) (ret int64, err error)
```
SeekLine 设置下一次读取日志的行号。
#### func (tho *LogsReader) CurrentLine
```go
func (tho *LogsReader) CurrentLine() int64
```
获取下次读取日志的行号。
#### func (tho *LogsReader) GetALog
```go
func (tho *LogsReader) GetALog() (a_log ALog, err error)
```
从日志文件读取一条日志。
#### func (tho *LogsReader) FindAllBegin
```go
func (tho *LogsReader) FindAllBegin() ([]ALog, error)
```
查询日志文件的所有 begin 日志。
#### func (tho *LogsReader) FindAllEnd
```go
func (tho *LogsReader) FindAllEnd() ([]ALog, error)
```
查询日志文件的所有 end 日志。
#### func (tho *LogsReader) FindAllDebug
```go
func (tho *LogsReader) FindAllDebug() ([]ALog, error)
```
查询日志文件的所有 debug 日志。
#### func (tho *LogsReader) FindAllInfo
```go
func (tho *LogsReader) FindAllInfo() ([]ALog, error)
```
查询日志文件的所有 info 日志。
#### func (tho *LogsReader) FindAllWarning
```go
func (tho *LogsReader) FindAllWarning() ([]ALog, error)
```
查询日志文件的所有 warning 日志。
#### func (tho *LogsReader) FindAllError
```go
func (tho *LogsReader) FindAllError() ([]ALog, error)
```
查询日志文件的所有 error 日志。
#### func (tho *LogsReader) GetLogs
```go
func (tho *LogsReader) GetLogs(start_line int64, end_line int64) ([]ALog, error)
```
获取指定行区间的所有日志。
#### func (tho *LogsReader) GetDebug
```go
func (tho *LogsReader) GetDebug(start_line int64, end_line int64) ([]ALog, error)
```
获取指定行区间的所有 debug 日志。
#### func (tho *LogsReader) GetInfo
```go
func (tho *LogsReader) GetInfo(start_line int64, end_line int64) ([]ALog, error)
```
获取指定行区间的所有 info 日志。
#### func (tho *LogsReader) GetWarning
```go
func (tho *LogsReader) GetWarning(start_line int64, end_line int64) ([]ALog, error)
```
获取指定行区间的所有 warning 日志。
#### func (tho *LogsReader) GetError
```go
func (tho *LogsReader) GetError(start_line int64, end_line int64) ([]ALog, error)
```
获取指定行区间的所有 error 日志。