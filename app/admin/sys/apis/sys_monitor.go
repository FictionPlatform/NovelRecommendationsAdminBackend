package apis

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	"go-admin/core/lang"
	"go-admin/core/utils/fileutils"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/strutils"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	_ "go-admin/core/dto/response"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Monitor struct {
	api.Api
}

// GetMonitor admin-获取服务器信息
// @Summary 获取服务器信息
// @Description 获取服务器信息
// @Tags 系统监控管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-monitor [get]
func (e Monitor) GetMonitor(c *gin.Context) {
	e.MakeContext(c)

	sysInfo, err := host.Info()
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()
	osDic["ip"] = iputils.GetLocaHost()
	osDic["projectDir"] = fileutils.GetCurrentPath()
	osDic["hostName"] = sysInfo.Hostname
	osDic["time"] = time.Now().Format("2006-01-02 15:04:05")

	dis, err := disk.Usage("/")
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	memUsedMB := int(memInfo.Used) / GB
	memTotalMB := int(memInfo.Total) / GB
	memFreeMB := int(memInfo.Free) / GB
	memUsedPercent := int(memInfo.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuInfo, err := cpu.Info()
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	cpuDic["cpuInfo"] = cpuInfo
	percent, err := cpu.Percent(0, false)
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	if len(percent) > 0 {
		cpuDic["Percent"] = strutils.Round(percent[0], 2)
	} else {
		cpuDic["Percent"] = 0
	}
	cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	cpuDic["cpus"] = cpus
	cpuNum, err := cpu.Counts(false)
	if err != nil {
		e.Error(baseLang.DataQueryCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err).Error())
		return
	}
	cpuDic["cpuNum"] = cpuNum

	//服务器磁盘信息
	disklist := make([]disk.UsageStat, 0)
	//所有分区
	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				disklist = append(disklist, *diskDetail)
			}
		}
	}

	result := map[string]interface{}{
		"os":       osDic,
		"mem":      memDic,
		"cpu":      cpuDic,
		"disk":     diskDic,
		"diskList": disklist,
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Ping admin-ping测试
// @Summary ping测试
// @Description ping测试
// @Tags 系统监控管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-monitor/ping [get]
func (e Monitor) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
