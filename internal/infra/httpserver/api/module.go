package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/usunil0/go-dsp/internal/core/usecase"
	"github.com/usunil0/go-dsp/internal/infra/config/envcfg"
)

func provideEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/health", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	return r
}

func latencyMiddleware(log zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		d := time.Since(start)
		log.Info().
			Str("path", c.FullPath()).
			Int("status", c.Writer.Status()).
			Dur("dur", d).
			Msg("http_request")
	}
}

func provideServer(cfg envcfg.Config, r *gin.Engine) *http.Server {
	return &http.Server{
		Addr:              cfg.Addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func run(
	lc fx.Lifecycle,
	srv *http.Server,
	r *gin.Engine,
	svc usecase.BidService,
	cfg envcfg.Config,
	log zerolog.Logger,
) {
	r.Use(latencyMiddleware(log))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			mountBid(r, svc, cfg, log)
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return srv.Shutdown(ctx2)
		},
	})
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(provideEngine),
		fx.Provide(provideServer),
		fx.Invoke(run),
	)
}
