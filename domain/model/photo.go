package model

import (
	"angya-backend/pkg/utils"
	"time"

	"github.com/pkg/errors"
)

type Photo struct {
	Id        string
	Src       string
	Spot      string
	CreatedAt time.Time
}

func NewPhoto(src, spot *string) (Photo, error) {
	if src == nil {
		return Photo{}, errors.New("ErrSrcIsRequired")
	}

	if spot == nil {
		return Photo{}, errors.New("ErrSpotIsRequired")
	}

	return Photo{
		Id:        utils.GenId(),
		Src:       *src,
		Spot:      *spot,
		CreatedAt: utils.GetNow(),
	}, nil
}
