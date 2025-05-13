package postgresstore

import (
	"github.com/meokg456/userservice/dbconst"
	"github.com/meokg456/userservice/domain/user"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (userStore *UserStore) Register(user *user.User) error {
	data := UserQuerySchema{
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
	}
	result := userStore.db.Table(dbconst.UserTableName).Create(&data)
	user.ID = int(data.ID)
	return result.Error
}

func (userStore *UserStore) GetUserByUsername(username string) (*user.User, error) {
	var data UserQuerySchema
	result := userStore.db.Table(dbconst.UserTableName).First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	user := user.NewUserWithId(int(data.ID), data.Username, data.Password, data.FullName)
	return &user, nil
}

func (userStore *UserStore) GetUserById(id int) (*user.User, error) {
	data := UserQuerySchema{}
	result := userStore.db.Debug().Table(dbconst.UserTableName).First(&data, id)

	if result.Error != nil {
		return nil, result.Error
	}

	user := user.NewUserWithId(int(data.ID), data.Username, data.Password, data.FullName)
	return &user, nil
}

func (userStore *UserStore) CheckIfUserExist(id int) error {
	data := UserQuerySchema{
		Model: gorm.Model{ID: uint(id)},
	}
	result := userStore.db.Table(dbconst.UserTableName).First(&data)

	return result.Error
}

func (userStore *UserStore) UpdateProfile(user *user.User) error {
	data := UserQuerySchema{
		Model:    gorm.Model{ID: uint(user.ID)},
		FullName: user.FullName,
		Avatar:   user.Avatar,
	}
	result := userStore.db.Table(dbconst.UserTableName).Updates(&data)

	return result.Error
}
