package smerrors

import (
	"go-to-do/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

/* -------------------------------------------------------------------------- */
/*                            VALIDATION ERROR 422                            */
/* -------------------------------------------------------------------------- */
func Validation(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var smerror Error
	errorCode := http.StatusUnprocessableEntity
	smerror.Code = errorCode
	smerror.Type = errorType.validation
	smerror.Message = err
	res.Error = smerror
	// logs.Info("VALIDATION",
	// 	logs.String("message", err),
	// 	logs.Int("code", errorCode),
	// 	logs.String("type", smerror.Type),
	// )

	// var slackReq models.SlackRequest
	// slackReq.ErrorType = errorType.validation
	// slackReq.StatusCode = errorCode
	// slackReq.Message = err
	// slackReq.Context = ctx
	// go slack.Notify(slackReq)

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
