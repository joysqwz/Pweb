package db_models

import (
	pq "github.com/lib/pq"
)

type User struct {
	Login    string        `gorm:"primarykey"`
	Password string        `gorm:"type:text"`
	Sheets   pq.Int64Array `gorm:"type:integer[]"`
	LoggedIn bool          `gorm:"type:bool"`
}
