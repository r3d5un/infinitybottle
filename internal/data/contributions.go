package data

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"infinitybottle.islandwind.me/internal/validator"
)

type Contribution struct {
	ID               int64     `json:"id"`
	InfinityBottleID int64     `json:"infinityBottleID"`
	AddedAt          time.Time `json:"addedAt"`
	Amount           int64     `json:"amount"`
	BrandName        string    `json:"brandName"`
	Tags             []string  `json:"tags,omitempty"`
}

func ValidateContribution(v *validator.Validator, contributionPost *Contribution) {
	v.Check(contributionPost.InfinityBottleID != 0, "infinityBottleID", "must be provided")
	v.Check(contributionPost.Amount != 0, "amount", "must be provided")
	v.Check(contributionPost.Amount > 0, "amount", "must be greater than 0")
	v.Check(contributionPost.BrandName != "", "brandName", "must be provided")
	v.Check(
		len(contributionPost.BrandName) <= 255,
		"brandName",
		"must not be more than 255 bytes long",
	)
	if contributionPost.Tags != nil {
		v.Check(
			validator.Unique(contributionPost.Tags),
			"tags",
			"must not contain duplicate values",
		)
	}
}

type ContributionModel struct {
	DB *sql.DB
}

func (m ContributionModel) Insert(contribution *Contribution) error {
	query := `
        INSERT INTO contributions (infinitybottle_id, brand_name, tags)
        VALUES ($1, $2, $3)
        RETURNING id, added_at`

	args := []any{
		contribution.InfinityBottleID,
		contribution.BrandName,
		pq.Array(contribution.Tags),
	}
	return m.DB.QueryRow(query, args...).Scan(&contribution.ID, &contribution.AddedAt)
}

func (m ContributionModel) Get(id int64) (*Contribution, error) {
	return nil, nil
}

func (m ContributionModel) Update(contribution *Contribution) error {
	return nil
}

func (m ContributionModel) Delete(id int64) error {
	return nil
}
