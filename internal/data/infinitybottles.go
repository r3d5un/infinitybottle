package data

import (
	"time"

	"infinitybottle.islandwind.me/internal/validator"
)

type InfinityBottle struct {
	ID                    int64          `json:"id"`
	BottleName            string         `json:"bottleName"`
	NumberOfContributions int64          `json:"numberOfContributions"`
	EmptyStart            bool           `json:"emptyStart"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	Contributions         []Contribution `json:"contributions,omitempty"`
}

func ValidateInfinityBottle(v *validator.Validator, infinityBottle *InfinityBottle) {
	v.Check(infinityBottle.BottleName != "", "bottleName", "must be provided")
	v.Check(
		len(infinityBottle.BottleName) <= 255,
		"bottleName",
		"must not be more than 255 bytes long",
	)
}
