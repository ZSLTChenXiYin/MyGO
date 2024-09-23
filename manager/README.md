# MyGO manager #
```go
import "github.com/ZSLTChenXiYin/MyGO/manager"
```
manager 包实现了用于进程信号量管理的简单框架。
## 索引 ##
* [func CreateManager(channel_size int, events_size int) error](#func-createmanager)
* [func Run()](#func-run)
* [func GetManagerInfo() (int, int)](#func-getmanagerinfo)
* [func Event(event func() bool, event_signal os.Signal) error](#func-event)
* [func Events(event func() bool, event_signals ...os.Signal) error](#func-events)
## API ##
#### func CreateManager
```go
func CreateManager(channel_size int, events_size int) error
```
创建 manager，完成初始化工作。channel_size 决定信号量管道大小，events_size 决定 manager 存放 event 的最大容量。
#### func Run
```go
func Run()
```
运行 manager，自动接收进程中的信号量，按指定的 event 处理信号量，此方法一般用于 main 方法的末尾。
#### func GetManagerInfo
```go
func GetManagerInfo() (int, int)
```
获取 manager 的信息，分别为 manager_channel_size 和 manager_events_size。
#### func Event
```go
func Event(event func() bool, event_signal os.Signal) error
```
为 manager 添加 event，event_signal 为一个进程信号量， event 为该接收到 event_signal 进程信号量所执行的方法。
#### func Events
```go
func Events(event func() bool, event_signals ...os.Signal) error
```
为 manager 添加 event，event_signals 为多个进程信号量， event 为该接收到 event_signals 进程信号量中任意一种时所执行的方法。