// Package app
// @Description: 多语言消息管理，根据业务需要新增消息内容，标识内容勿操作
package lang

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitxx/logger/logbase"
)

// MsgByCode
// @Description: i18n
// @param errCode query int true "错误码"
// @param lang query string true "语言"
// @return string
func MsgByCode(errCode int, lang string) string {
	switch lang {
	case "en":
		return EnLang.T(MsgInfo[errCode])
	default:
		return MsgInfo[errCode]
	}
}

// MsgByValue
// @Description: 直接根据值返回对应语言
// @param value query string true "值"
// @param lang query string true "语言"
// @return string
func MsgByValue(value string, lang string) string {
	switch lang {
	case "en":
		return EnLang.T(value)
	default:
		return value
	}
}

// MsgErr
// @Description: 获取error
// @param errCode query int true "错误码"
// @param lang query string true "语言"
// @return error
func MsgErr(errCode int, lang string) error {
	return errors.New(MsgByCode(errCode, lang))
}

// MsgErrf
// @Description:
// @param errCode query int true "错误码"
// @param lang query string true "语言"
// @param f query object false "格式化参数"
// @return error
func MsgErrf(errCode int, lang string, f ...interface{}) error {
	return fmt.Errorf(MsgByCode(errCode, lang), f)
}

// MsgLogErrf
// @Description: 带有参数，有些底层消息不应当被使用者感知，该类消息记录在日志中，并返回应用层可理解的消息
// @param log query object true "日志句柄"
// @param lang query string true "语言"
// @param errCodeReplace query int true "对外返回的消息码"
// @param errCode query int true "真实消息码"
// @param f query object false "格式化参数"
// @return error
func MsgLogErrf(log *logbase.Helper, lang string, errCodeReplace, errCode int, f ...interface{}) error {
	err := MsgErrf(errCode, lang, f)
	log.Error(err)
	if errCodeReplace <= 0 || errCodeReplace == errCode {
		return err
	}
	return MsgErr(errCodeReplace, lang)
}

// MsgLogErr
// @Description: 无参数，有些底层消息不应当被使用者感知，该类消息记录在日志中，并返回应用层可理解的消息
// @param log query object true "日志句柄"
// @param lang query string true "语言"
// @param errCodeReplace query int true "对外返回的消息码"
// @param errCode query int true "真实消息码"
// @return error
func MsgLogErr(log *logbase.Helper, lang string, errCodeReplace, errCode int) error {
	err := MsgErr(errCode, lang)
	log.Error(err)
	if errCodeReplace <= 0 || errCodeReplace == errCode {
		return err
	}
	return MsgErr(errCodeReplace, lang)
}

// TranslationText
// @Description: 仅支持 - 分隔符
// @param l query string true "语言"
// @param text query string true "文本"
// @return string
func TranslationText(l string, text string) string {
	values := strings.Split(text, "-")
	if len(values) <= 0 {
		return text
	}
	newValue := MsgByValue(values[0], l)
	return strings.Replace(text, values[0], newValue, 1)
}
