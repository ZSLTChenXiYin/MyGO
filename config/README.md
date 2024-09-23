# MyGO config #
```go
import "github.com/ZSLTChenXiYin/MyGO/config"
```
config 包实现了对 json 文件的便捷读取，以及对 logs 和 manager 框架的快速配置。
## 索引 ##
* [func ImportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error](#func-importconfig)
* [func ExportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error](#func-exportconfig)
* [func ImportLogsConfig(config_path string) error](#func-importlogsconfig)
* [func ExportLogsConfig(config_path string) error](#func-exportlogsconfig)
* [func ImportLogsReaderConfig(config_path string) (logs_reader *logs.LogsReader, err error)](#func-importlogsreaderconfig)
* [func ImportManagerConfig(config_path string) error](#func-importmanagerconfig)
* [func ExportManagerConfig(config_path string) error](#func-exportmanagerconfig)
## API ##
#### func ImportConfig
```go
func ImportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error
```
将 json 文件导入指定结构。
#### func ExportConfig
```go
func ExportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error
```
将指定结构导出至 json 文件。
#### func ImportLogsConfig
```go
func ImportLogsConfig(config_path string) error
```
将指定 json 文件配置导入 logs 框架。
#### func ExportLogsConfig
```go
func ExportLogsConfig(config_path string) error
```
将 logs 框架配置导出至指定 json 文件。
#### func ImportLogsReaderConfig
```go
func ImportLogsReaderConfig(config_path string) (logs_reader *logs.LogsReader, err error)
```
将指定 json 文件配置导入并生成 logs reader。
#### func ImportManagerConfig
```go
func ImportManagerConfig(config_path string) error
```
将指定 json 文件配置导入 manager 框架。
#### func ExportManagerConfig
```go
func ExportManagerConfig(config_path string) error
```
将 manager 框架配置导出至指定 json 文件。