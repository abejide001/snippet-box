package mysql

import (
    "database/sql"

    "alexedwards.net/snippetbox/pkg/models"
)

// UserModel struct
type UserModel struct {
    DB *sql.DB
}

// Insert Method
func (m *UserModel) Insert(name, email, password string) error {
    return nil
}

// Authenticate maethod
func (m *UserModel) Authenticate(email, password string) (int, error) {
    return 0, nil
}

// Get Method
func (m *UserModel) Get(id int) (*models.User, error) {
    return nil, nil
}