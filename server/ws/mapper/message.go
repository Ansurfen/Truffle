package mapper

import (
	"strconv"
	"truffle/common"
	"truffle/db"
	"truffle/utils"
	"truffle/ws/ddd/dao"
	"truffle/ws/ddd/po"

	"go.uber.org/zap"
)

const storeMsg = `
local topic = KEYS[1]

local msg = ARGV[1]

local targetId = "tf_id_" .. topic

redis.call("INCR", targetId)
local id = redis.call("GET", targetId)

local target = "tf_zset_" .. topic

redis.call("ZADD", target, "NX", id, msg)
return id
`

type MessageMapper struct {
	Node *utils.SnowFlake
}

func NewMessageMapper() *MessageMapper {
	return &MessageMapper{
		Node: utils.NewSnowFlake(1),
	}
}

func (mapper *MessageMapper) GetMsgId(topic string) int {
	id, err := strconv.Atoi(db.GetRedis().Get("tf_id_" + topic))
	if err != nil {
		zap.S().Warn(err)
	}
	return id
}

func (mapper *MessageMapper) CreateMsgCache(topic, msg string) (string, error) {
	id, err := db.GetRedis().Eval(storeMsg, []string{topic}, msg)
	return id.(string), err
}

func (mapper *MessageMapper) CreateMessage(msg po.Message) {
	if err := db.GetDB().Model(&po.Message{}).Create(msg).Error; err != nil {
		zap.S().Warn(err)
		return
	}
}

func (mapper *MessageMapper) GetMessage(req *dao.GetMsgRequest) *dao.GetMsgResponse {
	res := new(dao.GetMsgResponse)
	err := db.GetDB().Model(&po.Message{}).Where("path = ?", req.Path).Offset((req.Base - 1) * int64(req.PageSize)).Limit(req.Limit).Order("id DESC").Find(&res.Msgs).Error
	if err != nil {
		res.Meta = common.CommonResponse{
			Err:  err,
			Code: common.DB_ERR_CODE,
		}
		zap.S().Warn(err)
	}
	return res
}
