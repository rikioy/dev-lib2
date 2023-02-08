package slog

import (
	senv "github.com/rikioy/dev-lib/senv"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func init() {
	klog.SetLogger(kitexlogrus.NewLogger())
}

func InitKLog(level string) {
	switch level {
	case senv.ENV_DEV:
		klog.SetLevel(klog.LevelDebug)
	}
}
