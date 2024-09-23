package config

import "github.com/ZSLTChenXiYin/MyGO/manager"

type managerConfig struct {
	ChannelSize int
	EventsSize  int
}

func getManagerConfig() *managerConfig {
	channel_size, events_size := manager.GetManagerInfo()
	return &managerConfig{
		ChannelSize: channel_size,
		EventsSize:  events_size,
	}
}

// 将指定 json 文件配置导入 manager 框架。
func ImportManagerConfig(config_path string) error {
	var manager_config managerConfig
	err := ImportConfig(&manager_config, config_path)
	if err != nil {
		return err
	}

	err = manager.CreateManager(manager.GetManagerInfo())
	if err != nil {
		return err
	}

	return nil
}

// 将 manager 框架配置导出至指定 json 文件。
func ExportManagerConfig(config_path string) error {
	err := ExportConfig(getManagerConfig(), config_path)
	if err != nil {
		return err
	}

	return nil
}
