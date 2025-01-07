package main

import (
	// "fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var level zap.AtomicLevel

func initLogger() {
	level = zap.NewAtomicLevel() // 创建一个动态控制的日志级别
	level.SetLevel(zap.InfoLevel)
	config := zap.Config{
		Level:            level, // 使用动态控制的日志级别
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func main() {
	initLogger()
	defer logger.Sync()

	// 启动HTTP服务器以监听日志级别修改请求
	http.Handle("/set-log-level", level)
	go http.ListenAndServe(":8080", nil)

	logger.Info("This is an info message")
	logger.Debug("This debug message will not be printed")

	go func() {
		t := time.NewTicker(time.Second * 3)
		for range t.C {
			logger.Info("info")
			logger.Error("error")
			logger.Debug("debug")
		}
	}()

	// 模拟程序运行
	select {}
}

// func setLogLevel(w http.ResponseWriter, r *http.Request) {
// 	levelStr := r.URL.Query().Get("level")
// 	if levelStr == "" {
// 		http.Error(w, "level query parameter is required", http.StatusBadRequest)
// 		return
// 	}

// 	var zapLevel zapcore.Level
// 	if err := zapLevel.UnmarshalText([]byte(levelStr)); err != nil {
// 		http.Error(w, fmt.Sprintf("invalid log level: %v", err), http.StatusBadRequest)
// 		return
// 	}

// 	level.SetLevel(zapLevel) // 动态修改日志级别
// 	w.Write([]byte(fmt.Sprintf("Log level updated to %s", zapLevel)))
// }
