package memory

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(provideCampaignRepo),
		fx.Provide(provideBudgetStore),
	)
}
