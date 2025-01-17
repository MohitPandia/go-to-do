package models

import "net/http"

type BaseResponse struct {
	Success    bool        `json:"success,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	MetaData   interface{} `json:"metadata,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
}

func InternalServerError() BaseResponse {
	var baseResponse BaseResponse
	baseResponse.Success = false
	baseResponse.Message = "Failure"
	baseResponse.StatusCode = http.StatusInternalServerError
	return baseResponse
}
