package users

import (
	"github.com/jmoiron/sqlx"

	"github.com/hashicorp/go-hclog"
)

type Store interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
	DeleteUser(id string) error
	UpdateUser(id string, user *User) error
}

type store struct {
	db  *sqlx.DB
	log hclog.Logger
}

func NewStore(db *sqlx.DB, log hclog.Logger) Store {
	return &store{
		db:  db,
		log: log,
	}
}

func (s *store) GetAllUsers() ([]User, error) {
	users := []User{}
	query := `SELECT * FROM users`

	err := s.db.Select(&users, query)
	if err != nil {
		s.log.Error("failed on getting all users")
		return nil, err
	}

	return users, nil
}

func (s *store) CreateUser(user *User) error {
	query := `INSERT INTO users (full_name, phone, address) VALUES (:fullName, :phone, :address)`
	queryValues := map[string]interface{}{"fullName": user.FullName, "phone": user.Phone, "address": user.Address}

	_, err := s.db.NamedExec(query, queryValues)
	if err != nil {
		s.log.Error("failed on creating user", "error", err.Error())
		return err
	}

	return nil
}

func (s *store) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=:id`
	queryValue := map[string]interface{}{"id": id}

	_, err := s.db.NamedExec(query, queryValue)
	if err != nil {
		s.log.Error("failed on deleting user", "error", err.Error())
		return err
	}
	return nil
}

func (s *store) UpdateUser(id string, user *User) error {
	query := `UPDATE users SET full_name=:fullName, phone=:phone, address=:address WHERE id=:id`
	queryValues := map[string]interface{}{"fullName": user.FullName, "phone": user.Phone, "address": user.Address, "id": id}

	_, err := s.db.NamedExec(query, queryValues)
	if err != nil {
		s.log.Error("failed on updating user", "error", err.Error())
		return err
	}
	return nil
}
