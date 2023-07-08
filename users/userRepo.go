package users

import (
	"Chat/base"
	"Chat/users/models"

	"gorm.io/gorm"
)

// IUserRepo
//
// Repository that contains methods to manipulate users
type IUserRepo interface {
	base.IBaseRepo[models.User]
	GetWithEmail(email string, password string) models.User
}

// UserRepo
//
// Implementations of IUserRepo
type UserRepo struct {
	base.BaseRepo[models.User]
}

func (ur *UserRepo) GetWithEmail(email string, password string) models.User {
	var user models.User
	ur.Db.Find(&user, "email= ?", email)
	return user
}

// NewUserRepo
//
// Returns the instance of UserRepository
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseRepo: base.BaseRepo[models.User]{
			Db: db,
		},
	}
}
