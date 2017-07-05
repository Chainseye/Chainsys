package service

import (
	"chainseye.com/model"
)

//GetUserInfoByUID 根据UID获取个人信息
func GetUserInfoByUID(uid uint64) (*model.User, error) {
	um := model.NewUserModel(model.DEFAULT_MODEL)
	ur, err := um.GetUserInfoByUID(1)
	if err != nil {
		return nil, err
	}
	return ur, nil
}
