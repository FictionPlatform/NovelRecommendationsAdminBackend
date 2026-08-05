package service

import (
	"encoding/json"
	"errors"

	"go-admin/app/app/novel/models"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"gorm.io/gorm"
)

// getUserSnapshot 读取读者昵称/头像快照（app_user 复用）
func getUserSnapshot(e *service.Service, userId int64) (string, string, int, error) {
	u := &userModels.User{}
	err := e.Orm.First(u, userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return u.UserName, u.Avatar, baseLang.SuccessCode, nil
}

// fillShelfBooks 批量填充书架书本信息
func fillShelfBooks(e *service.Service, list []models.NovelBookshelf) {
	if len(list) == 0 {
		return
	}
	bookIds := make([]int64, 0, len(list))
	for _, s := range list {
		bookIds = append(bookIds, s.BookId)
	}
	var books []models.NovelBook
	err := e.Orm.Where("id in (?)", bookIds).Find(&books).Error
	if err != nil {
		return
	}
	bookMap := map[int64]models.NovelBook{}
	for _, b := range books {
		if b.Tags != "" {
			_ = json.Unmarshal([]byte(b.Tags), &b.TagList)
		}
		bookMap[b.Id] = b
	}
	for i := range list {
		if b, ok := bookMap[list[i].BookId]; ok {
			list[i].Book = &b
		}
	}
}
