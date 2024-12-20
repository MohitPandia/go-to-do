package smerrors

import (
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                       SERVICE UNAVAILABLE ERROR 503                        */
/* -------------------------------------------------------------------------- */
func ServiceUnavailable(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusServiceUnavailable
	smerror.Code = errorCode
	smerror.Type = errorType.ServiceUnavailable
	smerror.Message = "service unavailable, please try again after some time"
	res.Error = smerror
	// logs.Warn("SERVICE UNAVAILABLE",
	// 	logs.String("message", err),
	// 	logs.Int("code", errorCode),
	// 	logs.String("type", smerror.Type),
	// )

	// var slackReq models.SlackRequest
	// slackReq.ErrorType = errorType.ServiceUnavailable
	// slackReq.StatusCode = errorCode
	// slackReq.Message = err
	// slackReq.Context = ctx
	// go slack.Notify(slackReq)

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
