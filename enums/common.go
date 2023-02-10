package enums

// web app 响应代码
type ResponseCode int

const (
	SUCCESS    ResponseCode = 200 // 成功
	BADREQUEST ResponseCode = 400 // 参数错误
	NOTLOGIN   ResponseCode = 401 // 未登录/未授权
	FORBIDDEN  ResponseCode = 403 // 无权限
	ERROR      ResponseCode = 500 // 异常
	FAILURE    ResponseCode = 501 // 失败
)
