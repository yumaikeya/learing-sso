package model

import (
	"angya-backend/pkg/utils"
	"time"

	"github.com/pkg/errors"
)

type Spot struct {
	Name      string
	CreatedAt time.Time
}

func NewSpot(name *string) (Spot, error) {
	if name == nil {
		return Spot{}, errors.New("ErrNameIsRequired")
	}

	return Spot{
		Name:      *name,
		CreatedAt: utils.GetNow(),
	}, nil
}
