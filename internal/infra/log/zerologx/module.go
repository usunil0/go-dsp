package zerologx

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func Provide() zerolog.Logger {
	return zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
}

func Module() fx.Option { return fx.Provide(Provide) }
