package model

import "time"

type Record struct {
	ID      int        `gorm:"primary_key;auto_increment;column:id"`
	EnterAt time.Time  `gorm:"column:enter_at"`
	ExitAt  *time.Time `gorm:"column:exit_at"`
	UserID  int        `gorm:"column:user_id"`
	User    User       `gorm:"foreignKey:UserID"`
}
