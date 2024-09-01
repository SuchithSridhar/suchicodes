package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

// SQL commands stored at the bottom portion of the file

func InitializeDatabase(db *sql.DB, logger *slog.Logger) error {

	schema := []string{
		createAdminTable,
		createServerLogsTable,
		createContactTable,
		createAccessLogsTable,
		createUrlRedirectsTable,
		createCategoriesTable,
		createNotesTable,
	}

	tx, err := db.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to begin transaction during tables creations: %v", err))
		return err
	}

	for index, operation := range schema {
		if _, err := tx.Exec(operation); err != nil {
			logger.Error(fmt.Sprintf("Error in initialization of database for table with index %v: %v", index, err))
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logger.Error(fmt.Sprintf("Failed to rollback transaction during tables creation: %v", rollbackErr))
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		logger.Error(fmt.Sprintf("Failed to commit transaction during tables creation: %v", err))
		return err
	}

	logger.Debug("Database initialized successfully.")
	return nil
}

var (
	createAdminTable = `
CREATE TABLE IF NOT EXISTS Admin (
    id VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password_hash VARCHAR NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);
`
	createServerLogsTable = `
CREATE TABLE IF NOT EXISTS ServerLogs (
	id VARCHAR NOT NULL,
	keyword VARCHAR,
	destination_url VARCHAR,
	PRIMARY KEY (id)
);
`
	createContactTable = `
CREATE TABLE IF NOT EXISTS Contact (
    id VARCHAR NOT NULL,
    subject TEXT,
    message TEXT,
    ip_address VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);
`
	createAccessLogsTable = `
CREATE TABLE IF NOT EXISTS AccessLogs (
	id VARCHAR NOT NULL, 
	timestamp TIMESTAMP, 
	url TEXT, 
	ip_address VARCHAR, 
	referrer VARCHAR, 
	platform VARCHAR, 
	mobile VARCHAR, 
	user_agent VARCHAR, 
	user_agent_brand VARCHAR, 
	PRIMARY KEY (id)
);
`
	createUrlRedirectsTable = `
CREATE TABLE IF NOT EXISTS UrlRedirects (
	id VARCHAR NOT NULL,
	keyword VARCHAR,
	destination_url VARCHAR,
	PRIMARY KEY (id)
);
`
	createCategoriesTable = `
CREATE TABLE IF NOT EXISTS Categories (
	id VARCHAR NOT NULL,
	parent_id VARCHAR,
	name VARCHAR,
	relative_position DECIMAL,
	PRIMARY KEY (id),
	FOREIGN KEY(parent_id) REFERENCES Categories (id)
);
`
	createNotesTable = `
CREATE TABLE IF NOT EXISTS Notes (
	id VARCHAR NOT NULL,
	category_id VARCHAR NOT NULL,
	create_at TIMESTAMP,
	update_at TIMESTAMP,
	title VARCHAR,
	html TEXT,
	markdown TEXT,
	brief TEXT,
	assets_map TEXT,
	PRIMARY KEY (id),
	FOREIGN KEY (category_id) REFERENCES Categories (id)
);
`
)
