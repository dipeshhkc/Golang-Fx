package infrastructure

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger sets up logger
func NewLogger() Logger {

	config := zap.NewDevelopmentConfig()

	if os.Getenv("ENVIRONMENT") == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if os.Getenv("ENVIRONMENT") == "production" && os.Getenv("LOGOUTPUT") != "" {
		config.OutputPaths = []string{os.Getenv("LOGOUTPUT")}
	}

	logger, _ := config.Build()

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
