package db_models

type User struct {
	Login    string `gorm:"primarykey"`
	Password string `gorm:"type:text"`
	Email    string `gorm:"type:string"`
	LoggedIn bool   `gorm:"type:bool"`
}
