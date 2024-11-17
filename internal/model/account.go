package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UserID		uuid.UUID	`json:"user_id"`
	Username	string		`json:"username" binding:"required,min=3"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required,min=6"`
	CreatedAt	time.Time	`json:"created_at"`
}