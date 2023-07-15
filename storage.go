package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateBook(*Book) error
	DeleteBook(int) error
	UpdateBook(*Book) error
	GetBookById(int) (*Book, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:gohon@localhost?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateBook(*Book) error {
	return nil
}

func (s *PostgresStore) DeleteBook(int) error {
	return nil
}

func (s *PostgresStore) UpdateBook(*Book) error {
	return nil
}

func (s *PostgresStore) GetBookById(int) (*Book, error) {
	return nil, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateBookTable()
}

func (s *PostgresStore) CreateBookTable() error {
	query := `create table if not exists book(
		id uuid primary key,
		title varchar(255) not null)`

	_, err := s.db.Exec(query)
	return err
}
