package dao

type GetCGroupsRequset struct {
	Path string
	User string
}

type GetCGroupResponse struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Cname string `json:"cname"`
}
