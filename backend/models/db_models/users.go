package db_models

import (
	pq "github.com/lib/pq"
)

type User struct {
	ID       uint64        `gorm:"primarykey" json:"id"`
	Login    string        `gorm:"type:text"`
	Password string        `gorm:"type:text"`
	Sheets   pq.Int64Array `gorm:"type:integer[]"`
}
