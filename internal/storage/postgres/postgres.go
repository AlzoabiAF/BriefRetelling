package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Data struct {
	url  string
	info []string
}

func New(connStr string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS textStorage(id serial PRIMARY KEY, header TEXT NOT NULL, info TEXT NOT NULL, url TEXT NOT NULL UNIQUE);")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db}, nil
}

func (s *Storage) NewData(header string, info string, url string) error {
	const op = "storage.postgres.NewData"

	_, err := s.db.Exec("INSERT INTO textStorage(header, info, url) VALUES ($1, $2, $3);", header, info, url)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
