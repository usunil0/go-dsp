package usecase

import "github.com/usunil0/go-dsp/internal/core/valueobject"

type BidPolicy interface {
	PriceUSD(c valueobject.CPMMicros) float64
}

type FixedMargin struct{ Margin float64 }

func (p FixedMargin) PriceUSD(c valueobject.CPMMicros) float64 {
	return c.USD() * (1 - p.Margin)
}
