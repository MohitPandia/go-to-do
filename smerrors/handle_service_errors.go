package smerrors

import (
	"net/http"
	"go-to-do/models"

	"github.com/gin-gonic/gin"
)

func HandleServiceCodes(ctx *gin.Context, baseRes models.BaseResponse) {
	switch baseRes.StatusCode {
	case http.StatusUnauthorized:
		Unauthorized(ctx, baseRes.Message)
	case http.StatusServiceUnavailable:
		ServiceUnavailable(ctx, baseRes.Message)
	case http.StatusUnprocessableEntity:
		Validation(ctx, baseRes.Message)
	case 550:
		Downstream(ctx, baseRes.Message)
	default:
		InternalServer(ctx, baseRes.Message)
	}
}
