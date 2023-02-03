package pojo

type CGroupChan struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type CGroupDTO struct {
	Title    string       `json:"title"`
	Channels []CGroupChan `json:"channels"`
}