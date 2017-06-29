package model

import (
	"fmt"
)

//User 用户信息结构
type User struct {
	UID      uint64 `gorm:"type:int(11);primary_key;AUTO_INCREMENT" json:"uid"`
	Mobile   string `gorm:"column:mobile;type:varchar(45);not null;DEFAULT:''" json:"mobile"`
	Password string `gorm:"column:password;type:varchar(45);not null;DEFAULT:''" json:"password"`
}

type UserModel struct {
	BaseModel
}

//TableName 返回表名
func (user User) TableName() string {
	return "user"
}

//NewUserModel 创建新用户模型
func NewUserModel(name string) *UserModel {
	m := new(UserModel)
	m.DB = dbConnection.dbs[name]
	return m
}

//GetUser 获取个人信息
func (m *UserModel) GetUser(user *User) (*User, error) {
	u := &User{}
	fmt.Println(user.UID)
	db := m.DB.Table("user").Where(user).First(u)
	if db.RecordNotFound() {
		return nil, nil
	}
	if db.Error != nil {
		return nil, db.Error
	}
	return u, nil
}

//GetUserInfoByUID 根据UID获取个人信息
func (m *UserModel) GetUserInfoByUID(uid uint64) (*User, error) {
	return m.GetUser(&User{UID: uid})
}
