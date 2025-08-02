package main

import (
	"go.uber.org/fx"

	"github.com/usunil0/go-dsp/internal/core/usecase"
	"github.com/usunil0/go-dsp/internal/infra/config/envcfg"
	"github.com/usunil0/go-dsp/internal/infra/httpserver/api"
	"github.com/usunil0/go-dsp/internal/infra/log/zerologx"
	"github.com/usunil0/go-dsp/internal/infra/memory"
)

func main() {
	fx.New(
		zerologx.Module(),
		envcfg.Module(),
		memory.Module(),
		usecase.Module(),
		api.Module(),
	).Run()
}
