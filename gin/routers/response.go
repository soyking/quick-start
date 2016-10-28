package routers

import . "github.com/soyking/quick-start/gin/errors"

const (
	RESP_SUCCESS_CODE  = 0
	RESP_SUCCESS_MSG   = "success"
	UNKNOWN_ERROR_CODE = 500
)

type response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

func buildSuccessResponse(result interface{}) *response {
	return &response{
		Code:   RESP_SUCCESS_CODE,
		Msg:    RESP_SUCCESS_MSG,
		Result: result,
	}
}

func buildErrorResponse(err error) *response {
	r := new(response)
	if iErr, ok := err.(*InternalError); ok {
		r.Code = iErr.Code
		r.Msg = iErr.Msg
	} else {
		r.Code = UNKNOWN_ERROR_CODE
		r.Msg = err.Error()
	}
	return r

}

func buildResponse(result interface{}, err error) *response {
	if err != nil {
		return buildErrorResponse(err)
	} else {
		return buildSuccessResponse(result)
	}
}

type pageResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data" bson:"data"`
}

func buildPageResponse(total int, data interface{}, err error) *response {
	if err != nil {
		return buildErrorResponse(err)
	} else {
		return buildSuccessResponse(
			&pageResponse{
				Total: total,
				Data:  data,
			},
		)
	}
}
