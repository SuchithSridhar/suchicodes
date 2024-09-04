package database

import (
	"log/slog"

	"gorm.io/gorm"
)

type Admin struct {
	Id           string `gorm:"primaryKey"`
	Email        string
	PasswordHash string
	Note         string
	CreatedAt    int64
}

type ServerLog struct {
	Id             string `gorm:"primaryKey"`
	Keyword        string
	DestinationUrl string
}

type Contact struct {
	Id        string `gorm:"primaryKey"`
	Subject   string
	Message   string
	IpAddress string
	CreatedAt int64
}

type AccessLog struct {
	Id             string `gorm:"primaryKey"`
	Timestamp      int64
	Url            string
	IpAddress      string
	Referrer       string
	Platform       string
	Mobile         string
	UserAgent      string
	UserAgentBrand string
}

type UrlRedirect struct {
	Id             string `gorm:"primaryKey"`
	Keyword        string
	DestinationUrl string
}

type Category struct {
	Id               string `gorm:"primaryKey"`
	ParentId         string
	Name             string
	RelativePosition float64
}

type Note struct {
	Id         string `gorm:"primaryKey"`
	CategoryId string
	Category   Category `gorm:"foreignKey:CategoryId"`
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
		&Admin{}, &ServerLog{}, &Contact{}, &AccessLog{},
		&UrlRedirect{}, &Category{}, &Note{},
	)

	if err != nil {
		logger.Error("Failed to migrate database schema: ", err)
		return err
	}

	logger.Debug("Database initialized successfully.")
	return nil
}
