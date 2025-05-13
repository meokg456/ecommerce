package postgresstore

import (
	"github.com/meokg456/sampleservice/dbconst"
	"github.com/meokg456/sampleservice/domain/user"

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
