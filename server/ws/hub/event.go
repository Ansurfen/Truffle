package hub

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

type Event struct {
	Type uint8  `json:"type"`
	Data string `json:"data"`
}

func NewEvent(msg []byte) Event {
	var e Event
	err := json.Unmarshal(msg, &e)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

type IEvent interface {
	Type() uint8
}

type InitEvent struct {
	Subs []string `json:"subs"`
	User string   `json:"user"`
}

func NewInitEvent(data string) InitEvent {
	var e InitEvent
	err := json.Unmarshal([]byte(data), &e)
	if err != nil {
		zap.S().Warn(err)
	}
	return e
}

func (e InitEvent) Type() uint8 {
	return INIT
}

type UnicastEvent struct {
	Topic   string `json:"topic"`
	Msg     string `json:"msg"`
	User    string `json:"user"`
	Ok      bool
	Targets []string
}

func NewUnicastEvent(data string) UnicastEvent {
	var e UnicastEvent
	err := json.Unmarshal([]byte(data), &e)
	if err != nil {
		zap.S().Warn(err)
	}
	return e
}

func (e UnicastEvent) Type() uint8 {
	return UNICAST
}

type InfoEvent struct {
	Topic string `json:"topic"`
	User  string `json:"user"`
}

func NewInfoEvent(data string) InfoEvent {
	var e InfoEvent
	err := json.Unmarshal([]byte(data), &e)
	if err != nil {
		zap.S().Warn(err)
	}
	return e
}

func (e InfoEvent) Type() uint8 {
	return INFO
}

type UpdateEvent struct {
	// 这是要清空内存的
}
