# MyGO #
本项目为Golang开发者提供了一个轻量级服务端程序开发的工具库，其中包含日志服务框架、信号量管理框架、json文件配置导入导出等开发工具，能帮助开发者快速搭建日志服务、信号量管理等服务模块。
## 内容导引 ##
* [介绍](#介绍)
* [部署](#部署)
* [快速上手](#快速上手)
* [问题反馈](#问题反馈)
## 介绍 ##
#### 目前为止，MyGO 可以提供以下支持：
* 完整的日志服务框架，可以将程序运行日志输出至指定文件（包括标准输出文件、标准错误输出文件）
* 简洁的信号量管理框架，可以用于管理进程接受到的信号量（例如接收到用户键入终端中断字符Ctrl+C后，管理器会执行您预先设置的方法，帮您更简洁和优雅地完成收尾工作）
* 方便的json文件配置导入导出模块，可以轻松地读取配置文件，模块中已经提供了日志服务、信号量管理服务等工具的快速部署和导出方法
## 部署 ##
MyGO 的部署依赖 Go modules，如果你还没有 go mod，你需要首先初始化:
```sh
go mod init myproject
```
安装 MyGO
```sh
go get -u github.com/ZSLTChenXiYin/MyGO
```
## 快速上手 ##
### 使用日志服务框架
> 使用默认配置将日志信息输出至标准输出文件
```go
package main

import (
  "os"

  "github.com/ZSLTChenXiYin/MyGO/logs"
)

func main() {
  // open log serve with stdout file
  logs.UseDefault(os.Stdout)
  defer logs.CloseLogs()

  logs.Debug("The log service has been enabled")
}
```
> 使用默认配置将日志信息输出至指定配置文件
```go
package main

import "github.com/ZSLTChenXiYin/MyGO/logs"

func main() {
  // open log serve with log file
  logs.Default("log")
  defer logs.CloseLogs()

  logs.Debug("The log service has been enabled")
}
```
### 使用信号量管理框架
> 用管理器处理收尾工作
```go
package main

import (
  "syscall"
  "time"

  "github.com/ZSLTChenXiYin/MyGO/logs"
  "github.com/ZSLTChenXiYin/MyGO/manager"
)

func MyServe() {
  // simulate service
  for {
    logs.Info("Serving users")
    time.Sleep(time.Second * 10)
  }
}

func main() {
  // open log serve
  logs.Default("log")

  // create manager serve
  manager.CreateManager(1, 1)
  // add events
  manager.Events(func() bool {
    logs.Debug("The process has stopped running")
    logs.CloseLogs()
    return true
  }, syscall.SIGINT, syscall.SIGTERM)

  // run my work serve
  go MyServe()

  // run manager serve
  manager.Run()
}
```
### 使用快速配置构建项目
> 快速配置并构建日志服务和管理器
>> logs_config.json
```json
{
  "LogFilePath": "log",

  "BEFlags": {
    "Date": true,
    "Time": true,
    "Microseconds": true,
    "Longfile": false,
    "Shortfile": false,
    "UTC": false,
    "MsgPrefix": true,
    "STDFlags": false
  },
  "PreBegin": "=> logs begin: ",
  "PreEnd": "<= logs end: ",

  "Flags": {
    "Date": true,
    "Time": true,
    "Microseconds": false,
    "Longfile": false,
    "Shortfile": false,
    "UTC": false,
    "MsgPrefix": true,
    "STDFlags": false
  },
  "PreDebug": "[debug] ",
  "PreInfo": "#info# ",
  "PreWarning": "|warning| ",
  "PreError": "<error> "
}
```
>> manager_config.json
```json
{
  "ChannelSize": 1,
  "EventsSize": 1
}
```
>> logs_config.go
```go
package main

import (
  "syscall"

  "github.com/ZSLTChenXiYin/MyGO/config"
  "github.com/ZSLTChenXiYin/MyGO/logs"
  "github.com/ZSLTChenXiYin/MyGO/manager"
)

func main() {
  config.ImportLogsConfig("logs_config.json")

  config.ImportManagerConfig("manager_config.json")
  manager.Events(func() bool {
    logs.Debug("The process has stopped running")
    logs.CloseLogs()
    return true
  }, syscall.SIGINT, syscall.SIGTERM)

  logs.Debug("The log service has been enabled")

  manager.Run()
}
```
## 问题反馈 ##
* 陈汐胤会在每周五至周日查看 [Issues](https://github.com/ZSLTChenXiYin/MyGO/issues)，还会不定期地在 bilibili 直播
>> 陈汐胤的 e-mail: imjfoy@163.com
>> 
>> 陈汐胤的 bilibili UID: 352456302
## 引用 ##
[CodE Dream! It's MyGO!!!!!](https://www.bing.com/search?q=BanG+Dream%21+It%27s+MyGO%21%21%21%21%21&qs=n&form=QBRE&sp=-1&lq=0&pq=bang+dream%21+it%27s+mygo%21%21%21%21%21&sc=9-26&sk=&cvid=E85531D5035D4A0CA14ED56EAD735E44&ghsh=0&ghacc=0&ghpl=)
