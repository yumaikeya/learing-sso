package databases

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresのLocal
func NewLocalPostgres() *gorm.DB {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=db user=postgres password=postgres dbname=angya port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Fail to generate lcoal postgres clinet: %s", err.Error()))
	}

	return db
}
