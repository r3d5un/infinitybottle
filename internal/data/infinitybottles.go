package data

import (
	"time"
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
