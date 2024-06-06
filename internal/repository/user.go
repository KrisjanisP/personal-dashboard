// repository/user_repository.go
package repository

import (
	"database/sql"
	"errors"

	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/go-jet/jet/v2/sqlite"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) internal.UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUsername(username string) (*domain.User, error) {
	stmt := sqlite.SELECT("*").FROM("users").WHERE(sqlite.COLUMN("username").EQ(username))
	query, args := stmt.Sql()

	var user domain.User
	err := r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) error {
	stmt := sqlite.INSERT("users").COLUMNS("username", "password").VALUES(user.Username, user.Password)
	query, args := stmt.Sql()

	_, err := r.db.Exec(query, args...)
	return err
}
