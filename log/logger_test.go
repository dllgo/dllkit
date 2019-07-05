package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	llog := Logger("applog")
	llog.Debug("hello world")
	llog.Info("hello world")
	llog.Info("hello world")
	llog.Info("hello world")
	llog.Info("hello world")
	llog.Info("hello world")
	llog.Info("Info", zap.Any("user-srv", "user-srv is start ..."))
}
