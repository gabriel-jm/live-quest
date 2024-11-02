package users

import (
	"live-quest/internal/database"
	"live-quest/internal/id"
	"time"
)

type UserRepository struct{}

func (u *UserRepository) Add(user ReqUser) (*User, error) {
	userData := User{
		Id:          id.NanoId(),
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Picture:     user.Picture,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `
		insert into users
		values (:id, :username, :displayName, :email, :picture, :createdAt, :updatedAt)
	`

	_, err := database.Db.NamedExec(query, &userData)

	if err != nil {
		return nil, err
	}

	return &userData, nil
}
