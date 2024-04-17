package controllers

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ReturnSuccess(data interface{}) *JsonStruct {
	return &JsonStruct{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func ReturnError(code int, msg string) *JsonStruct {
	return &JsonStruct{
		Code: code,
		Msg:  msg,
	}
}
