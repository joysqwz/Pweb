package db_models

import (
	"encoding/json"
)

type Sheet struct {
	Login string          `gorm:"primarykey" json:"login"`
	Data  json.RawMessage `gorm:"type:jsonb;NOT NULL" json:"-"`
}
