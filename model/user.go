package model

import "blogServer/database"

var UserManager = UserDao{}

type User struct {
	UserName string `gorm:"user_name;primaryKey;not null" json:"userName" form:"userName"`
	Password string `gorm:"password;not null" json:"password" form:"password"`
}

func (user User) TableName() string {
	return "user"
}

type UserDao struct {
}

// GetUserWithCheck  验证并返回 user 数据
func (userDao UserDao) GetUserWithCheck(userName string, password string) (*User, error) {
	var user User
	tx := database.MysqlDb.Where("user_name = ? and password = ?", userName, password).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
