// internal/infra/httpserver/api/bid_handler.go
package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/usunil0/go-dsp/internal/core/dto"
	"github.com/usunil0/go-dsp/internal/core/usecase"
	"github.com/usunil0/go-dsp/internal/infra/config/envcfg"
)

func mountBid(r *gin.Engine, svc usecase.BidService, cfg envcfg.Config, log zerolog.Logger) {
	r.POST("/bid", func(c *gin.Context) {
		var br dto.BidRequest
		if err := c.ShouldBindJSON(&br); err != nil {
			log.Error().Err(err).Msg("bind bid request")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if br.ID == "" || len(br.Imp) == 0 {
			log.Warn().Msg("missing id or imp")
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing id or imp"})
			return
		}

		var bids []dto.Bid
		for _, imp := range br.Imp {
			// simple imp validation
			if imp.Banner == nil && imp.Video == nil && imp.Native == nil {
				continue
			}
			in := usecase.BidInput{
				ReqID: br.ID,
				ImpID: imp.ID,
			}
			if br.Device.Geo != nil {
				in.Country = br.Device.Geo.Country
			}
			if br.Device.DeviceType != 0 {
				in.Device = strconv.Itoa(int(br.Device.DeviceType))
			}

			start := time.Now()
			res, ok, err := svc.Bid(c.Request.Context(), in)
			log.Debug().
				Str("req", in.ReqID).
				Str("imp", in.ImpID).
				Dur("hot_ms", time.Since(start)).
				Msg("bid_hotpath")
			if err != nil {
				log.Error().Err(err).Msg("service error")
				continue
			}
			if !ok {
				continue
			}

			bidID := uuid.NewString()
			nurl := fmt.Sprintf("%s/win?impid=%s&price=%.2f", cfg.Host, in.ImpID, res.PriceUSD)
			bids = append(bids, dto.Bid{
				ID:    bidID,
				ImpID: in.ImpID,
				Price: res.PriceUSD,
				AdID:  res.CreativeID,
				AdM:   "<html>banner</html>",
				NURL:  nurl,
			})
		}

		if len(bids) == 0 {
			c.Status(http.StatusNoContent)
			return
		}

		resp := dto.BidResponse{
			ID: br.ID,
			SeatBid: []dto.SeatBid{{
				Seat: cfg.Seat,
				Bid:  bids,
			}},
		}
		c.JSON(http.StatusOK, resp)
	})
}
