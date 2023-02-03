package routes

type WsRoutes struct {
	WsChannelRoutes
	WsMessageRoutes
	WsTopicRoutes
	WsServeRoutes
}

var WsRouter = new(WsRoutes)
