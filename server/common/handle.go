package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinWrap[Request any](f func(req Request) Result) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request Request
		if err := ctx.Bind(&request); err != nil {
			zap.S().Warn(err)
			ctx.JSON(http.StatusInternalServerError, JSON_UNMARSHAL_ERR)
		}
		ret := f(request)
		switch ret.Code {
		case http.StatusOK:
			ctx.JSON(http.StatusOK, ret)
		default:

		}
	}
}
