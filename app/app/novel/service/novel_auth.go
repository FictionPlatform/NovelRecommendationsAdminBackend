package service

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"gorm.io/gorm"
)

type NovelAuth struct {
	service.Service
}

// NewNovelAuthService app-实例化读者注册登录服务
func NewNovelAuthService(s *service.Service) *NovelAuth {
	var srv = new(NovelAuth)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// LoginVerify app-读者登录验证（失败统一返回“用户名或密码错误”，防用户名枚举）
func (e *NovelAuth) LoginVerify(login *dto.NovelAuthLoginReq) (*userModels.User, int, error) {
	user := &userModels.User{}
	err := e.Orm.Where("user_name = ? AND status = ?", login.Username, global.SysStatusOk).First(user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("novel reader login query user error:%s", err.Error())
		}
		return nil, baseLang.NovelUserPwdErrCode, lang.MsgErr(baseLang.NovelUserPwdErrCode, e.Lang)
	}
	if !user.CheckPwd(login.Password) {
		return nil, baseLang.NovelUserPwdErrCode, lang.MsgErr(baseLang.NovelUserPwdErrCode, e.Lang)
	}
	return user, baseLang.SuccessCode, nil
}

// Register app-读者注册（自动创建读者扩展资料，注册成功即可登录）
func (e *NovelAuth) Register(req *dto.NovelAuthRegisterReq) (*userModels.User, int, error) {
	var count int64
	if err := e.Orm.Model(&userModels.User{}).Where("user_name = ?", req.Username).Count(&count).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if count > 0 {
		return nil, baseLang.NovelUsernameExistCode, lang.MsgErr(baseLang.NovelUsernameExistCode, e.Lang)
	}

	now := time.Now()
	user := &userModels.User{
		LevelId:   1,
		UserName:  req.Username,
		TrueName:  "- -",
		Money:     decimal.NewFromInt(0),
		Pwd:       req.Password,
		Status:    global.SysStatusOk,
		TreeLeaf:  global.SysStatusOk,
		CreateBy:  0,
		UpdateBy:  0,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := e.Orm.Create(user).Error
	if err != nil {
		return nil, baseLang.NovelRegisterFailCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.NovelRegisterFailCode, baseLang.DataInsertLogCode, err)
	}

	profile := &models.NovelReaderProfile{
		UserId:           user.Id,
		Nickname:         req.Username,
		Avatar:           req.Avatar,
		NotifyComment:    1,
		NotifyBookUpdate: 1,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}
	if err := e.Orm.Create(profile).Error; err != nil {
		e.Log.Errorf("novel reader profile create error:%s", err.Error())
		_ = e.Orm.Delete(&userModels.User{}, user.Id)
		return nil, baseLang.NovelRegisterFailCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.NovelRegisterFailCode, baseLang.DataInsertLogCode, err)
	}
	return user, baseLang.SuccessCode, nil
}
