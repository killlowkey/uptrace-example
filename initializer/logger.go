package initializer

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"uptrace-example/global"
)

// InitLogger https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelzap
// https://opentelemetry.io/ecosystem/registry
func InitLogger() {
	// 全局日志处理
	logger := otelzap.New(zap.NewExample())
	otelzap.ReplaceGlobals(logger)
	global.Log = logger
	global.Sugar = logger.Sugar()
}
