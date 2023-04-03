package migrations

import (
	"tugas_golang_alfi_ahsani/entity"

	"gorm.io/gorm"
)

func Migration(connection *gorm.DB) error {
	return connection.AutoMigrate(
		&entity.Books{},
	)
}
