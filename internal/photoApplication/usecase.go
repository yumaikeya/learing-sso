package photoApplication

import (
	"angya-backend/domain/model"
	"angya-backend/pkg/databases"
	"angya-backend/pkg/utils"
	"context"
	"encoding/json"
	"time"
)

type (
	Usecase struct{}

	command struct {
		Id    *string
		Src   *string
		PoiId *string
		Spot  *string
	}

	DTO struct {
		Id        string    `json:"id"`
		PoiId     *string   `json:"poiId"`
		Src       string    `json:"src"`
		Spot      string    `json:"spot"`
		CreatedAt time.Time `json:"createdAt"`
	}

	dbModel struct {
		Id        string `gorm:"primaryKey;type:varchar(36)"` // uuidが36byte
		PoiId     *string
		Src       string `gorm:"not null"`
		Spot      string `gorm:"not null"`
		CreatedAt int64  `gorm:"not null"`
	}
)

// This function returns a pointer to Usecase.
func NewUsecase() *Usecase {
	return &Usecase{}
}

func (usecase *Usecase) Register(ctx context.Context, b []byte) (dto DTO, err error) {
	cmd := command{}
	if err = json.Unmarshal(b, &cmd); err != nil {
		return
	}

	photo, err := model.NewPhoto(cmd.Src, cmd.Spot)
	if err != nil {
		return dto, err
	}

	db := databases.NewLocalPostgres()
	if res := db.Debug().Table("photos").Save(&dbModel{Id: photo.Id, Src: photo.Src, Spot: photo.Spot, CreatedAt: photo.CreatedAt.Unix()}); res.Error != nil {
		return dto, res.Error
	}
	if d, _ := db.DB(); d != nil {
		defer d.Close()
	}

	utils.MarshalAndInsert(photo, &dto)

	return
}

func (usecase *Usecase) Update(ctx context.Context, id string, b []byte) (dto DTO, err error) {
	cmd := command{}
	if err = json.Unmarshal(b, &cmd); err != nil {
		return
	}
	cmd.Id = &id

	db := databases.NewLocalPostgres()
	dbPhoto, photo := dbModel{}, model.Photo{}
	if res := db.Table("photos").Where("id = ?", id).First(&dbPhoto); res.Error != nil {
		return dto, res.Error
	}
	utils.MarshalAndInsert(dbPhoto, &photo)

	photo.UpdateNewPhoto(cmd.PoiId, cmd.Src, cmd.Spot)

	if res := db.Debug().Table("photos").Save(&dbModel{Id: photo.Id, PoiId: photo.PoiId, Src: photo.Src, Spot: photo.Spot}); res.Error != nil {
		return dto, res.Error
	}
	if d, _ := db.DB(); d != nil {
		defer d.Close()
	}

	utils.MarshalAndInsert(photo, &dto)

	return
}

func (usecase *Usecase) List(ctx context.Context) (dtos []DTO, err error) {
	dbPhotos := []dbModel{}

	db := databases.NewLocalPostgres()
	if res := db.Debug().Table("photos").Where("poi_id IS NULL").Find(&dbPhotos); res.Error != nil {
		return dtos, res.Error
	}
	spots := func() (s []model.Photo) {
		for i := range dbPhotos {
			s = append(s, model.Photo{Id: dbPhotos[i].Id, Src: dbPhotos[i].Src, Spot: dbPhotos[i].Spot, CreatedAt: time.Unix(dbPhotos[i].CreatedAt, 0)})
		}
		return
	}()
	if d, _ := db.DB(); d != nil {
		defer d.Close()
	}

	utils.MarshalAndInsert(spots, &dtos)
	return
}
