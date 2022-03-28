package model

type ResponseModel struct {
	ResCode    int         `json:"code"`
	ResMessage string      `json:"message"`
	ResData    interface{} `json:"data"`
}
