package repository

import (
	"fmt"

	"github.com/ZSLTChenXiYin/MyGO/configure"
	"github.com/ZSLTChenXiYin/MyGO/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Init(conf configure.Configuration, logger *logger.ZapGormLogger, tables []any, back func(db *gorm.DB) error) error {
	// 创建 Gorm 配置
	gorm_config := &gorm.Config{
		Logger: logger,
	}

	database_conf := conf.Database()

	database_driver, input_driver := database_conf.Driver()
	database_dsn := database_conf.DSN()

	// 创建数据库连接
	var err error
	switch database_driver {
	case configure.DATABASE_CONFIG_DRIVER_SQLITE:
		r.database, err = gorm.Open(sqlite.Open(database_dsn), gorm_config)
	case configure.DATABASE_CONFIG_DRIVER_MYSQL:
		r.database, err = gorm.Open(mysql.Open(database_dsn), gorm_config)
	case configure.DATABASE_CONFIG_DRIVER_POSTGRESQL:
		r.database, err = gorm.Open(postgres.Open(database_dsn), gorm_config)
	default:
		return fmt.Errorf("无效数据库驱动: %s", input_driver)
	}
	if err != nil {
		return fmt.Errorf("数据库连接错误: %v", err)
	}

	err = r.database.AutoMigrate(tables...)
	if err != nil {
		return fmt.Errorf("自动迁移错误: %v", err)
	}

	err = back(r.database)
	if err != nil {
		return fmt.Errorf("回调错误: %v", err)
	}

	return nil
}

func (r *Repository) DB() *gorm.DB {
	return r.database
}
