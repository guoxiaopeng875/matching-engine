package log

import (
	msLog "github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var Logger = msLog.QyLogger

// 处理 docker 中 log 到 /var/log/qy 中
func Init() {

	switch qyenv.GetDockerEnv() {
	case "", "0":
	default:
		logNameSuffix := "matching_engine"
		dockerLogDir := filepath.Join("/var/log/qy", qyenv.GetProductName()+logNameSuffix)
		if err := os.MkdirAll(dockerLogDir, 0755); err != nil {
			panic(err)
		}
		podName := os.Getenv("HOSTNAME")
		if podName == "" {
			panic("can't get HOSTNAME from env")
		}
		msLog.InitLoggerWithCaller(zapcore.DebugLevel, dockerLogDir, podName+".log", true)
	}
}
