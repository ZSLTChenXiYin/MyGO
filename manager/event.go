package manager

import (
	"errors"
	"os"
	"os/signal"
)

// 为 manager 添加 event，event_signal 为一个进程信号量， event 为该接收到 event_signal 进程信号量所执行的方法。
func Event(event func() bool, event_signal os.Signal) error {
	if int(events_count) == len(events) {
		map_events := make(map[os.Signal]func() bool, len(events)*2)
		if pan := recover(); pan != nil {
			map_events = make(map[os.Signal]func() bool, len(events)+1)
			if pan = recover(); pan != nil {
				return errors.New("manager: Failed to create the memory allocation required for the events of manager")
			}
		}
		for key, value := range events {
			map_events[key] = value
		}
		events = map_events
	}
	events_count++
	signal.Notify(signal_channel, event_signal)
	events[event_signal] = event
	return nil
}

// 为 manager 添加 event，event_signals 为多个进程信号量， event 为该接收到 event_signals 进程信号量中任意一种时所执行的方法。
func Events(event func() bool, event_signals ...os.Signal) error {
	if int(events_count) == len(events) {
		map_events := make(map[os.Signal]func() bool, len(events)+len(event_signals))
		if pan := recover(); pan != nil {
			return errors.New("manager: Failed to create the memory allocation required for the events of manager")
		}
		for key, value := range events {
			map_events[key] = value
		}
		events = map_events
	}
	events_count += uint(len(event_signals))
	signal.Notify(signal_channel, event_signals...)
	for _, event_signal := range event_signals {
		events[event_signal] = event
	}
	return nil
}
