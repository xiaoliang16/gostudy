package serializer

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}

const (
	CodeCheckLogin = 401
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

func CheckLogin() Response {
	return Response{
		Code:  0,
		Msg:   "",
	}
}

func Err(errCode int , msg string, err error) Response {
	res := Response{
		Code:  0,
		Msg:   "",
	}
	return res
}

func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库失败"
	}
	return Err(CodeDBError, msg, err)
}

func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}