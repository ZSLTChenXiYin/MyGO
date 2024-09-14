package manager

import (
	"errors"
	"os"
)

var (
	user_channel_size int
	user_events_size  int
)

var (
	signal_channel chan os.Signal
	events_count   uint
	events         map[os.Signal]func() bool
)

func CreateManager(channel_size int, events_size int) error {
	signal_channel = make(chan os.Signal, channel_size)
	if pan := recover(); pan != nil {
		return errors.New("manager: Failed to create the memory allocation required for the manager")
	}
	events = make(map[os.Signal]func() bool, events_size)
	if pan := recover(); pan != nil {
		return errors.New("manager: Failed to create the memory allocation required for the manager")
	}

	user_channel_size = channel_size
	user_events_size = events_size

	return nil
}

func Run() {
	for event_signal := range signal_channel {
		if events[event_signal]() {
			break
		}
	}
}

func GetUserInfo() (int, int) {
	return user_channel_size, user_events_size
}
