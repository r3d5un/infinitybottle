package data

import (
	"context"
	"database/sql"
	"errors"
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
	Version          int64     `json:"version"`
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&contribution.ID, &contribution.AddedAt)
}

func (m ContributionModel) Get(id int64) (*Contribution, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, infinitybottle_id, added_at, amount, brand_name, tags
        FROM contributions
        WHERE id = $1`

	var contribution Contribution

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&contribution.ID,
		&contribution.InfinityBottleID,
		&contribution.AddedAt,
		&contribution.Amount,
		&contribution.BrandName,
		pq.Array(&contribution.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &contribution, nil
}

func (m ContributionModel) Update(contribution *Contribution) error {
	query := `
        UPDATE contributions
        SET infinitybottle_id = $1, brand_name = $2, tags = $3, version = version + 1
        WHERE id = $4
        RETURNING version`

	args := []any{
		contribution.InfinityBottleID,
		contribution.BrandName,
		pq.Array(contribution.Tags),
		contribution.ID,
		contribution.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&contribution.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}
	return nil
}

func (m ContributionModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM contributions
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
