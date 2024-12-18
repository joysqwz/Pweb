package httpserver

type StorageItf interface {
	CheckLogin(login, password string) (bool, error)
	FindLogin() (string, error)
	Logout() error
	Register(login, password string) error
}
