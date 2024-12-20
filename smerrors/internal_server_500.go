package smerrors

import (
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            INTERNAL SERVER ERROR                           */
/* -------------------------------------------------------------------------- */
func InternalServer(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusInternalServerError
	smerror.Code = errorCode
	smerror.Type = errorType.server
	smerror.Message = "something went wrong"
	res.Error = smerror
	// logs.Error("SERVER",
	// 	logs.String("message", err),
	// 	logs.Int("code", errorCode),
	// 	logs.String("type", smerror.Type),
	// )

	// var slackReq models.SlackRequest
	// slackReq.ErrorType = errorType.server
	// slackReq.StatusCode = errorCode
	// slackReq.Message = err
	// slackReq.Context = ctx
	// go slack.Notify(slackReq)

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
