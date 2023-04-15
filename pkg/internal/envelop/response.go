package envelop

type ResponseEnveloped struct {
	Code int32  `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data,omitempty"`
}

func NewResponseSuccess(data any) *ResponseEnveloped {
	return &ResponseEnveloped{
		Code: 0,
		Desc: "success",
		Data: data,
	}
}

func NewResponseError(code int32, desc string) *ResponseEnveloped {
	return &ResponseEnveloped{
		Code: code,
		Desc: desc,
	}
}

func NewValidateError(code int32) *ResponseEnveloped {
	return &ResponseEnveloped{
		Code: code,
		Desc: "invalid request",
	}
}
