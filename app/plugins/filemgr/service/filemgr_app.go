package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/plugins/filemgr/models"
	"go-admin/app/plugins/filemgr/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/ossutils"
	"mime/multipart"
	"net/url"
	"path"
	"strings"

	"go-admin/core/utils/excelutils"
	"gorm.io/gorm"
	"time"
)

type FilemgrApp struct {
	service.Service
}

// maxAppUploadSize APP 安装包上传大小上限 200MB
const maxAppUploadSize = 200 << 20

// allowedAppExt 允许上传的 APP 安装包扩展名白名单
var allowedAppExt = map[string]bool{".apk": true, ".ipa": true, ".zip": true}

// validDownloadScheme 校验下载根地址仅允许 http/https 协议，
// 拒绝 javascript:、data: 等可执行/伪协议拼接进下载链接
func validDownloadScheme(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil || u.Scheme == "" {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// NewFilemgrAppService plugins-实例化APP管理
func NewFilemgrAppService(s *service.Service) *FilemgrApp {
	var srv = new(FilemgrApp)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage plugins-获取APP管理分页列表
func (e *FilemgrApp) GetPage(c *dto.FilemgrAppQueryReq, p *middleware.DataPermission) ([]models.FilemgrApp, int64, int, error) {
	var data models.FilemgrApp
	var list []models.FilemgrApp
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	for index, item := range list {
		if item.DownloadType == constant.AppDownloadTypeOss {
			url := ""
			url, err = e.generateAppOssUrl(&item)
			if err == nil {
				item.DownloadUrl = url
			}
		}
		list[index] = item
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get plugins-获取APP管理详情
func (e *FilemgrApp) Get(id int64, p *middleware.DataPermission) (*models.FilemgrApp, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.FilemgrApp{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne plugins-获取APP管理一条记录
func (e *FilemgrApp) QueryOne(queryCondition *dto.FilemgrAppQueryReq, p *middleware.DataPermission) (*models.FilemgrApp, int, error) {
	data := &models.FilemgrApp{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取APP管理数据总数
func (e *FilemgrApp) Count(queryCondition *dto.FilemgrAppQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.FilemgrApp{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert plugins-新增APP管理
func (e *FilemgrApp) Insert(c *dto.FilemgrAppInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Platform == "" {
		return 0, baseLang.AppPlatformEmptyCode, lang.MsgErr(baseLang.AppPlatformEmptyCode, e.Lang)
	}
	if c.Version == "" {
		return 0, baseLang.AppVersionEmptyCode, lang.MsgErr(baseLang.AppVersionEmptyCode, e.Lang)
	}
	if c.DownloadType == "" {
		return 0, baseLang.AppDownloadTypeEmptyCode, lang.MsgErr(baseLang.AppDownloadTypeEmptyCode, e.Lang)
	}
	if c.DownloadType == constant.AppDownloadTypeOss || c.DownloadType == constant.AppDownloadTypeLocal {
		if c.AppType == "" {
			return 0, baseLang.AppTypeCode, lang.MsgErr(baseLang.AppTypeCode, e.Lang)
		}
		if c.LocalAddress == "" {
			return 0, baseLang.AppUploadCode, lang.MsgErr(baseLang.AppUploadCode, e.Lang)
		}
	}
	if c.Remark == "" {
		return 0, baseLang.AppRemarkCode, lang.MsgErr(baseLang.AppRemarkCode, e.Lang)
	}

	query := dto.FilemgrAppQueryReq{}
	query.Platform = c.Platform
	query.AppType = c.AppType
	query.Version = c.Version
	count, respCode, err := e.Count(&query)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.AppExistCode, lang.MsgErr(baseLang.AppExistCode, e.Lang)
	}

	//oss上传
	if c.DownloadType == constant.AppDownloadTypeOss {
		err = e.uploadOssFile(c.AppType, c.Version, c.Platform, c.LocalAddress)
		if err != nil {
			return 0, baseLang.AppOssUploadLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.AppUploadCode, baseLang.AppOssUploadLogCode, err)
		}
	}
	if c.DownloadType == constant.AppDownloadTypeLocal {
		if c.LocalRootUrl == "" {
			return 0, baseLang.AppLocalUrlEmptyCode, lang.MsgErr(baseLang.AppLocalUrlEmptyCode, e.Lang)
		}
		//L25：仅允许 http/https 协议，拒绝 javascript: 等可执行协议注入下载链接
		if !validDownloadScheme(c.LocalRootUrl) {
			return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
		}
		//c.LocalAddress = strings.Replace(c.LocalAddress, config.ApplicationConfig.FileRootPath, "", -1)
		c.DownloadUrl = c.LocalRootUrl + c.LocalAddress
	}

	now := time.Now()
	var data models.FilemgrApp
	data.Version = c.Version
	data.Platform = c.Platform
	data.AppType = c.AppType
	data.LocalAddress = c.LocalAddress
	data.DownloadType = c.DownloadType
	data.DownloadUrl = c.DownloadUrl
	data.Remark = c.Remark
	data.Status = constant.AppPublishWait
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update plugins-更新APP管理
func (e *FilemgrApp) Update(c *dto.FilemgrAppUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete plugins-删除APP管理
func (e *FilemgrApp) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	//L25：先统计需清理的 OSS 资源（不删除），DB 删除成功后再清理，
	//避免 DB 失败时 OSS 文件已丢导致记录悬空、下载链接失效
	var pendingOssObjectKeys []string
	for _, id := range ids {
		result, respCode, err := e.Get(id, p)
		if respCode != baseLang.DataNotFoundCode && err != nil {
			return respCode, err
		}

		//同一个完全相同的版本，可能因为网路有多条记录，但这些记录都指向一个oss资源，此时只有最后一条记录，才能删除oss资源
		query := dto.FilemgrAppQueryReq{}
		query.Platform = result.Platform
		query.AppType = result.AppType
		query.Version = result.Version
		var count int64
		count, respCode, err = e.Count(&query)
		if respCode != baseLang.DataNotFoundCode && err != nil {
			return respCode, err
		}
		if count <= 1 {
			objectKey, _ := e.generateAppOssObjectKey(result)
			if objectKey != "" {
				pendingOssObjectKeys = append(pendingOssObjectKeys, objectKey)
			}
		}
	}
	var data models.FilemgrApp
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	//DB 删除成功后清理 OSS 资源（删除失败仅告警，不阻断主流程）
	oss, _ := e.getOssClient()
	if oss != nil {
		for _, objectKey := range pendingOssObjectKeys {
			if err := oss.Bucket.DeleteObject(objectKey, nil); err != nil {
				e.Log.Warnf("delete oss object %s error:%s", objectKey, err.Error())
			}
		}
	}
	return baseLang.SuccessCode, nil
}

// GetSingleUploadFileInfo admin-获取APP管理单个上传文件信息
func (e *FilemgrApp) GetSingleUploadFileInfo(form *multipart.Form, file *multipart.FileHeader, dst *string) (int, error) {
	if len(form.File) != 1 {
		return baseLang.AppSelectOneFileUploadCode, lang.MsgErr(baseLang.AppSelectOneFileUploadCode, e.Lang)
	}
	for _, files := range form.File {
		if len(files) != 1 {
			return baseLang.AppSelectOneFileUploadCode, lang.MsgErr(baseLang.AppSelectOneFileUploadCode, e.Lang)
		}
		for _, item := range files {
			// 扩展名白名单校验
			ext := strings.ToLower(path.Ext(item.Filename))
			if !allowedAppExt[ext] {
				return baseLang.AppUploadCode, lang.MsgErr(baseLang.AppUploadCode, e.Lang)
			}
			// 文件大小校验
			if item.Size <= 0 || item.Size > maxAppUploadSize {
				return baseLang.AppUploadCode, lang.MsgErr(baseLang.AppUploadCode, e.Lang)
			}
			*dst = config.ApplicationConfig.FileRootPath + "app/" + idgen.UUID() + ext
			*file = *item
			return baseLang.SuccessCode, nil
		}
	}
	return baseLang.SuccessCode, nil
}

// Export plugins-导出APP管理
func (e *FilemgrApp) Export(list []models.FilemgrApp) ([]byte, error) {
	sheetName := "FilemgrApp"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "J", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"App编号", "版本号", "系统平台", "App类型", "下载类型", "发布状态", "下载地址", "服务器本地地址",
		"更新内容", "创建时间"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		platform := dictService.GetLabel("plugin_filemgr_app_platform", item.Platform)              //平台
		appType := dictService.GetLabel("plugin_filemgr_app_type", item.AppType)                    //app类型
		downloadType := dictService.GetLabel("plugin_filemgr_app_download_type", item.DownloadType) //下载类型
		publishStatus := dictService.GetLabel("plugin_filemgr_publish_status", item.Status)         //下载类型
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, excelutils.SafeRow(
			item.Id, item.Version, platform, appType, downloadType, publishStatus, item.DownloadUrl, item.Remark, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		))
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// uploadOssFile plugins-内部方法，上传文件到oss
func (e *FilemgrApp) uploadOssFile(appType, version, platform, localAddress string) error {
	app := models.FilemgrApp{}
	app.AppType = appType
	app.Version = version
	app.Platform = platform
	key, err := e.generateAppOssObjectKey(&app)
	if err != nil {
		return err
	}
	client, err := e.getOssClient()
	if err != nil {
		return err
	}
	err = client.UploadWithSpace(key, localAddress)
	if err != nil {
		e.Log.Errorf("上传失败，失败原因:%s \r\n", err)
		return err
	}
	return nil
}

// generateAppOssUrl plugins-内部方法，获取APP管理的下载链接
func (e *FilemgrApp) generateAppOssUrl(App *models.FilemgrApp) (string, error) {
	appPath, err := e.generateAppOssObjectKey(App)
	if err != nil {
		return "", err
	}
	oss, err := e.getOssClient()
	if err != nil {
		return "", err
	}
	return oss.GeneratePresignedUrl(appPath)
}

// getOssClient plugins-内部方法，APP管理获取oss客户端
func (e *FilemgrApp) getOssClient() (*ossutils.ALiYunOSS, error) {
	var sysConfService = adminService.NewSysConfigService(&e.Service)
	endPoint, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_endpoint")
	key, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_access_key_id")
	secret, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_access_key_secret")
	bucketName, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_bucket")
	if endPoint == "" || key == "" || secret == "" || bucketName == "" {
		return nil, errors.New("oss param config empty")
	}
	oss := ossutils.ALiYunOSS{}
	err := oss.InitOssClient(key, secret, endPoint, bucketName)
	if err != nil {
		return nil, err
	}
	return &oss, nil
}

// generateAppOssObjectKey plugins-内部方法，生成oss key
func (e *FilemgrApp) generateAppOssObjectKey(App *models.FilemgrApp) (string, error) {
	var sysConfService = adminService.NewSysConfigService(&e.Service)
	var dictDataService = adminService.NewSysDictDataService(&e.Service)

	//app目录
	appPath, _, err := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_root_path")
	if err != nil {
		return "", errors.New(fmt.Sprintf("oss config err: %s", err))
	}
	appPath += dictDataService.GetLabel("plugin_filemgr_app_type", App.AppType) + "_"
	appPath += App.Version

	switch App.Platform {
	case constant.AppPlatformAndroid:
		appPath += ".apk"
	case constant.AppPlatformIOS:
		appPath += ".ipa"
	default:
		return "", errors.New("app platform error")
	}
	return appPath, nil
}
