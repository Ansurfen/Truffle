package hub

const (
	// msg event
	INIT = iota
	UNICAST
	MULTICAST
	BROADCAST
	SYNC
	INFO
	ASYNC  // dial sync content
	CLOASE // active close conncet

	// internal
	UPDATE
	PING // 完成初始化了在挂上客户端
)
