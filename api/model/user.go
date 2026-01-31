package model

import (
	//"gorm.io/gorm"
)

type RegisterRequest struct {
    Username string `json:"username" binding:"required,min=2,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,max=32"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}