package misc

import (
	"os"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Config Application config parameters
type Config struct {
	Mongo MongoConfig
	HTTP  HTTPConfig
	Log   LogConfig
}

// MongoConfig is for MongoDB
type MongoConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
	PoolSize *uint64
}

// HTTPConfig HTTP config parameters
type HTTPConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Domain       string
	CookieExpiry int
}

// LogConfig Logging configuration
type LogConfig struct {
	Level      string
	Format     string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

// InitLogging Initialize logging framework
func InitLogging(lc *LogConfig) {
	switch lc.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		fallthrough
	case "text":
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	}

	switch lc.Level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		fallthrough
	case "info":
		log.SetLevel(log.InfoLevel)
	}

	if lc.Filename == "" {
		log.SetOutput(os.Stdout)
	} else {
		l := &lumberjack.Logger{
			Filename:   lc.Filename,
			MaxSize:    lc.MaxSize,
			MaxAge:     lc.MaxAge,
			MaxBackups: lc.MaxBackups,
			LocalTime:  lc.LocalTime,
			Compress:   lc.Compress,
		}

		log.SetOutput(l)
	}
}
