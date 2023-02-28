package logs

import (
	"github.com/hashicorp/go-hclog"
)

type Logger hclog.Logger

func NewLogger() hclog.Logger {
	Logger := hclog.New(&hclog.LoggerOptions{
		Name:  "NsInformer",
		Level: hclog.LevelFromString("DEBUG"),
	})

	return Logger
}
