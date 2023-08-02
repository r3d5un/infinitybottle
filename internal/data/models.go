package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	InfinityBottles InfinityBottleModel
	Contributions   ContributionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		InfinityBottles: InfinityBottleModel{DB: db},
		Contributions:   ContributionModel{DB: db},
	}
}
