// internal/infra/config/envcfg/config.go
package envcfg

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type Config struct {
	// HTTP listen address, e.g. ":8080"
	Addr string `env:"ADDR" envDefault:":8080"`

	// Seat ID to return in the OpenRTB response
	Seat string `env:"SEAT" envDefault:"123"`

	Host string `env:"HOST" envDefault:"http://localhost:8080"`

}

// Provide parses environment into Config, or returns an error if invalid
func Provide() (Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("envcfg: failed to parse config: %w", err)
	}
	return c, nil
}

func Module() fx.Option {
	return fx.Provide(Provide)
}
