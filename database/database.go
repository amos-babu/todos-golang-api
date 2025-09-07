package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(dbUsername, dbPassword, dbName string) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbName, url.QueryEscape(dbPassword))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTableMigrations(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTableMigrations(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            description VARCHAR(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Errorf("Error Creating Tables: %w", err)
	}

	fmt.Println("✅ Table created succefully!")

	return nil
}
