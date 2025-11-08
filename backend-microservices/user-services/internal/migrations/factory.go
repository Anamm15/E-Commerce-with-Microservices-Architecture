package migrations

import (
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := SeedUsers(db); err != nil {
		return err
	}

	return nil
}
