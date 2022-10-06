package data

func InitStorage() error {
	storage, err := GetStorage()
	if err != nil {
		return err
	}
	storage.Init()
	return nil
}

func GetStorage() (
	Storage,
	error,
) {
	return NewRAMstorage(), nil
}

type Storage interface {
	NewUser(name, password string) error
	NewToken(token, name string) error

	CheckAccess(name, password string) error
	DeleteUser(name string)
	Init() error

	ChangePassword(token, password string) error
}

const (
	ErrorBadToken      = "bad token"
	ErrorAlreadyExist  = "already exist"
	ErrorWrongUserName = "wrong username"
	ErrorWrongPassword = "wrong password"
)
