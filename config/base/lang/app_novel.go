package lang

import "go-admin/core/lang"

const (
	NovelBookNotExistCode         = 41000
	NovelBookTitleEmptyCode       = 41001
	NovelBookTitleExistCode       = 41002
	NovelReviewRatingRangeCode    = 41003
	NovelPostNotExistCode         = 41005
	NovelPostContentLenCode       = 41006
	NovelPostTitleEmptyCode       = 41007
	NovelPostCommentNotExistCode  = 41008
	NovelShelfNotExistCode        = 41009
	NovelNoPermissionCode         = 41010
	NovelActionTypeErrCode        = 41012
	NovelPostRefBookNotExistCode  = 41013
	NovelShelfAlreadyExistCode    = 41014
	NovelBookCategoryEmptyCode    = 41016
	NovelParentNotSamePostCode    = 41017
	NovelBookTitleTooLongCode     = 41018
	NovelBookAuthorTooLongCode    = 41019
	NovelBookSloganTooLongCode    = 41020
	NovelBookCoverTooLongCode     = 41021
	NovelBookDescTooLongCode      = 41022
	NovelPostTitleTooLongCode     = 41023
	NovelContentTooLongCode       = 41024
	NovelNicknameTooLongCode      = 41025
	NovelBioTooLongCode           = 41026
	NovelEmailTooLongCode         = 41027
	NovelPostCommentEmptyCode     = 41028
	NovelFollowSelfCode           = 41029
	NovelFollowAlreadyExistCode   = 41030
	NovelFollowNotExistCode       = 41031
	NovelUserNotExistCode         = 41032
	NovelUserPwdErrCode           = 41033
	NovelUsernameExistCode        = 41034
	NovelRegisterFailCode         = 41035
	NovelFeedbackEmptyCode        = 41036
	NovelFeedbackTooLongCode      = 41037
	NovelFeedbackTypeErrCode      = 41038
	NovelNoticeTitleEmptyCode     = 41039
	NovelNoticeContentEmptyCode   = 41040
	NovelNoticeTitleTooLongCode   = 41041
	NovelNoticeContentTooLongCode = 41042
	NovelPostBannedCode           = 41043
	NovelUserStatusErrCode        = 41044
	NovelBanTimeErrCode           = 41045
	NovelPostStatusErrCode        = 41046
	NovelUserCancelledCode        = 41047
	NovelUserPwdWrongCode         = 41048
	NovelRegLimitCode             = 41049
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[NovelBookNotExistCode] = "书籍不存在"
	lang.MsgInfo[NovelBookTitleEmptyCode] = "书名不得为空"
	lang.MsgInfo[NovelBookTitleExistCode] = "已存在同名书籍"
	lang.MsgInfo[NovelReviewRatingRangeCode] = "评分需在 1~5 之间"
	lang.MsgInfo[NovelPostNotExistCode] = "话题不存在或已删除"
	lang.MsgInfo[NovelPostContentLenCode] = "话题正文需在 5000~10000 字之间"
	lang.MsgInfo[NovelPostTitleEmptyCode] = "话题标题不得为空"
	lang.MsgInfo[NovelPostCommentNotExistCode] = "评论不存在"
	lang.MsgInfo[NovelShelfNotExistCode] = "书架记录不存在"
	lang.MsgInfo[NovelNoPermissionCode] = "无权操作该数据"
	lang.MsgInfo[NovelActionTypeErrCode] = "不支持该互动类型"
	lang.MsgInfo[NovelPostRefBookNotExistCode] = "关联小说不存在"
	lang.MsgInfo[NovelShelfAlreadyExistCode] = "该书已在书架上"
	lang.MsgInfo[NovelBookCategoryEmptyCode] = "书籍分类不得为空"
	lang.MsgInfo[NovelParentNotSamePostCode] = "父评论不属于该话题"
	lang.MsgInfo[NovelBookTitleTooLongCode] = "书名不能超过 100 字"
	lang.MsgInfo[NovelBookAuthorTooLongCode] = "作者名不能超过 64 字"
	lang.MsgInfo[NovelBookSloganTooLongCode] = "推荐语不能超过 30 字"
	lang.MsgInfo[NovelBookCoverTooLongCode] = "封面地址不能超过 500 字"
	lang.MsgInfo[NovelBookDescTooLongCode] = "简介不能超过 500 字"
	lang.MsgInfo[NovelPostTitleTooLongCode] = "话题标题不能超过 200 字"
	lang.MsgInfo[NovelContentTooLongCode] = "评论/书评内容不能超过 1000 字"
	lang.MsgInfo[NovelNicknameTooLongCode] = "昵称不能超过 64 字"
	lang.MsgInfo[NovelBioTooLongCode] = "简介不能超过 500 字"
	lang.MsgInfo[NovelEmailTooLongCode] = "邮箱不能超过 128 字"
	lang.MsgInfo[NovelPostCommentEmptyCode] = "评论内容不得为空"
	lang.MsgInfo[NovelFollowSelfCode] = "不能关注自己"
	lang.MsgInfo[NovelFollowAlreadyExistCode] = "已关注该用户"
	lang.MsgInfo[NovelFollowNotExistCode] = "尚未关注该用户"
	lang.MsgInfo[NovelUserNotExistCode] = "用户不存在"
	lang.MsgInfo[NovelUserPwdErrCode] = "用户名或密码错误"
	lang.MsgInfo[NovelUsernameExistCode] = "用户名已存在"
	lang.MsgInfo[NovelRegisterFailCode] = "注册失败，请稍后重试"
	lang.MsgInfo[NovelFeedbackEmptyCode] = "反馈内容不得为空"
	lang.MsgInfo[NovelFeedbackTooLongCode] = "反馈内容不能超过 500 字"
	lang.MsgInfo[NovelFeedbackTypeErrCode] = "反馈类型无效"
	lang.MsgInfo[NovelNoticeTitleEmptyCode] = "公告标题不得为空"
	lang.MsgInfo[NovelNoticeContentEmptyCode] = "公告内容不得为空"
	lang.MsgInfo[NovelNoticeTitleTooLongCode] = "公告标题不能超过 100 字"
	lang.MsgInfo[NovelNoticeContentTooLongCode] = "公告内容不能超过 2000 字"
	lang.MsgInfo[NovelPostBannedCode] = "您已被禁止发帖，请于解禁时间后再试"
	lang.MsgInfo[NovelUserStatusErrCode] = "用户状态无效"
	lang.MsgInfo[NovelBanTimeErrCode] = "禁言截止时间必须晚于当前时间"
	lang.MsgInfo[NovelPostStatusErrCode] = "话题状态无效"
	lang.MsgInfo[NovelUserCancelledCode] = "该账号已注销"
	lang.MsgInfo[NovelUserPwdWrongCode] = "密码错误"
	lang.MsgInfo[NovelRegLimitCode] = "今日注册次数已达上限，请明天再试"
}
