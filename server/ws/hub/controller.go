package hub

type HubController struct {
	topics *Trie
}

func NewHubController() *HubController {
	return &HubController{
		topics: NewRoot(),
	}
}

func (con *HubController) parseMsg(msg []byte) IEvent {
	e := NewEvent(msg)
	switch e.Type {
	case INIT:
		return con.Init(e.Data)
	case UNICAST:
		return con.Unicast(e.Data)
	case CLOASE:
	}
	return nil
}

func (con *HubController) Init(data string) InitEvent {
	e := NewInitEvent(data)
	for _, topic := range e.Subs {
		con.topics.Insert(topic)
		con.topics.UpdateTopic(topic, []string{e.User})
	}
	return e
}

func (con *HubController) Unicast(data string) UnicastEvent {
	e := NewUnicastEvent(data)
	e.Targets, e.Ok = con.topics.Unicast(e.User, e.Topic, e.Msg)
	return e
}

// abstract interface for WsController

func (con *HubController) GenPath(name string) string {
	path := con.topics.AddChild(ROOT)
	if node := con.topics.FindNode(path); node != nil && node.Topic != nil {
		node.Topic.Name = name
	}
	return path
}

func (con *HubController) Decoding(path string) string {
	return string(Decoding([]byte(path)))
}

func (con *HubController) Encoding(path string) string {
	return string(Encoding([]byte(path)))
}
