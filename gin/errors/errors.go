package errors

import "strconv"

type InternalError struct {
	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
}

func (e *InternalError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Msg
}

var (
	ErrorAny = &InternalError{Code: 400, Msg: "bad request"}
)
