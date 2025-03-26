package api

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type SuccessResBody struct {
	Meta Meta
	Data any
}

type ErrorResBody struct {
	Meta Meta
	Err  any
}

func NewSuccessResponse(message string, code int, data any) SuccessResBody {
	return SuccessResBody{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "success",
		},
		Data: data,
	}
}

func NewErrorResponse(message string, code int, err interface{}) ErrorResBody {
	return ErrorResBody{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "error",
		},
		Err: err,
	}
}
