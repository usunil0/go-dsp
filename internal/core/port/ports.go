package port

import (
	"context"

	"github.com/usunil0/go-dsp/internal/core/entity"
)

type CampaignRepo interface {
	Active(ctx context.Context) ([]entity.Campaign, error)
}

type BudgetStore interface {
	Reserve(ctx context.Context, campaignID int, perImpMicros, limitMicros int64) (bool, error)
}
