package configure

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	lock sync.Mutex

	conf Configuration

	viper *viper.Viper
}

func NewViperConfig(conf Configuration) *ViperConfig {
	vc := &ViperConfig{
		conf:  conf,
		viper: viper.New(),
	}

	return vc
}

func (vc *ViperConfig) Init(name string, back func(v *viper.Viper) error) error {
	vc.viper.SetConfigFile(name + ".conf.yaml")

	// 读取配置文件
	if err := vc.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("viper读取配置错误: %v", err)
	}

	// 监听配置文件变化
	vc.viper.WatchConfig()
	vc.viper.OnConfigChange(func(in fsnotify.Event) {
		var err error
		// 加载配置
		err = vc.load()
		if err != nil {
			panic(fmt.Errorf("viper加载配置错误: %v", err))
		}
	})

	if back != nil {
		if err := back(vc.viper); err != nil {
			return err
		}
	}

	// 首次加载配置
	err := vc.load()
	if err != nil {
		return fmt.Errorf("viper加载配置错误: %v", err)
	}

	return nil
}

func (vc *ViperConfig) load() error {
	// 并发安全
	vc.lock.Lock()
	defer vc.lock.Unlock()

	err := vc.viper.Unmarshal(vc.conf)
	return err
}

func (vc *ViperConfig) Viper() *viper.Viper {
	return vc.viper
}
