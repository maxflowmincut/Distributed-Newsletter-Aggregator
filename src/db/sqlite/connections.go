package sqlite

import (
	"database/sql"
	"log"
)

func ConnectDB(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func InitializeDB(db *SQLiteDB, schema string) error {
	_, err := db.DB.Exec(schema)
	return err
}

func CloseDB(db *SQLiteDB) {
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}