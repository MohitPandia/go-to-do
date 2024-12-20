package utils

import "encoding/json"

func WrongJsonValidator(body []byte) []byte {
	if string(body) == "" {
		return nil
	}

	if len(body) == 0 {
		return nil
	} else if isValid := json.Valid(body); !isValid {
		resp, _ := json.Marshal([]interface{}{
			struct {
				WrongJsonReq string `json:"wrong_json"`
			}{
				WrongJsonReq: string(body),
			},
		})
		return resp
	}
	return body
}
