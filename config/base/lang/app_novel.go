package lang

import "go-admin/core/lang"

const (
	NovelBookNotExistCode        = 41000
	NovelBookTitleEmptyCode      = 41001
	NovelBookTitleExistCode      = 41002
	NovelReviewRatingRangeCode   = 41003
	NovelReviewNotExistCode      = 41004
	NovelPostNotExistCode        = 41005
	NovelPostContentLenCode      = 41006
	NovelPostTitleEmptyCode      = 41007
	NovelPostCommentNotExistCode = 41008
	NovelShelfNotExistCode       = 41009
	NovelNoPermissionCode        = 41010
	NovelProfileNotExistCode     = 41011
	NovelActionTypeErrCode       = 41012
	NovelPostRefBookNotExistCode = 41013
	NovelShelfAlreadyExistCode   = 41014
	NovelCommentDeleteErrCode    = 41015
	NovelBookCategoryEmptyCode   = 41016
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[NovelBookNotExistCode] = "书籍不存在"
	lang.MsgInfo[NovelBookTitleEmptyCode] = "书名不得为空"
	lang.MsgInfo[NovelBookTitleExistCode] = "已存在同名书籍"
	lang.MsgInfo[NovelReviewRatingRangeCode] = "评分需在 1~5 之间"
	lang.MsgInfo[NovelReviewNotExistCode] = "书评不存在"
	lang.MsgInfo[NovelPostNotExistCode] = "帖子不存在或已删除"
	lang.MsgInfo[NovelPostContentLenCode] = "帖子正文需在 5000~10000 字之间"
	lang.MsgInfo[NovelPostTitleEmptyCode] = "帖子标题不得为空"
	lang.MsgInfo[NovelPostCommentNotExistCode] = "评论不存在"
	lang.MsgInfo[NovelShelfNotExistCode] = "书架记录不存在"
	lang.MsgInfo[NovelNoPermissionCode] = "无权操作该数据"
	lang.MsgInfo[NovelProfileNotExistCode] = "读者资料不存在"
	lang.MsgInfo[NovelActionTypeErrCode] = "不支持该互动类型"
	lang.MsgInfo[NovelPostRefBookNotExistCode] = "关联小说不存在"
	lang.MsgInfo[NovelShelfAlreadyExistCode] = "该书已在书架上"
	lang.MsgInfo[NovelCommentDeleteErrCode] = "评论删除失败"
	lang.MsgInfo[NovelBookCategoryEmptyCode] = "书籍分类不得为空"
}
