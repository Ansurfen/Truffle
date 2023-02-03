package vo

import "truffle/ws/ddd/po"

// @NewTopicRequestVo2Dto req:v2d topic.go
type NewTopicRequest struct {
	User string `json:"user" form:"user"`
	Name string `json:"name" form:"name"`
}

type NewTopicResponse struct {
	Path string `json:"path"`
}

// @JoinTopicRequestVo2Dto req:v2d topic.go
type JoinTopicRequest struct {
	User     string `json:"user" form:"user"`
	Passcode string `json:"passcode" form:"passcode"`
}

type JoinTopicResponse struct {
	Topic po.Topic `json:"topic"`
}
