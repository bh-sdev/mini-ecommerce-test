package logger

import (
	"github.com/superdev/ecommerce/gateway/internal/config"
	"go.uber.org/zap"
)

func NewZapLogger(
	conf *config.Config,
) (zap.Logger, error) {
	var logger = zap.NewExample()
	return *logger, nil
}
