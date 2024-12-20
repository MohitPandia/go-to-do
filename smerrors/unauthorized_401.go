package smerrors

import (
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            INTERNAL SERVER ERROR                           */
/* -------------------------------------------------------------------------- */
func Unauthorized(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusUnauthorized
	smerror.Code = errorCode
	smerror.Type = errorType.Unauthorized
	smerror.Message = err
	res.Error = smerror
	// logs.Info("AUTHORIZATION",
	// 	logs.String("message", err),
	// 	logs.Int("code", errorCode),
	// 	logs.String("type", smerror.Type),
	// )

	// var slackReq models.SlackRequest
	// slackReq.ErrorType = errorType.Unauthorized
	// slackReq.StatusCode = errorCode
	// slackReq.Message = err
	// slackReq.Context = ctx
	// go slack.Notify(slackReq)

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
