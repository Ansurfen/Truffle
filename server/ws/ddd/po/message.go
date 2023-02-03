package po

import "truffle/utils"

const (
	AttrMask       = 0b11111111
	IsBot          = 0b00000001
	IsBotMask      = AttrMask ^ IsBot
	IsMarkDown     = 0b00000010
	IsMarkDownMask = AttrMask ^ IsMarkDown
	HasCite        = 0b00000100
	HasCiteMask    = AttrMask ^ HasCite
	HasDec         = 0b00001000
	HasDecMask     = AttrMask ^ HasDec
	Highlight      = 0b00010000
	HighlightMask  = AttrMask ^ Highlight
)

type Message struct {
	Id         int64          `json:"id" gorm:"primary_key;not null"`
	Path       string         `json:"-" gorm:"not null"`
	Author     string         `json:"author" gorm:"not null"`
	Timestamp  int32          `json:"timestamp" gorm:"not null"`
	Attr       uint8          `json:"-" gorm:"not null"`
	Msg        string         `json:"msg" gorm:"not null"`
	Avatar     string         `json:"avatar" gorm:"not null"`
	RoleColor  string         `json:"role_color" gorm:"not null"`
	Declare    DeclareMessage `json:"dec" gorm:"-"`
	Cite       CiteMessage    `json:"cite" gorm:"-"`
	IsBot      bool           `json:"isbot" gorm:"-"`
	IsMarkdown bool           `json:"ismarkdown" gorm:"-"`
	HasCite    bool           `json:"hascite" gorm:"-"`
	HasDec     bool           `json:"hasdec" gorm:"-"`
	Highlight  bool           `json:"highlight" gorm:"-"`
	DeclareMessage
	CiteMessage
}

type DeclareMessage struct {
	Dec_Title   string `json:"-"`
	Dec_Md_msg  string `json:"-"`
	Dec_Ctx_msg string `json:"-"`
	Dec_Bottom  string `json:"-"`
}

type CiteMessage struct {
	Cite_Author string `json:"-" gorm:"column:cite_author;not null"`
	Cite_Msg    string `json:"-" gorm:"column:cite_msg;not null"`
	Cite_Avatar string `json:"-" gorm:"column:cite_avatar;not null"`
}

func (Message) TableName() string {
	return "message"
}

func (msg Message) AdapterIMsg() Message {
	if msg.Attr&IsMarkDownMask == IsMarkDown {
		msg.IsMarkdown = true
	}
	if msg.Attr&IsBotMask == IsBot {
		msg.IsBot = true
	}
	if msg.Attr&HighlightMask == Highlight {
		msg.Highlight = true
	}
	if msg.Attr&HasCiteMask == HasCite {
		msg.HasCite = true
	}
	if msg.Attr&HasDecMask == HasDec {
		msg.HasDec = true
	}
	msg.Cite = CiteMessage{
		Cite_Author: msg.Cite_Author,
		Cite_Msg:    msg.Cite_Msg,
		Cite_Avatar: msg.Cite_Avatar,
	}
	msg.Declare = DeclareMessage{
		Dec_Title:   msg.Dec_Title,
		Dec_Ctx_msg: msg.Dec_Ctx_msg,
		Dec_Md_msg:  msg.Dec_Md_msg,
		Dec_Bottom:  msg.Dec_Bottom,
	}
	return msg
}

func (msg Message) AdapterDB(id int64) Message {
	if msg.IsMarkdown {
		msg.Attr |= IsMarkDown
	}
	if msg.IsBot {
		msg.Attr |= IsBot
	}
	if msg.Highlight {
		msg.Attr |= Highlight
	}
	if msg.HasCite {
		msg.Attr |= HasCite
	}
	if msg.HasDec {
		msg.Attr |= HasDec
	}
	msg.Cite_Author = msg.Cite.Cite_Author
	msg.Cite_Avatar = msg.Cite.Cite_Avatar
	msg.Cite_Msg = msg.Cite.Cite_Msg
	msg.Dec_Title = msg.Declare.Dec_Title
	msg.Dec_Md_msg = msg.Declare.Dec_Md_msg
	msg.Dec_Ctx_msg = msg.Declare.Dec_Ctx_msg
	msg.Dec_Bottom = msg.Declare.Dec_Bottom
	msg.Id = id
	msg.Timestamp = int32(utils.NowTimestamp())
	return msg
}
