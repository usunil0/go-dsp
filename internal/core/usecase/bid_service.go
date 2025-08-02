// internal/core/usecase/bid_service.go
package usecase

import (
	"context"
	"time"

	"github.com/usunil0/go-dsp/internal/core/entity"
	"github.com/usunil0/go-dsp/internal/core/port"
)

type BidInput struct {
	ReqID   string
	ImpID   string
	Country string
	Device  string
}

type BidResult struct {
	CampaignID int
	CreativeID string
	PriceUSD   float64
}

type BidService struct {
	Repo   port.CampaignRepo
	Budget port.BudgetStore
}

func NewBidService(repo port.CampaignRepo, budget port.BudgetStore) BidService {
	return BidService{Repo: repo, Budget: budget}
}

func (s BidService) Bid(ctx context.Context, in BidInput) (BidResult, bool, error) {

	if in.ReqID == "" || in.ImpID == "" {
		return BidResult{}, false, nil
	}


	camps, err := s.Repo.Active(ctx)
	if err != nil || len(camps) == 0 {
		return BidResult{}, false, err
	}

	reqCtx := entity.RequestContext{
		Country:   in.Country,
		Timestamp:   time.Now(),
		ContentCats: nil,
	}

	for _, c := range camps {
		if !c.CanServe() {
			continue
		}
		if !c.Matches(reqCtx) {
			continue
		}
		perImp := c.Price.PerImpMicros()
		ok, err := s.Budget.Reserve(ctx, c.ID, perImp, c.BudgetMicros)
		if err != nil || !ok {
			continue
		}
		return BidResult{
			CampaignID: c.ID,
			CreativeID: c.CreativeID,
			PriceUSD:   c.PriceUSD(),
		}, true, nil
	}

	return BidResult{}, false, nil
}
