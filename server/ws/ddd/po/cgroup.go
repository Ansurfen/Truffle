package po

type CGroup struct {
	Id   int64  `json:"id" gorm:"primary_key;not null"`
	User string `json:"user" gorm:"not null"`
	Name string `json:"name" gorm:"not null"` // group name
	Path string `json:"path" gorm:"not null"` // sole identification to channel
}

func (CGroup) TableName() string {
	return "cgroup"
}
