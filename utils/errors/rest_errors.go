package errors

type RestErrors struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Error   string      `json:"error"`
}

type ValidationErrors struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}
