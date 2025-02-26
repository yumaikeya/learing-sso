package model

import (
	"angya-backend/pkg/utils"
	"time"

	"github.com/pkg/errors"
)

type Photo struct {
	Id        string
	PoiId     *string
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
		PoiId:     nil,
		Src:       *src,
		Spot:      *spot,
		CreatedAt: utils.GetNow(),
	}, nil
}

func (photo *Photo) UpdateNewPhoto(poiId, src, spot *string) {
	if poiId != nil {
		photo.PoiId = poiId
	}

	if src != nil {
		photo.Src = *src
	}

	if spot != nil {
		photo.Spot = *spot
	}
}
