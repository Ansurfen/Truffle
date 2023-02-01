package server

import (
	"truffle/db"
	"truffle/utils"
)

type UserMapper struct {
	node *utils.SnowFlake
}

func NewUserMapper() *UserMapper {
	return &UserMapper{
		node: utils.NewSnowFlake(1),
	}
}

func (mapper *UserMapper) CreateUser(user User) error {
	user.Id = mapper.node.Generate().Int64()
	return db.GetDB().Model(&User{}).Create(user).Error
}

func (*UserMapper) DeleteUser(user User) {
	db.GetDB().Model(&User{})
}

func (*UserMapper) FindByName(name string) (user User) {
	db.GetDB().Model(&User{}).Where("username = ?", name).First(&user)
	return user
}

func (*UserMapper) FindByTelephone(telephone string) (user User) {
	db.GetDB().Model(&User{}).Where("telephone = ?", telephone).First(&user)
	return user
}

func (*UserMapper) FindByEmail(email string) (user User) {
	db.GetDB().Model(&User{}).Where("email = ?", email).First(&user)
	return user
}

func (*UserMapper) Hash(user User) (string, int64) {
	now := utils.NowTimestamp()
	return utils.EncodeAESWithKey(utils.ToString(now)+"TRUFFLE", user.Username), now
}
