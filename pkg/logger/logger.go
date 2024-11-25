package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func New() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

				if a.Key == "error" {
					v, ok := a.Value.Any().(error)
					if !ok {
						return a
					}
					return tint.Err(v)
				}

				return a
			},
		}),
	)

	slog.SetDefault(log)

	return log
}
