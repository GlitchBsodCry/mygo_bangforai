package service

import (
	stderr "errors"
	"mygo_bangforai/api/model"
	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/utils"
	"mygo_bangforai/pkg/interfacer"
	"mygo_bangforai/api/errors"

	"gorm.io/gorm"
	"go.uber.org/zap"
)

var logger = interfacer.GetLogger()

func Register(req model.RegisterRequest) (*model.User, error) {
	db := config.GetDB()

	var existingUser model.User
	result := db.Where("username = ?", req.Username).First(&existingUser)
	if result.Error == nil {
		logger.Error("用户名已存在", zap.String("username", req.Username))
		err:=stderr.New("用户名已存在")
		return nil, err
	} else if !stderr.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.WrapError(result.Error, errors.DatabaseError, "查询用户失败", "internal/service/user.go/Register")
	}

	result = db.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		logger.Error("邮箱已被注册", zap.String("email", req.Email))
		err:=stderr.New("邮箱已被注册")
		return nil, err
	} else if !stderr.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.WrapError(result.Error, errors.DatabaseError, "查询用户失败", "internal/service/user.go/Register")
	}

	hashedPassword,err:= utils.HashPassword(req.Password)
	if err!=nil{
		return nil, errors.WrapError(err, errors.UtilsError, "密码哈希失败", "internal/service/user.go/Register")
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	err = db.Create(&user).Error
	if err!= nil{
		return nil, errors.WrapError(err, errors.DatabaseError, "创建用户失败", "internal/service/user.go/Register")
	}
	
	return &user, nil
}

func Login(req model.LoginRequest) (*model.User, error) {
	db := config.GetDB()

	var user model.User
	result := db.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		if stderr.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Error("用户名不存在", zap.String("username", req.Username))
			err:=stderr.New("用户名不存在")
			return nil, err
		}
		return nil, errors.WrapError(result.Error, errors.DatabaseError, "查询用户失败", "internal/service/user.go/Login")
	}

	if !utils.CheckPassword(req.Password, user.Password){
		logger.Error("密码错误", zap.String("username", req.Username))
		err:=stderr.New("密码错误")
		return nil, err
	}

	return &user, nil
}