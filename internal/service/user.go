package service

import (
	"errors"
	//"mygo_bangforai/api/error/response"
	"mygo_bangforai/api/model"
	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/utils"

	"gorm.io/gorm"
)

func Register(req model.RegisterRequest) (*model.User, error) {
	db := config.GetDB()

	var existingUser model.User
	result := db.Where("username = ?", req.Username).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("用户名已存在")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	result = db.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("邮箱已被注册")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	hashedPassword,err:= utils.HashPassword(req.Password)
	if err!=nil{
		return nil,err
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	result = db.Create(&user)
	
	return &user, nil
}

func Login(req model.LoginRequest) (*model.User, error) {
	db := config.GetDB()

	var user model.User
	result := db.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if !utils.CheckPassword(req.Password, user.Password){
		return nil,errors.New("密码错误")
	}

	return &user, nil
}