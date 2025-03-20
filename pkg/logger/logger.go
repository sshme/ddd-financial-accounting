package logger

import (
	"ddd-financial-accounting/pkg/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	return setupPrettySlog()
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
