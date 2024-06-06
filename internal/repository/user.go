// repository/user_repository.go
package repository

import (
	"errors"

	"github.com/KrisjanisP/personal-dashboard/internal/database/model"
	"github.com/KrisjanisP/personal-dashboard/internal/database/table"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/sqlite"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUsername(username string) (*domain.User, error) {
	stmt := sqlite.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(
		table.Users.Username.EQ(sqlite.String(username)))

	var userRecord model.Users
	err := stmt.Query(r.db, &userRecord)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user := domain.User{
		ID:        *userRecord.ID,
		Username:  userRecord.Password,
		Password:  userRecord.Password,
		CreatedAt: *userRecord.CreatedAt,
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(userID int32) (*domain.User, error) {
	stmt := sqlite.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(
		table.Users.ID.EQ(sqlite.Int32(userID)))

	var userRecord model.Users
	err := stmt.Query(r.db, &userRecord)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user := domain.User{
		ID:        *userRecord.ID,
		Username:  userRecord.Password,
		Password:  userRecord.Password,
		CreatedAt: *userRecord.CreatedAt,
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (int32, error) {
	stmt := table.Users.INSERT(table.Users.Username, table.Users.Password).
		VALUES(user.Username, user.Password).RETURNING(table.Users.ID)

	var userRecord model.Users
	err := stmt.Query(r.db, &userRecord)
	if err != nil {
		return 0, err
	}

	return *userRecord.ID, nil
}
