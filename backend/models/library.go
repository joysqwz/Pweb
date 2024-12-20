package db_models

type Library struct {
	ID       uint64 `gorm:"primarykey" json:"login"`
	Category string `gorm:"type:string" json:"category"`
	Data     string `gorm:"type:string" json:"data"`
}
