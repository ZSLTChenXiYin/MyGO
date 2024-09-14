package manager

import (
	"errors"
	"os"
	"os/signal"
)

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
