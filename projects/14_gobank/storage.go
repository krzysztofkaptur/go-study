package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(int, *Account) error
	GetAccountById(int) (*Account, error)
	FetchAccounts() ([]*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=golang_test password=password sslmode=disable"
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

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `
		create table if not exists account (
			id serial primary key, 
			first_name varchar(40) not null, 
			last_name varchar(40) not null, 
			number integer, 
			balance integer, 
			created_at timestamp with time zone default current_timestamp
		);
	`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	query := `
		insert into account (first_name, last_name, number, balance, created_at)
		values ($1, $2, $3, $4, $5)
	`

	_, err := s.db.Query(query, a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt)

	return err
}

func (s *PostgresStore) UpdateAccount(id int, a *Account) error {
	query := `
		update account 
		set 
			first_name=coalesce($2, first_name),
			last_name=coalesce($3, last_name)
		where id=$1
	`

	_, err := s.db.Query(query, id, a.FirstName, a.LastName)

	return err
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := `
		delete from account where id=$1
	`

	_, err := s.db.Query(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	query := `
		select * from account where id=$1
	`

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) FetchAccounts() ([]*Account, error) {
	query := `
		select * from account;
	`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByNumber(number int) (*Account, error) {
	rows, err := s.db.Query("select * from account where number=$1", number)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number %d not found", number)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	
	err := rows.Scan(
		&account.ID, 
		&account.FirstName, 
		&account.LastName, 
		&account.Number, 
		&account.Balance, 
		&account.CreatedAt)

	return account, err
}