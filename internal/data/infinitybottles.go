package data

import (
	"database/sql"
	"errors"
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
	Version               int64          `json:"version"`
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
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, bottle_name, number_of_contributions, empty_start, created_at, updated_at
        FROM infinitybottles
        WHERE id = $1`

	var infinityBottle InfinityBottle
	err := m.DB.QueryRow(query, id).Scan(
		&infinityBottle.ID,
		&infinityBottle.BottleName,
		&infinityBottle.NumberOfContributions,
		&infinityBottle.EmptyStart,
		&infinityBottle.CreatedAt,
		&infinityBottle.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &infinityBottle, nil
}

func (m InfinityBottleModel) Update(infinityBottle *InfinityBottle) error {
	query := `
        UPDATE infinitybottles
        SET bottle_name = $1, empty_start = $2, updated_at = CURRENT_TIMESTAMP, version = version + 1
        WHERE id = $3
        RETURNING updated_at, version`

	args := []any{
		infinityBottle.BottleName,
		infinityBottle.EmptyStart,
		infinityBottle.ID,
		infinityBottle.Version,
	}

	err := m.DB.QueryRow(query, args...).Scan(&infinityBottle.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m InfinityBottleModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM infinitybottles
        WHERE id = $1`

	result, err := m.DB.Exec(query, id)
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
