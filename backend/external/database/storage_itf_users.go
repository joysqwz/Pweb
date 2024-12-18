package database

import (
	"pweb/backend/external/httpserver"
	"pweb/backend/models/db_models"
)

var _ httpserver.StorageItf = &Database{}

func (db *Database) CheckLogin(login, password string) (bool, error) {
	var User db_models.User
	err := db.Table("users").Where("login = ?", login).Find(&User).Error
	if err != nil {
		return false, err
	}
	if User.Password == password {
		db.Logout()
		err = db.Table("users").Where("login = ?", login).Update("logged_in", true).Error
	}
	return User.Password == password, err
}

func (db *Database) FindLogin() (string, error) {
	var User db_models.User
	result := db.Table("users").Where("logged_in = ?", true).Find(&User)
	return User.Login, result.Error
}

func (db *Database) Register(login, password string) error {
	db.Logout()
	User := db_models.User{
		Password: password,
		Login:    login,
		LoggedIn: true,
	}
	return db.Table("users").Create(User).Error
}

func (db *Database) Logout() error {
	result := db.Table("users").Where("logged_in = ?", true).Update("logged_in", false)
	return result.Error
}
