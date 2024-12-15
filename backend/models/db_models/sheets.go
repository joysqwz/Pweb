package db_models

import (
	"encoding/json"
)

type Sheet struct {
	ID   uint64          `gorm:"primarykey" json:"id"`
	Data json.RawMessage `gorm:"type:jsonb;NOT NULL" json:"-"`
}
