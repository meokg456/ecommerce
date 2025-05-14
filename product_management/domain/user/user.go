package user

type User struct {
	ID       int
	Username string
	Password string
	FullName string
}

type Storage interface {
	Register(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CheckIfUserExist(id int) error
}

func NewUser(username string, password string, fullName string) User {
	return User{
		Username: username,
		Password: password,
		FullName: fullName,
	}
}

func NewUserWithId(id int, username string, password string, fullName string) User {
	return User{
		ID:       id,
		Username: username,
		Password: password,
		FullName: fullName,
	}
}
