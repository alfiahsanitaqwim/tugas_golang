package migrations

import (
	"gorm.io/gorm"
)

func Migration(connection *gorm.DB) error {
	//rollback.Rollback(connection)
	return connection.AutoMigrate(
	// &entityBlog.BlogsLabel{},
	)
}
