package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {

	connStr := "user=postgres dbname=postgres password=sarasa sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil

}

func (s *PostgressStore) CreateAccount() error {
	return nil
}

func (s *PostgressStore) UpdateAccount() error {
	return nil
}

func (s *PostgressStore) DeleteAccount() error {
	return nil
}

func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
