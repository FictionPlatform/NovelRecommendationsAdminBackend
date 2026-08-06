package lang

import "go-admin/core/lang"

const (
	SuccessCode   = 200
	BadRequestErr = 400
	AuthErr       = 401
	ForbitErr     = 403
	NotFoundErr   = 404
	ServerErr     = 500

	ParamErrCode         = 1001
	OpErrCode            = 1002
	DataDecodeCode       = 1013
	DataDecodeLogCode    = 1012
	DataQueryCode        = 1003
	DataQueryLogCode     = 1004
	DataInsertLogCode    = 1005
	DataInsertCode       = 1006
	DataNotUpdateCode    = 1014
	DataUpdateCode       = 1007
	DataUpdateLogCode    = 1008
	DataDeleteCode       = 1009
	DataDeleteLogCode    = 1010
	DataNotFoundCode     = 1011
	ServerErrLogCode     = 1012
	BadRequestErrLogCode = 1013
	AuthErrLogCode       = 1014

	// 通用安全/平台类错误码（多用于中间件，响应层按 response.Error 规则透传）
	AdminOnlyErrCode       = 1015
	RateLimitServerErrCode = 1016
	RateLimitErrCode       = 1017
	IpBlacklistCode        = 1018
	DbConnFailCode         = 1019
	PermissionErrCode      = 1020
	LoginFailCode          = 1021
)

func init() {
	//1-基础通用
	lang.MsgInfo[SuccessCode] = "操作成功"
	lang.MsgInfo[BadRequestErr] = "错误请求"
	lang.MsgInfo[AuthErr] = "状态失效，请重新登录"
	lang.MsgInfo[ForbitErr] = "对不起，您权限不足，操作异常，请联系管理员"
	lang.MsgInfo[NotFoundErr] = "对不起，资源未找到"
	lang.MsgInfo[ServerErr] = "内部错误"
	lang.MsgInfo[ParamErrCode] = "参数错误"
	lang.MsgInfo[OpErrCode] = "操作异常，请检查"
	lang.MsgInfo[DataDecodeCode] = "数据解析异常"
	lang.MsgInfo[DataDecodeLogCode] = "数据解析异常：%s"
	lang.MsgInfo[DataQueryCode] = "数据查询失败"
	lang.MsgInfo[DataQueryLogCode] = "数据查询失败：%s"
	lang.MsgInfo[DataInsertLogCode] = "数据新增失败：%s"
	lang.MsgInfo[DataInsertCode] = "数据新增失败"
	lang.MsgInfo[DataNotUpdateCode] = "数据未变更"
	lang.MsgInfo[DataUpdateCode] = "数据更新异常"
	lang.MsgInfo[DataUpdateLogCode] = "数据更新异常：%s"
	lang.MsgInfo[DataDeleteCode] = "数据删除失败"
	lang.MsgInfo[DataDeleteLogCode] = "数据删除失败：%s"
	lang.MsgInfo[DataNotFoundCode] = "数据不存在"
	lang.MsgInfo[ServerErrLogCode] = "内部错误：%s"
	lang.MsgInfo[BadRequestErrLogCode] = "错误请求：%s"
	lang.MsgInfo[AuthErrLogCode] = "认证失败：%s"
	lang.MsgInfo[AdminOnlyErrCode] = "仅超级管理员可访问"
	lang.MsgInfo[RateLimitServerErrCode] = "限流服务异常"
	lang.MsgInfo[RateLimitErrCode] = "请求过于频繁，请稍后再试"
	lang.MsgInfo[IpBlacklistCode] = "禁止访问"
	lang.MsgInfo[DbConnFailCode] = "数据库连接获取失败"
	lang.MsgInfo[PermissionErrCode] = "数据权限处理失败，请联系管理员"
	lang.MsgInfo[LoginFailCode] = "账号或密码错误"
}
