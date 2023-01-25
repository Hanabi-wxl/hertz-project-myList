package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted int
}

const (
	PassWordCost = 12 // 密码加密难度
)

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func BuildUserDetailResponse(userReg *User) *UserDetailResponse {
	return &UserDetailResponse{
		Code: 200,
		UserDetail: &UserModel{
			ID:       int64(userReg.ID),
			Username: userReg.Username,
		},
	}
}
