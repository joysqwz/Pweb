package database

import (
	"pweb/backend/external/httpserver"
	"pweb/backend/models/db_models"
)

var _ httpserver.StorageItf = &Database{}

func (db *Database) CheckLogin(login, password string) (bool, error) {
	var User db_models.User
	result := db.Table("users").Where("login = ?", login).Find(&User)
	return User.Password == password, result.Error
}

func (db *Database) FindLogin() (string, error) {
	var User db_models.User
	result := db.Table("users").Where("logged_in = ?", true).Find(&User)
	return User.Login, result.Error
}

func (db *Database) Logout() error {
	result := db.Table("users").Where("logged_in = ?", true).Update("logged_in", false)
	return result.Error
}
