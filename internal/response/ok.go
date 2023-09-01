package response

import "encoding/json"

func NewOk() *HttpResponse {
	return &HttpResponse{}
}

type HttpResponse struct {
	Message string `json:"message"`
}

func (r *HttpResponse) NewResponse(msg string) []byte {
	r.Message = msg
	jsonDate, _ := json.Marshal(&r)
	return jsonDate
}
