package hub

import (
	"encoding/json"
	"net/http"
	"time"
	truffle_log "truffle/log/proto"

	"github.com/gorilla/websocket"
)

const (
	WAITING_INIT = iota
	WAITING_SYC

	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

var (
	hub      *Hub
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type (
	CStatus int

	Client struct {
		conn   *websocket.Conn
		buffer chan []byte
		status CStatus
		lc     truffle_log.LogClient
	}

	Hub struct {
		channel chan IEvent
		clis    map[string]*Client
		repo    map[string][]string
		con     *HubController
	}
)

func ServeWs(lc truffle_log.LogClient, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &Client{conn: conn, buffer: make(chan []byte, 20), lc: lc}
	go c.Read()
	go c.Write()
}

func newHub() *Hub {
	return &Hub{
		channel: make(chan IEvent, 20),
		clis:    make(map[string]*Client),
		repo:    make(map[string][]string),
		con:     NewHubController(),
	}
}

func InitHub() {
	hub = newHub()
}

func GetHub() *Hub {
	return hub
}

func (hub *Hub) WaitSignal() chan IEvent {
	return hub.channel
}

func (hub *Hub) Run() {
	for {
		select {
		case evt := <-hub.channel:
			switch evt.Type() {
			case UPDATE:
				// 一个注意的点，dict 注册过的路由要是以后删了就不能用了，等于得缓存起来做映射就行...

			case INIT:
				ie := evt.(InitEvent)
				hub.clis[ie.User].buffer <- []byte("初始化成功")
				// 查缓存告诉他还有啥没更新
				for _, topic := range ie.Subs {
					if len(hub.repo[topic]) > 0 {
						hub.clis[ie.User].buffer <- []byte(topic + "有消息")
					}
				}
			case CLOASE:

			case BROADCAST:

			case UNICAST:
				ue := evt.(UnicastEvent)
				for _, topic := range ue.Targets {
					hub.clis[topic].buffer <- []byte(ue.Msg)
				}
				hub.repo[ue.Topic] = append(hub.repo[ue.Topic], ue.Msg)
				// client.Logger(hub.clis[ue.User].lc, ue.Topic, ue.Msg)
			case INFO:
				ie := evt.(InfoEvent)
				msgs, _ := json.Marshal(hub.repo[ie.Topic])
				hub.clis[ie.User].buffer <- msgs
			}
		}
	}
}

func (hub *Hub) GetAbstractHub() *HubController {
	return hub.con
}
