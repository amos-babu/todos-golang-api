package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB(dbUsername, dbPassword, dbName string) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUsername, dbPassword, dbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	createTableMigrations(DB)
	return DB, nil
}

func createTableMigrations(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS todos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            description VARCHAR(255) NOT NULL,
            createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating table: ", err)
	}

	fmt.Println("Table created succefully!")
}
