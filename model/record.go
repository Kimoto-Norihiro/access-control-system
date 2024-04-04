package model

import "time"

type Record struct {
	ID      int        `gorm:"primary_key;auto_increment;column:id" json:"id"`
	EnterAt time.Time  `gorm:"column:enter_at" json:"enter_at"`
	ExitAt  *time.Time `gorm:"column:exit_at" json:"exit_at"`
	UserID  int        `gorm:"column:user_id" json:"user_id"`
	User    User       `gorm:"foreignKey:UserID" json:"user"`
}
