package models

import (
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	LevelId     int64           `json:"levelId" gorm:"column:level_id;type:int;comment:用户等级编号"`
	UserName    string          `json:"userName" gorm:"column:user_name;type:varchar(100);comment:用户昵称"`
	TrueName    string          `json:"trueName" gorm:"column:true_name;type:varchar(100);comment:真实姓名"`
	Money       decimal.Decimal `json:"money" gorm:"column:money;type:decimal(30,18);comment:余额"`
	Email       string          `json:"email" gorm:"column:email;type:varchar(300);comment:电子邮箱"`
	MobileTitle string          `json:"mobileTitle" gorm:"column:mobile_title;type:varchar(255);comment:用户手机号国家前缀"`
	Mobile      string          `json:"mobile" gorm:"column:mobile;type:varchar(100);comment:手机号码"`
	Avatar      string          `json:"avatar" gorm:"column:avatar;type:varchar(1000);comment:头像路径"`
	PayPwd      string          `json:"-" gorm:"column:pay_pwd;type:varchar(100);comment:提现密码"`
	Pwd         string          `json:"-" gorm:"column:pwd;type:varchar(100);comment:登录密码"`
	RefCode     string          `json:"refCode" gorm:"column:ref_code;type:varchar(255);comment:推荐码"`
	ParentId    int64           `json:"parentId" gorm:"column:parent_id;type:int;comment:父级编号"`
	ParentIds   string          `json:"parentIds" gorm:"column:parent_ids;type:varchar(1000);comment:所有父级编号"`
	TreeSort    int64           `json:"treeSort" gorm:"column:tree_sort;type:decimal(10,0);comment:本级排序号（升序）"`
	TreeSorts   string          `json:"treeSorts" gorm:"column:tree_sorts;type:varchar(1000);comment:所有级别排序号"`
	TreeLeaf    string          `json:"treeLeaf" gorm:"column:tree_leaf;type:char(1);comment:是否最末级"`
	TreeLevel   int64           `json:"treeLevel" gorm:"column:tree_level;type:int;comment:层次级别"`
	Status      string          `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-异常)"`
	Remark      string          `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	CreateBy    int64           `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy    int64           `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt   *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`

	// 扩展
	UserLevel      *UserLevel `json:"userLevel" gorm:"foreignkey:level_id"`
	ParentUserName string     `json:"parentUserName" gorm:"-"`
	ParentRefCode  string     `json:"parentRefCode" gorm:"-"`
}

func (User) TableName() string {
	return "app_user"
}

// isBcryptHash 判断是否已是 bcrypt 哈希，避免重复加密
func isBcryptHash(s string) bool {
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}

// Encrypt 对登录密码/提现密码做 bcrypt 哈希
func (e *User) Encrypt() (err error) {
	if e.Pwd != "" && !isBcryptHash(e.Pwd) {
		var hash []byte
		if hash, err = bcrypt.GenerateFromPassword([]byte(e.Pwd), bcrypt.DefaultCost); err != nil {
			return
		}
		e.Pwd = string(hash)
	}
	if e.PayPwd != "" && !isBcryptHash(e.PayPwd) {
		var hash []byte
		if hash, err = bcrypt.GenerateFromPassword([]byte(e.PayPwd), bcrypt.DefaultCost); err != nil {
			return
		}
		e.PayPwd = string(hash)
	}
	return
}

func (e *User) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *User) BeforeUpdate(_ *gorm.DB) error {
	return e.Encrypt()
}

// CheckPwd 校验登录密码
func (e *User) CheckPwd(pwd string) bool {
	if e.Pwd == "" {
		return false
	}
	if !isBcryptHash(e.Pwd) {
		return e.Pwd == pwd
	}
	return bcrypt.CompareHashAndPassword([]byte(e.Pwd), []byte(pwd)) == nil
}

// CheckPayPwd 校验提现密码
func (e *User) CheckPayPwd(payPwd string) bool {
	if e.PayPwd == "" {
		return false
	}
	if !isBcryptHash(e.PayPwd) {
		return e.PayPwd == payPwd
	}
	return bcrypt.CompareHashAndPassword([]byte(e.PayPwd), []byte(payPwd)) == nil
}
