package exception

type BizError struct {
	Code int32
	Msg  string
}

func NewBizError(code int32, msg string) *BizError {
	return &BizError{Code: code, Msg: msg}
}

func (b BizError) Error() string {
	return b.Msg
}
