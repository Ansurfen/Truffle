package server

import "database/sql"

const (
// 平台管理员
// 平台用户
// 群管理员...
)

type User struct {
	Username  string         `json:"username" gorm:"varchar(110);not null;unique"`
	Alias     string         `json:"alias" gorm:"varchar(110);not null"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	Telephone sql.NullString `json:"telephone" gorm:"varchar(110);unique"`
	Email     sql.NullString `json:"email" gorm:"varchar(110);unique"`
	Id        int64          `json:"id" gorm:"primary_key;not null"`
	Power     uint8          `json:"power" gorm:"not null"`
	Profile   string         `json:"profile" gorm:"varchar(32);"`
	Lang      string         `json:"lang" gorm:"varchar(5);"`
}

func (User) TableName() string {
	return "user"
}
