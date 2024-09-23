package manager

import (
	"errors"
	"os"
)

var (
	manager_channel_size int
	manager_events_size  int
)

var (
	signal_channel chan os.Signal
	events_count   uint
	events         map[os.Signal]func() bool
)

// 创建 manager，完成初始化工作。channel_size 决定信号量管道大小，events_size 决定 manager 存放 event 的最大容量。
func CreateManager(channel_size int, events_size int) error {
	signal_channel = make(chan os.Signal, channel_size)
	if pan := recover(); pan != nil {
		return errors.New("manager: Failed to create the memory allocation required for the manager")
	}
	events = make(map[os.Signal]func() bool, events_size)
	if pan := recover(); pan != nil {
		return errors.New("manager: Failed to create the memory allocation required for the manager")
	}

	manager_channel_size = channel_size
	manager_events_size = events_size

	return nil
}

// 运行 manager，自动接收进程中的信号量，按指定的 event 处理信号量，此方法一般用于 main 方法的末尾。
func Run() {
	for event_signal := range signal_channel {
		if events[event_signal]() {
			break
		}
	}
}

// 获取 manager 的信息，分别为 manager_channel_size 和 manager_events_size。
func GetManagerInfo() (channel_size int, events_size int) {
	return manager_channel_size, manager_events_size
}
