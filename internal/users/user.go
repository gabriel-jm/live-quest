package users

import "time"

type User struct {
	Id          string    `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	DisplayName string    `json:"displayName" db:"displayName"`
	Email       string    `json:"email" db:"email"`
	Picture     string    `json:"picture" db:"picture"`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedAt"`
}
