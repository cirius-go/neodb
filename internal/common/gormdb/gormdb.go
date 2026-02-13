package gormdb

import (
	"database/sql"
	"fmt"
	"time"

	"cirius-go/neodb/internal/common"
	"cirius-go/neodb/internal/common/logger"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// logWriter wraps common.Logger to implement gorm/logger.Interface.
type logWriter struct {
	silent bool
	inner  common.Logger
}

// Printf implements logger.Writer.
func (g *logWriter) Printf(msg string, args ...any) {
	g.inner.Printf(msg, args...)
}

func convertGormLogLevel(level common.LoggerLevel) gormlogger.LogLevel {
	switch level {
	case common.LoggerLevelWarn:
		return gormlogger.Warn
	case common.LoggerLevelError:
		return gormlogger.Error
	default:
		return gormlogger.Info
	}
}

var _ gormlogger.Writer = (*logWriter)(nil)

// NewGormLogger creates a new GormLogger instance.
func NewGormLogger(inner common.Logger) *logWriter {
	return &logWriter{
		inner: inner,
	}
}

// GormConfig represents the configuration for GORM.
type GormConfig struct {
	Silent          bool
	SlowThreshold   time.Duration
	Logger          common.Logger
	Dialector       gorm.Dialector
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func (c *GormConfig) fillDefaults() {
	if c.SlowThreshold <= 0 {
		c.SlowThreshold = 200 * time.Millisecond
	}
	if c.Logger == nil {
		c.Logger = logger.Get()
	}
	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = 10
	}
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = 100
	}
	if c.ConnMaxIdleTime <= 0 {
		c.ConnMaxIdleTime = 10 * time.Minute
	}
	if c.ConnMaxLifetime <= 0 {
		c.ConnMaxLifetime = 10 * time.Minute
	}
}

// NewGorm creates new database connection to the database server with default
// configuration.
func NewGorm(cfg GormConfig) (*gorm.DB, *sql.DB, error) {
	cfg.fillDefaults()
	lw := &logWriter{
		inner: cfg.Logger,
	}
	logLevel := convertGormLogLevel(lw.inner.LogLevel())
	if cfg.Silent {
		logLevel = gormlogger.Silent
	}
	lg := gormlogger.New(lw, gormlogger.Config{
		SlowThreshold:             cfg.SlowThreshold,
		LogLevel:                  logLevel,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false,
		ParameterizedQueries:      cfg.Logger.LogLevel() == common.LoggerLevelDebug,
	})
	db, err := gorm.Open(cfg.Dialector, &gorm.Config{
		Logger:                                   lg,
		AllowGlobalUpdate:                        false,
		CreateBatchSize:                          1000,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("cannot establish connection: %w", err)
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get generic db instance: %w", err)
	}
	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqldb.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	sqldb.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	return db, sqldb, nil
}
