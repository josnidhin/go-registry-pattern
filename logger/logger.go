/**
 * @author Jose Nidhin
 */
package logger

import (
	"go.uber.org/zap"
)

var DefaultLogger *zap.Logger

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("Unexpected error while creating logger")
	}

	DefaultLogger = logger
}
