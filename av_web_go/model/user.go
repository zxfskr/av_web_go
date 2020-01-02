package model

import (
	"av_process/av_web_go/pkg/db"
)

// UserModel 控制user的数据库行为
type UserModel struct {
}

// User 系统用户
type User struct {
	UserID   int    `gorm:"column:user_id;" json:"userId"`
	Username string `json:"username"`
	Password string `gorm:"column:password_hash;" json:"password"`
	UserType string `gorm:"column:user_type;" json:"userType"`
}

// TableName 设置user表名
func (User) TableName() string {
	return "sys_user"
}

// GetOneByID 通过id查找一个用户
func (UserModel) GetOneByID(userID int) (*User, error) {
	user := &User{}
	tx := db.DB.Where("user_id = ?", userID).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

// GetOneByName 通过id查找一个用户
func (UserModel) GetOneByName(username string) (*User, error) {
	user := &User{}
	tx := db.DB.Where("username = ?", username).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

// GetAllByPage 获取所有用户
func (UserModel) GetAllByPage(limit int, offset int) ([]User, error) {
	users := make([]User, 0)

	sql := db.DB.Limit(limit).Offset(offset)
	tx := sql.Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}

	return users, nil
}
