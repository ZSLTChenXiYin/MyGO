package configure

type StandardConfig struct {
	ServerConfig   ServerConfig    `mapstructure:"server"`
	LogConfig      LogConfig       `mapstructure:"log"`
	RedisConfig    *RedisConfig    `mapstructure:"redis"`
	DatabaseConfig *DatabaseConfig `mapstructure:"database"`
	EmailConfig    *EmailConfig    `mapstructure:"email"`
}

func (sc *StandardConfig) Server() ServerConfiguration {
	return &sc.ServerConfig
}

func (sc *StandardConfig) Log() LogConfiguration {
	return &sc.LogConfig
}

func (sc *StandardConfig) Database() DatabaseConfiguration {
	return sc.DatabaseConfig
}

const (
	SERVER_CONFIG_MODE_DEV  = "dev"
	SERVER_CONFIG_MODE_PROD = "prod"
)

type ServerConfig struct {
	Mode      string  `mapstructure:"mode"`
	HTTPPort  *uint   `mapstructure:"http_port"`
	TCPPort   *uint   `mapstructure:"tcp_port"`
	JWTSecret *string `mapstructure:"jwt_secret"`
}

func (sc *ServerConfig) Debug() bool {
	return sc.Mode == SERVER_CONFIG_MODE_DEV
}

func (sc *ServerConfig) JwtSecret() string {
	if sc.JWTSecret == nil {
		return ""
	}
	return *sc.JWTSecret
}

const (
	LOG_CONFIG_MODE_STANDARD = "standard"
	LOG_CONFIG_MODE_DATE     = "date" // 需要用有且只有一个`%s`占位符在`OutLog`和`ZapLog`中预留日期插入位置
)

type LogConfig struct {
	Mode       string `mapstructure:"mode"`
	OutLogPath string `mapstructure:"out_log"`
	ZapLogPath string `mapstructure:"zap_log"`
}

func (lc *LogConfig) Date() bool {
	return lc.Mode == LOG_CONFIG_MODE_DATE
}

func (lc *LogConfig) OutLog() string {
	return lc.OutLogPath
}

func (lc *LogConfig) ZapLog() string {
	return lc.ZapLogPath
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

const (
	DATABASE_CONFIG_DRIVER_SQLITE     = "sqlite"
	DATABASE_CONFIG_DRIVER_MYSQL      = "mysql"
	DATABASE_CONFIG_DRIVER_POSTGRESQL = "postgresql"
)

type DatabaseConfig struct {
	DBDriver string `mapstructure:"driver"`
	DBDSN    string `mapstructure:"dsn"`
}

func (dc *DatabaseConfig) Driver() (string, string) {
	input_driver := dc.DBDriver

	var database_driver string
	switch input_driver {
	case DATABASE_CONFIG_DRIVER_SQLITE:
		database_driver = DATABASE_CONFIG_DRIVER_SQLITE
	case DATABASE_CONFIG_DRIVER_MYSQL:
		database_driver = DATABASE_CONFIG_DRIVER_MYSQL
	case DATABASE_CONFIG_DRIVER_POSTGRESQL:
		database_driver = DATABASE_CONFIG_DRIVER_POSTGRESQL
	default:
		return "", input_driver
	}

	return database_driver, input_driver
}

func (dc *DatabaseConfig) DSN() string {
	return dc.DBDSN
}

type EmailConfig struct {
	Template        string `mapstructure:"template"`
	EmailServerHost string `mapstructure:"host"`
	EmailServerPort uint   `mapstructure:"port"`
	SenderEmail     string `mapstructure:"email"`
	SenderPassword  string `mapstructure:"password"`
}

func (ec *EmailConfig) Host() string {
	return ec.EmailServerHost
}

func (ec *EmailConfig) Port() int {
	return int(ec.EmailServerPort)
}

func (ec *EmailConfig) Email() string {
	return ec.SenderEmail
}

func (ec *EmailConfig) Password() string {
	return ec.SenderPassword
}
