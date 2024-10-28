package gormrepo

import (
	"fmt"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type GormUser struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"` // Store a hashed password
	Role     string `gorm:"default:user"`
}

func (us *GormUser) GetId() string {
	return fmt.Sprintf("%d", us.ID)
}
func (us *GormUser) GetUserName() string {
	return us.Username
}
func (us *GormUser) GetPassword() string {
	return "" // return error that saved user cannnot return password
}
func (us *GormUser) GetHashedPassword() string {
	return us.Password
}

func (us *GormUser) SetHashedPassword(password string) {
	us.Password = password
}

func (us *GormUser) GetRole() string {
	if us.Role == "" {
		us.Role = "user"
	}
	return us.Role
}

func fromUserEntity(user entity.User) *GormUser {
	if gUser, ok := user.(*GormUser); ok {
		return gUser
	} else {
		gUser = &GormUser{}
		ID, err := strconv.ParseUint(user.GetId(), 10, 64)
		if nil != err {
			// panic(err)
		} else {
			gUser.ID = ID
		}
		// gUser := &GormUser{}
		gUser.Username = user.GetUserName()
		gUser.Password = user.GetHashedPassword()
		gUser.Role = user.GetRole()
		return gUser
	}
}
