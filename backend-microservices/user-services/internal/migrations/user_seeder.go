package migrations

import (
	"errors"
	"log"
	"time"

	"user-services/internal/models"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	var count int64

	if count > 0 {
		log.Println("ℹ️  Users table already seeded, skipping...")
		return nil
	}

	users := []models.User{
		{
			FullName:    "Admin One",
			Username:    "admin1",
			Email:       "admin1@example.com",
			AvatarUrl:   "https://i.pravatar.cc/150?img=1",
			Password:    "password123",
			Role:        "admin",
			IsConfirmed: true,
			MemberSince: time.Now(),
		},
		{
			FullName:    "Admin Two",
			Username:    "admin2",
			Email:       "admin2@example.com",
			AvatarUrl:   "https://i.pravatar.cc/150?img=2",
			Password:    "password123",
			Role:        "admin",
			IsConfirmed: true,
			MemberSince: time.Now(),
		},
		{
			FullName:    "User One",
			Username:    "user1",
			Email:       "user1@example.com",
			AvatarUrl:   "https://i.pravatar.cc/150?img=3",
			Password:    "password123",
			Role:        "user",
			IsConfirmed: true,
			MemberSince: time.Now(),
		},
		{
			FullName:    "User Two",
			Username:    "user2",
			Email:       "user2@example.com",
			AvatarUrl:   "https://i.pravatar.cc/150?img=4",
			Password:    "password123",
			Role:        "user",
			IsConfirmed: false,
			MemberSince: time.Now(),
		},
		{
			FullName:    "User Three",
			Username:    "user3",
			Email:       "user3@example.com",
			AvatarUrl:   "https://i.pravatar.cc/150?img=5",
			Password:    "password123",
			Role:        "user",
			IsConfirmed: true,
			MemberSince: time.Now(),
		},
	}

	for _, data := range users {
		var user models.User
		err := db.Where(&models.User{Username: data.Username}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "username = ?", data.Username).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	log.Println("ℹ️  Users table seeded successfully...")

	return nil
}
