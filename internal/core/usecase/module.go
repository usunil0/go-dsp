package usecase

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			func() BidPolicy { return FixedMargin{Margin: 0.0} },
			func() Spec { return Always{} },
			NewBidService,
		),
	)
}
