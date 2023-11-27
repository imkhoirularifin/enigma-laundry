package repository

import (
	"database/sql"
	"fmt"
	"time"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/utils/common"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	// get by username ketika login
	GetByUsername(username string) (model.User, error)
	Create(payload model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.CreateUser, payload.Name, payload.Email, payload.Username, payload.Password, payload.Role, time.Now()).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetByUsername(username string) (model.User, error) {
	fmt.Println()
	var user model.User

	err := u.db.QueryRow(common.GetUserByUsername, username).Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &user.Role)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetUserById, id).
		Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
