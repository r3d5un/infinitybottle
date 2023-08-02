package data

import (
	"database/sql"
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

type InfinityBottleModel struct {
	DB *sql.DB
}

func (m InfinityBottleModel) Insert(infinityBottle *InfinityBottle) error {
	query := `
        INSERT INTO infinitybottles (bottle_name, empty_start)
        VALUES ($1, $2)
        RETURNING id, created_at`

	args := []any{infinityBottle.BottleName, infinityBottle.EmptyStart}
	return m.DB.QueryRow(query, args...).Scan(&infinityBottle.ID, &infinityBottle.CreatedAt)
}

func (m InfinityBottleModel) Get(id int64) (*InfinityBottle, error) {
	return nil, nil
}

func (m InfinityBottleModel) Update(infinityBottle *InfinityBottle) error {
	return nil
}

func (m InfinityBottleModel) Delete(id int64) error {
	return nil
}
