package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/ZSLTChenXiYin/MyGO/configure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

func NewStandardZapConfig(conf configure.Configuration) zap.Config {
	zap_conf := zap.NewProductionConfig()

	zap_conf.EncoderConfig.TimeKey = "timestamp"
	zap_conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var out_log string
	var zap_log string

	if conf.Log().Date() {
		date := time.Now().Format("2006-01-02")
		out_log = fmt.Sprintf(conf.Log().OutLog(), date)
		zap_log = fmt.Sprintf(conf.Log().ZapLog(), date)
	} else {
		out_log = conf.Log().OutLog()
		zap_log = conf.Log().ZapLog()
	}

	zap_conf.OutputPaths = []string{out_log}
	zap_conf.ErrorOutputPaths = []string{zap_log}

	return zap_conf
}

type ZapGormLogger struct {
	std_logger *StdLogger
	zap_logger *zap.Logger
	log_level  logger.LogLevel
}

func NewZapGormLogger(std_logger *StdLogger, zap_logger *zap.Logger) *ZapGormLogger {
	return &ZapGormLogger{
		std_logger: std_logger,
		zap_logger: zap_logger,
		log_level:  logger.Info, // 默认日志级别
	}
}

// LogMode 实现 logger.Interface 的 LogMode 方法
func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	new_logger := *l
	new_logger.log_level = level
	return &new_logger
}

// Info 实现 logger.Interface 的 Info 方法
func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...any) {
	if l.log_level >= logger.Info {
		l.zap_logger.Sugar().Infof(msg, data...)
	}
}

// Warn 实现 logger.Interface 的 Warn 方法
func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...any) {
	if l.log_level >= logger.Warn {
		l.zap_logger.Sugar().Warnf(msg, data...)
	}
}

// Error 实现 logger.Interface 的 Error 方法
func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...any) {
	if l.log_level >= logger.Error {
		l.zap_logger.Sugar().Errorf(msg, data...)
	}
}

// Trace 实现 logger.Interface 的 Trace 方法
func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.log_level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.Duration("elapsed", elapsed),
		zap.String("sql", sql),
		zap.Int64("rows", rows),
	}

	if err != nil {
		l.zap_logger.Error("GORM Trace", append(fields, zap.Error(err))...)
	} else {
		l.zap_logger.Info("GORM Trace", fields...)
		l.std_logger.Debugf("GORM Trace: %s, %dms, %d rows", sql, elapsed.Milliseconds(), rows)
	}
}
