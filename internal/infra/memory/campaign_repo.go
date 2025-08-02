package memory

import (
	"context"

	"github.com/usunil0/go-dsp/internal/core/entity"
	"github.com/usunil0/go-dsp/internal/core/port"
	"github.com/usunil0/go-dsp/internal/core/valueobject"
)

type campaignRepo struct{}

func (campaignRepo) Active(ctx context.Context) ([]entity.Campaign, error) {
	return []entity.Campaign{
		{ID: 1, Price: valueobject.CPMMicros(2_000_000), BudgetMicros: 10_000_000, CreativeID: "cre-1"},
	}, nil
}

func provideCampaignRepo() port.CampaignRepo { return campaignRepo{} }
