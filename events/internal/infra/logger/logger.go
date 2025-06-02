package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"
)

var Logger slog.Logger

func SetupLogger() {
	otelHandler := otelslog.NewHandler("events-service")
	consoleHandler := slog.NewJSONHandler(os.Stdout, nil)

	multiHandler := slogmulti.Fanout(consoleHandler, otelHandler)
	Logger := slog.New(multiHandler)
	slog.SetDefault(Logger)
}
