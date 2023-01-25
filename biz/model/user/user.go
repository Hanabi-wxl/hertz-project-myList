package user

import (
	"hertz-mylist/base/result"
	"hertz-mylist/biz/model/gorm"
)

func Register(req *UserRequest) (*UserDetailResponse, error) {
	username := req.Username
	password := req.Password
	if password == "" {
		return nil, result.NewIError(10003, "输入密码为空")
	}
	var count int64
	gorm.DB.Model(User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return nil, result.NewIError(10004, "用户名已存在")
	}
	var userReg User
	if err := userReg.SetPassword(password); err != nil {
		return nil, result.NewIError(10005, "密码生成错误")
	}
	userReg.Username = username
	if err := gorm.DB.Create(&userReg).Error; err != nil {
		return nil, result.NewIError(10006, "创建用户失败")
	}
	return BuildUserDetailResponse(&userReg), nil
}
