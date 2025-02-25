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
		Src  string
		Spot string
	}

	DTO struct {
		Id        string    `json:"id"`
		Src       string    `json:"src"`
		Spot      string    `json:"spot"`
		CreatedAt time.Time `json:"createdAt"`
	}

	dbModel struct {
		Id        string `gorm:"primaryKey;type:varchar(36)"` // uuidが36byte
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

	photo, err := model.NewPhoto(&cmd.Src, &cmd.Spot)
	if err != nil {
		return dto, err
	}

	db := databases.NewLocalPostgres()
	if res := db.Debug().Table("photos").Save(&dbModel{Id: photo.Id, Src: photo.Src, Spot: photo.Spot, CreatedAt: photo.CreatedAt.Unix()}); res.Error != nil {
		return dto, res.Error
	}

	utils.MarshalAndInsert(photo, &dto)

	return
}

func (usecase *Usecase) List(ctx context.Context) (dtos []DTO) {

	return
}
