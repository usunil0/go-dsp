package memory

import (
	"context"

	"github.com/usunil0/go-dsp/internal/core/port"
)

type budgetStore struct{}

func (budgetStore) Reserve(ctx context.Context, campaignID int, perImpMicros, limitMicros int64) (bool, error) {
	return true, nil
}

func provideBudgetStore() port.BudgetStore { return budgetStore{} }
