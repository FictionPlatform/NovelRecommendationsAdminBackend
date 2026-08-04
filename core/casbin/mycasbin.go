package mycasbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"gorm.io/gorm"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB, _ string) *casbin.SyncedEnforcer {
	Apter, err := NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// L19：关闭 casbin 内置日志——Enforce 每请求打日志拖累性能，鉴权结果已由 AuthCheckRole 中间件记录
	e.EnableLog(false)
	return e
}

func LoadPolicy(c *gin.Context) (*casbin.SyncedEnforcer, error) {
	enforcer := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
	if enforcer == nil {
		return nil, fmt.Errorf("casbin enforcer not found for host %s", c.Request.Host)
	}
	if err := enforcer.LoadPolicy(); err != nil {
		log.Errorf("casbin rbac_model or policy init error, %s ", err.Error())
		return nil, err
	}
	return enforcer, nil
}

// GetGlobalEnforcer 获取全局casbin enforcer（优先通配键，单实例部署场景）
func GetGlobalEnforcer() *casbin.SyncedEnforcer {
	if e := runtime.RuntimeConfig.GetCasbinKey("*"); e != nil {
		return e
	}
	for _, e := range runtime.RuntimeConfig.GetCasbin() {
		return e
	}
	return nil
}
