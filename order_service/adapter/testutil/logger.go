package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func SetupLogger(t *testing.T) *zap.SugaredLogger {
	t.Helper()

	logger, err := zap.NewProduction()
	defer logger.Sync()
	assert.NoError(t, err)
	return logger.Sugar()
}
