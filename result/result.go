package result

func Success(data any) map[string]any {
	result := make(map[string]any)
	result["code"] = 0
	result["msg"] = "success"
	result["data"] = data
	return result
}

func ParamError(msg string) map[string]any {
	result := make(map[string]any)
	result["code"] = 1
	result["msg"] = msg
	return result
}
func BizError(code int32, msg string) map[string]any {
	result := make(map[string]any)
	result["code"] = code
	result["msg"] = msg
	return result
}
func SysError(msg string) map[string]any {
	result := make(map[string]any)
	result["code"] = 3
	result["msg"] = msg
	return result
}
