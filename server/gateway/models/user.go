package models

type LoginRequest struct {
	Key string `json:"key" form:"key"`
	Pwd string `json:"pwd" form:"pwd"`
}

type RegisterRequest struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
	Key  string `json:"key" form:"key"`
}
