package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age:%d", "TipGo", 40) // like as fmt.Printf(format, a)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 40))

	// 2.
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// //Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	// 3.
	encoder := getEncoderLog()
	writerSync := getWriterSync()
	core := zapcore.NewCore(encoder, writerSync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// Format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//1753451126.856311 -> 2025-07-25T20:45:26.855+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> timestamp
	encoderConfig.TimeKey = "timestamp"
	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller":"cli/main.log.go:21
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//

func getWriterSync() zapcore.WriteSyncer {
	err := os.MkdirAll("./log", 0755)
	if err != nil {
		panic(err)
	}

	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
