package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"truffle/db"
	"truffle/utils"
	. "truffle/utils"
)

func init() {
	opt := LoadOpt(ENV_DEVELOP, DefaultOpt{}, db.NoSQLOpt{}, LoggerOpt{})
	InitLogger(opt.Opt(LOGGER).(LoggerOpt))
	db.InitRedis(opt.Opt(db.NOSQL).(db.NoSQLOpt))
	db.GetRedis().Ping()
}

func TestRedis(t *testing.T) {
	db.GetRedis().Set("key", "value", time.Second*60)
	println(db.GetRedis().Get("key"))
	println(db.GetRedis().Get("key2"))
}

func TestRedisHash(t *testing.T) {
	topics := db.NewRHash("trie")
	topics.Set("aba", 1)
	fmt.Println(topics.Get("aba"))
	topics.Del("aba")
}

func TestRedisList(t *testing.T) {
	topicMsg := db.NewRList("truffle_topic_aba")
	topicMsg.Push("world")
	topicMsg.Unshift("hello")
	fmt.Println(topicMsg.Range(0, -1))
}

func TestRedisScan(t *testing.T) {
	for i := 0; i < 1000; i++ {
		db.GetRedis().SetNX("tf_aba_"+strconv.Itoa(i), i, time.Second*60)
	}
	fmt.Println(len(db.GetRedis().GetWithPrefix("tf_aba_", 1000)))
	db.GetRedis().Del(db.GetRedis().GetKeysWithPrefix("tf_aba_", 3)...)
}

const script = `
local topic = KEYS[1]

local msg = ARGV[1]

local targetId = "tf_id_" .. topic

local res = redis.call("INCR", targetId)
local id = redis.call("GET", targetId)

local target = "tf_zset_" .. topic

res = redis.call("ZADD", target, "NX", id, msg)
return res
`

func TestMessage(t *testing.T) {
	node := utils.NewSnowFlake(1)
	for i := 0; i < 100; i++ {
		primaryKey := node.Generate()
		msg := "{" + primaryKey.String() + "}"
		db.GetRedis().Eval(script, []string{"test"}, msg)
	}
	id := db.GetRedis().Get("tf_id_test")
	fmt.Println(db.GetRedis().ZRange("tf_zset_test", "0", id, 100))
}
