package dto

type NovelProfileUpdateReq struct {
	Nickname            string   `json:"nickname" comment:"昵称"`
	Avatar              string   `json:"avatar" comment:"头像"`
	Bio                 string   `json:"bio" comment:"简介"`
	Email               string   `json:"email" comment:"邮箱"`
	PreferredCategories []string `json:"preferredCategories" comment:"偏好分类"`
	NotifyComment       int      `json:"notifyComment" comment:"评论消息通知 0|1"`
	NotifyBookUpdate    int      `json:"notifyBookUpdate" comment:"书籍更新通知 0|1"`
	CurrUserId          int64    `json:"-" comment:"当前登录用户"`
}
