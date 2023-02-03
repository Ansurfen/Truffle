package po

type Channel struct {
	Id       int64  `json:"id" gorm:"primary_key;not null"`
	Topic    string `json:"topic" gorm:"not null"`
	Path     string `json:"path" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Passcode string `json:"-" gorm:"not null"`
}

func (Channel) TableName() string {
	return "channel"
}
