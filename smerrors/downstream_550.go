package smerrors

import (
	"go-to-do/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                              DOWNSTREAM ERROR                              */
/* -------------------------------------------------------------------------- */
func Downstream(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := 550
	smerror.Code = errorCode
	smerror.Type = errorType.Downstream
	smerror.Message = err
	res.Error = smerror
	// logs.Error("DOWNSTREAM ERROR",
	// 	logs.String("message", err),
	// 	logs.Int("code", errorCode),
	// 	logs.String("type", smerror.Type),
	// )

	// var slackReq models.SlackRequest
	// slackReq.ErrorType = errorType.Downstream
	// slackReq.StatusCode = errorCode
	// slackReq.Message = err
	// slackReq.Context = ctx
	// go slack.Notify(slackReq)

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
