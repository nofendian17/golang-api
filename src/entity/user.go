package entity

import (
	"github.com/Kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UserName         string `json:"user_name" binding:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" binding:"required"`
}
