package ecode

// All common ecode
var (
	OK = add(200) // 正确

	NotModified        = add(-304) // 木有改动
	TemporaryRedirect  = add(-307) // 撞车跳转
	RequestErr         = add(-400) // 请求错误
	Unauthorized       = add(-401) // 未认证
	AccessDenied       = add(-403) // 访问权限不足
	NothingFound       = add(-404) // 啥都木有
	MethodNotAllowed   = add(-405) // 不支持该方法
	Conflict           = add(-409) // 冲突
	Canceled           = add(-498) // 客户端取消请求
	ServerErr          = add(-500) // 服务器错误
	ServiceUnavailable = add(-503) // 过载保护,服务暂不可用
	Deadline           = add(-504) // 服务调用超时
	LimitExceed        = add(-509) // 超出限制

	Degrade     = add(-1200) // 被降级过滤的请求
	RPCNoClient = add(-1201) // rpc服务的client都不可用
	RPCNoAuth   = add(-1202) // rpc服务的client没有授权
)

var HttpStatusCodeMsgFlags = map[Code]string{
	OK:           "请求成功",
	RequestErr:   "请求错误",
	Unauthorized: "未认证",
	AccessDenied: "访问权限不足",
	NothingFound: "请求的资源不存在",
}

var (
	Fail          = add(500) //错误
	ParamsInvalid = add(400) //参数错误

	SignInCodeSendFailed = add(10001) //发送动态验证码失败
	AuthCodeCheckFailed  = add(20001) //验证码不匹配
	AuthFailed           = add(20002) //验证失败

	UserPhoneNotExisted    = add(30001) //用户手机号不存在
	UserAssetTransferEqual = add(30002) //用户转移资产手机号不能跟源手机号一致
)

var BusinessStatusCodeMsgFlags = map[Code]string{
	Fail:                 "服务器发生错误",
	ParamsInvalid:        "参数错误",
	SignInCodeSendFailed: "发送动态验证码失败",
	AuthCodeCheckFailed:  "验证码不匹配",
	AuthFailed:           "验证失败",

	UserPhoneNotExisted:    "用户手机号不存在",
	UserAssetTransferEqual: "用户转移资产手机号不能跟源手机号一致",
}

// GetMsg get error information based on Code
func GetMsg(code Code) string {
	msg, ok := HttpStatusCodeMsgFlags[code]
	if ok {
		return msg
	}

	msg, ok = BusinessStatusCodeMsgFlags[code]
	if ok {
		return msg
	}

	return HttpStatusCodeMsgFlags[RequestErr]
}
