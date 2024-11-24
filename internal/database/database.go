package database

import (
	"log/slog"

	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	Email        string
	Username     string
	Fullname     string
	PasswordHash string
	Note         string
	CreatedAt    int64
}

type Contact struct {
	ID        string `gorm:"primaryKey"`
	Subject   string
	Message   string
	IPAddress string
	CreatedAt int64
}

type AccessLog struct {
	ID             string `gorm:"primaryKey"`
	Timestamp      int64
	URL            string
	IPAddress      string
	Referrer       string
	Platform       string
	Mobile         string
	UserAgent      string
	UserAgentBrand string
}

type URLRedirect struct {
	ID             string `gorm:"primaryKey"`
	Keyword        string
	DestinationURL string
}

type Category struct {
	ID               string `gorm:"primaryKey"`
	ParentID         string
	Name             string
	RelativePosition float64
}

type Note struct {
	ID         string `gorm:"primaryKey"`
	CategoryID string
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	AuthorID   string
	Author     User `gorm:"foreignKey:AuthorID;references:ID"`
	CreatedAt  int64
	UpdatedAt  int64
	Title      string
	Html       string
	Markdown   string
	Brief      string
	AssetsMap  string
}

func InitializeDatabase(db *gorm.DB, logger *slog.Logger) error {
	err := db.AutoMigrate(
		&User{}, &Contact{}, &AccessLog{},
		&URLRedirect{}, &Category{}, &Note{},
	)

	if err != nil {
		logger.Error("Failed to migrate database schema: ", err)
		return err
	}

	logger.Debug("Database initialized successfully.")
	return nil
}
