package configure

type Configuration interface {
	Server() ServerConfiguration
	Log() LogConfiguration
	Database() DatabaseConfiguration
	Email() EmailConfiguration
}

type ServerConfiguration interface {
	Debug() bool
	JwtSecret() string
}

type LogConfiguration interface {
	Date() bool
	OutLog() string
	ZapLog() string
}

type DatabaseConfiguration interface {
	Driver() (string, string)
	DSN() string
}

type EmailConfiguration interface {
	Host() string
	Port() int
	Email() string
	Password() string
}
