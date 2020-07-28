// package zap handles creating zap logger
package zap

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sergionunezgo/goservice/internal/logger"
	"go.uber.org/zap"
)

// RegisterLog initializes zap as the main logger for the app.
// Defer call to CloseLog after error check.
func RegisterLog() error {
	zLogger, err := initLog()
	if err != nil {
		return errors.Wrap(err, "initilize zap logger")
	}
	sugar := zLogger.Sugar()
	sugar.Info("zap logger initilized")

	logger.SetLogger(sugar)
	return nil
}

// CloseLog performs all necessary clean-up for zap logger.
func CloseLog() {
	if z, ok := logger.Log.(*zap.SugaredLogger); ok {
		_ = z.Sync()
	}
}

// initLog create logger
func initLog() (zap.Logger, error) {
	rawJSON := []byte(`{
		"level": "info",
		"Development": true,
		"DisableCaller": false,
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"timeKey":        	"ts",
			"levelKey":       	"level",
			"messageKey":     	"msg",
			"nameKey":        	"name",
			"stacktraceKey":  	"stacktrace",
			"callerKey":      	"caller",
			"lineEnding":     	"\n",
			"timeEncoder":     	"rfc3339",
			"levelEncoder":    	"lowercaseLevel",
			"durationEncoder": 	"stringDuration",
			"callerEncoder":   	"shortCaller"
		}
	}`)

	var cfg zap.Config
	var zLogger *zap.Logger
	//standard configuration
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return *zLogger, errors.Wrap(err, "unmarshal config")
	}
	//customize it from configuration file
	// err := customizeLogFromConfig(&cfg, lc)
	// if err != nil {
	// 	return *zLogger, errors.Wrap(err, "cfg.Build()")
	// }
	zLogger, err := cfg.Build()
	if err != nil {
		return *zLogger, errors.Wrap(err, "cfg.Build")
	}

	zLogger.Debug("logger construction succeeded")
	return *zLogger, nil
}
