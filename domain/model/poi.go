package model

import (
	"angya-backend/pkg/utils"
	"time"

	"github.com/pkg/errors"
)

type Poi struct {
	Id        string
	PhotoId   string
	Photo     Photo
	Latitude  float64
	Longitude float64
	Comment   *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPoi(photoId *string, latitude, longitude *float64) (Poi, error) {
	if photoId == nil {
		return Poi{}, errors.New("ErrPhotoIdIsRequired")
	}

	if latitude == nil {
		return Poi{}, errors.New("ErrLatitudeIsRequired")
	}

	if longitude == nil {
		return Poi{}, errors.New("ErrLongitudeIsRequired")
	}

	return Poi{
		Id:        utils.GenId(),
		PhotoId:   *photoId,
		Latitude:  *latitude,
		Longitude: *longitude,
		Comment:   nil,
		CreatedAt: utils.GetNow(),
		UpdatedAt: utils.GetNow(),
	}, nil
}
