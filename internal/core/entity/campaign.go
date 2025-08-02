// internal/core/entity/campaign.go
package entity

import (
	"time"

	"github.com/usunil0/go-dsp/internal/core/valueobject"
)

// Campaign is your core ad campaign entity, extended with
// pricing, budget and full targeting dimensions.
type Campaign struct {
	ID           int
	Price        valueobject.CPMMicros
	Margin       float64 // e.g. 0.1 for 10% off CPM
	BudgetMicros int64
	CreativeID   string

	// Geographic
	Countries   []string
	Cities      []string
	ISPs        []string
	MinConnKbps int

	// Demographic
	AgeMin, AgeMax int
	Genders        []string
	Languages      []string

	// Device
	Manufacturers []string
	OSs           []string
	Browsers      []string

	// Temporal
	DaysOfWeek []time.Weekday
	HoursOfDay []int

	// Content Categories
	WhitelistCats []string
	BlacklistCats []string
}

// CanServe returns true if the campaign still has creative & budget.
func (c Campaign) CanServe() bool {
	return c.CreativeID != "" && c.BudgetMicros > 0
}

// PriceUSD computes this campaign's bid price in dollars.
func (c Campaign) PriceUSD() float64 {
	return c.Price.USD() * (1.0 - c.Margin)
}

// RequestContext holds all info needed for matching.
type RequestContext struct {
	Country      string
	City         string
	ISP          string
	ConnKbps     int
	Age          int
	Gender       string
	Language     string
	Manufacturer string
	OS           string
	Browser      string
	Timestamp    time.Time
	ContentCats  []string
}

// Matches returns true if the request meets all the campaign's
// geographic, demographic, device, temporal and content rules.
func (c Campaign) Matches(ctx RequestContext) bool {
	// Geographic
	if len(c.Countries) > 0 && !contains(c.Countries, ctx.Country) {
		return false
	}
	if len(c.Cities) > 0 && !contains(c.Cities, ctx.City) {
		return false
	}
	if len(c.ISPs) > 0 && !contains(c.ISPs, ctx.ISP) {
		return false
	}
	if c.MinConnKbps > 0 && ctx.ConnKbps < c.MinConnKbps {
		return false
	}

	// Demographic
	if (c.AgeMin > 0 || c.AgeMax > 0) &&
		(ctx.Age < c.AgeMin || ctx.Age > c.AgeMax) {
		return false
	}
	if len(c.Genders) > 0 && !contains(c.Genders, ctx.Gender) {
		return false
	}
	if len(c.Languages) > 0 && !contains(c.Languages, ctx.Language) {
		return false
	}

	// Device
	if len(c.Manufacturers) > 0 && !contains(c.Manufacturers, ctx.Manufacturer) {
		return false
	}
	if len(c.OSs) > 0 && !contains(c.OSs, ctx.OS) {
		return false
	}
	if len(c.Browsers) > 0 && !contains(c.Browsers, ctx.Browser) {
		return false
	}

	// Temporal
	if len(c.DaysOfWeek) > 0 && !containsWeekday(c.DaysOfWeek, ctx.Timestamp.Weekday()) {
		return false
	}
	if len(c.HoursOfDay) > 0 && !containsInt(c.HoursOfDay, ctx.Timestamp.Hour()) {
		return false
	}

	// Content
	for _, cat := range ctx.ContentCats {
		if contains(c.BlacklistCats, cat) {
			return false
		}
	}
	if len(c.WhitelistCats) > 0 {
		ok := false
		for _, cat := range ctx.ContentCats {
			if contains(c.WhitelistCats, cat) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

// ---------- helpers ----------
func contains(list []string, val string) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}

func containsInt(list []int, val int) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}

func containsWeekday(list []time.Weekday, val time.Weekday) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}
