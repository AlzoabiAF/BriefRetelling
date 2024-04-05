package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // init sqlite driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS textStorage(id INTEGER PRIMARY KEY, url TEXT NOT NULL UNIQUE, info TEXT NOT NULL);")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db}, nil
}

func (s *Storage) NewData(url string, info string) error {
	const op = "storage.sqlite.NewData"

	_, err := s.db.Exec("INSERT INTO textStorage(url, info) VALUES ($1, $2);", url, info)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
