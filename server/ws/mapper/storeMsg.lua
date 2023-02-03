---@diagnostic disable: undefined-global
local topic = KEYS[1]

local msg = ARGV[1]

local targetId = "tf_id_" .. topic

redis.call("INCR", targetId)
local id = redis.call("GET", targetId)

local target = "tf_zset_" .. topic

redis.call("ZADD", target, "NX", id, msg)
return id
