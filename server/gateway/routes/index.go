package routes

type GatewayRoutes struct {
	UserRoutes
	SysRoutes
	CaptchaRoutes
	WsRoutes
}

var GatewayRouter = new(GatewayRoutes)