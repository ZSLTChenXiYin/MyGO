package config

import "github.com/ZSLTChenXiYin/MyGO/manager"

type managerConfig struct {
	ChannelSize int
	EventsSize  int
}

func getManagerConfig() *managerConfig {
	channel_size, events_size := manager.GetUserInfo()
	return &managerConfig{
		ChannelSize: channel_size,
		EventsSize:  events_size,
	}
}

// auto import manager config
func ImportManagerConfig(config_path string) error {
	var manager_config managerConfig
	err := ImportConfig(&manager_config, config_path)
	if err != nil {
		return err
	}

	err = manager.CreateManager(manager.GetUserInfo())
	if err != nil {
		return err
	}

	return nil
}

// auto export manager config
func ExportManagerConfig(config_path string) error {
	err := ExportConfig(getManagerConfig(), config_path)
	if err != nil {
		return err
	}

	return nil
}
