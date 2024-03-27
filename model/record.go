package model

import "time"

type Record struct {
	ID      int        `gorm:"primary_key;auto_increment;column:id"`
	EntryAt time.Time  `gorm:"column:entry_at"`
	ExitAt  *time.Time `gorm:"column:exit_at"`
	UserID  int        `gorm:"column:user_id"`
}
