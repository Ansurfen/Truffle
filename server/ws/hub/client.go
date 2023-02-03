package hub

import (
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (c *Client) Read() {
	defer c.conn.Close()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.S().Warn("error: %v", err)
			}
			break
		}
		evt := GetHub().con.parseMsg(message)
		if evt == nil {
			continue
		}
		if evt.Type() == INIT {
			GetHub().clis[evt.(InitEvent).User] = c
		}
		GetHub().channel <- evt
		GetHub().channel <- InitEvent{}
	}
}

func (c *Client) Write() {
	defer c.conn.Close()
	ticker := time.NewTicker(pingPeriod)
	for {
		select {
		case message, ok := <-c.buffer:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// 还有其他消息一块广播了
			n := len(c.buffer)
			for i := 0; i < n; i++ {
				w.Write([]byte("\n"))
				w.Write([]byte((<-c.buffer)))
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
