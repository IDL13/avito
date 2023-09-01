package response

import "encoding/json"

func NewErr() *HttpError {
	return &HttpError{}
}

type HttpError struct {
	Message string `json:"message"`
}

func (e *HttpError) NewError(msg string) []byte {
	e.Message = msg
	jsonDate, _ := json.Marshal(&e)
	return jsonDate
}
