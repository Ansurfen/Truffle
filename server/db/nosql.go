package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RClient struct {
	*redis.Client
	ctx context.Context
}

var rdb *RClient

// ? singleton pattern

func GetRedis() *RClient {
	return rdb
}

func InitRedis(opt NoSQLOpt) *RClient {
	if rdb == nil {
		rdb = &RClient{
			Client: redis.NewClient(&redis.Options{
				Addr:     opt.Addr,
				Password: opt.Pwd,
				DB:       0}),
			ctx: context.Background(),
		}
	}
	return rdb
}

func (rdb *RClient) Ping() {
	_, err := rdb.Client.Ping(rdb.ctx).Result()
	if err != nil {
		zap.S().Fatal("Connect redis failed! err : %v\n", err)
		return
	}
	zap.S().Info("Connect redis successfully!")
}

func (rdb *RClient) Set(key string, value any, expire time.Duration) {
	err := rdb.Client.Set(rdb.ctx, key, value, expire).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) Get(key string) string {
	val, err := rdb.Client.Get(rdb.ctx, key).Result()
	if err != nil {
		zap.S().Warn(err)
	}
	return val
}

func (rdb *RClient) Del(keys ...string) {
	err := rdb.Client.Del(rdb.ctx, keys...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) SetNX(key string, value any, expire time.Duration) {
	err := rdb.Client.SetNX(rdb.ctx, key, value, expire).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) Do(args ...any) {
	err := rdb.Client.Do(rdb.ctx, args...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) Eval(script string, keys []string, args ...any) (any, error) {
	res, err := rdb.Client.Eval(rdb.ctx, script, keys, args...).Result()
	if err != nil {
		zap.S().Warn(err)
		return res, err
	}
	return res, nil
}

func (rdb *RClient) HSet(key string, values ...any) {
	err := rdb.Client.HSet(rdb.ctx, key, values...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) HGet(key, filed string) string {
	val, err := rdb.Client.HGet(rdb.ctx, key, filed).Result()
	if err != nil {
		zap.S().Warn(err)
	}
	return val
}

func (rdb *RClient) HDel(key string, fields ...string) {
	err := rdb.Client.HDel(rdb.ctx, key, fields...).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) LPush(key string, values ...any) {
	err := rdb.Client.LPush(rdb.ctx, key, values).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) RPush(key string, values ...any) {
	err := rdb.Client.RPush(rdb.ctx, key, values).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) LRange(key string, start, end int64) []string {
	return rdb.Client.LRange(rdb.ctx, key, start, end).Val()
}

func (rdb *RClient) Incr(key string) {
	err := rdb.Client.Incr(rdb.ctx, key).Err()
	if err != nil {
		zap.S().Warn(err)
	}
}

func (rdb *RClient) GetWithPrefix(match string, count int64) []string {
	var cursor uint64
	match = "*" + match + "*"
	keys, cursor, err := rdb.Client.Scan(rdb.ctx, cursor, match, count+1).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	var ret []string
	for _, key := range keys {
		ret = append(ret, rdb.Get(key))
	}
	return ret
}

func (rdb *RClient) GetKeysWithPrefix(match string, count int64) []string {
	match = "*" + match + "*"
	var cursor uint64
	keys, cursor, err := rdb.Client.Scan(rdb.ctx, cursor, match, count+1).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	var ret []string
	for _, key := range keys {
		ret = append(ret, rdb.Get(key))
	}
	return ret
}

func (rdb *RClient) GetWithKeyAndPrefix(key, match string, count int64) []string {
	var cursor uint64
	keys, cursor, err := rdb.Client.SScan(rdb.ctx, key, cursor, match, count).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return keys
}

func (rdb *RClient) GetKeysWithKeyAndPrefix(key, match string, count int64) []string {
	var cursor uint64
	keys, cursor, err := rdb.Client.SScan(rdb.ctx, key, cursor, match, count).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return keys
}

func (rdb *RClient) ZRange(key, start, end string, count int64) []string {
	res, err := rdb.ZRangeByScore(rdb.ctx, key, &redis.ZRangeBy{
		Min:   start,
		Max:   end,
		Count: count,
	}).Result()
	if err != nil {
		zap.S().Warn(err)
		return nil
	}
	return res
}

type RList struct {
	instance string
}

func NewRList(instance string) *RList {
	return &RList{
		instance: instance,
	}
}

func (list *RList) Unshift(values ...any) {
	GetRedis().LPush(list.instance, values...)
}

func (list *RList) Push(values ...any) {
	GetRedis().RPush(list.instance, values...)
}

func (list *RList) Range(start, end int64) []string {
	return GetRedis().LRange(list.instance, start, end)
}

type RHash struct {
	instance string
}

func NewRHash(instance string) *RHash {
	return &RHash{
		instance: instance,
	}
}

func (hash *RHash) Set(values ...any) {
	GetRedis().HSet(hash.instance, values...)
}

func (hash *RHash) Get(filed string) string {
	return GetRedis().HGet(hash.instance, filed)
}

func (hash *RHash) Del(fileds ...string) {
	GetRedis().HDel(hash.instance, fileds...)
}
