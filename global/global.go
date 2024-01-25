package global

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Log   *otelzap.Logger
	Sugar *otelzap.SugaredLogger
)
