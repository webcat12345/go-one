package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Email      string    `json:"email" sql:"type:varchar(255),notnull,unique"`
	Password   []byte    `json:"-" sql:",notnull"`
	IsDeleted  bool      `json:"is_deleted" sql:",notnull;default:false"`
	IsVerified bool      `json:"is_verified sql:",notnull;default:false"`
	CreatedAt  time.Time `json:"created_at" sql:",notnull;default:now()"`
	UpdatedAt  time.Time `json:"updated_at" sql:",notnull;default:now()"`
}
