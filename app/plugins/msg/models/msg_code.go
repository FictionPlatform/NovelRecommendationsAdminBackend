package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MsgCode struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
	Code      string     `json:"-" gorm:"column:code;type:varchar(128);comment:验证码(bcrypt哈希)"`
	CodeType  string     `json:"codeType" gorm:"column:code_type;type:char(1);comment:验证码类型 1-邮箱；2-短信"`
	Remark    string     `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注异常"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:验证码状态 1-发送成功 2-发送失败"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (MsgCode) TableName() string {
	return "plugins_msg_code"
}

func (e *MsgCode) BeforeCreate(_ *gorm.DB) error {
	return e.HashCode()
}

func (e *MsgCode) BeforeUpdate(_ *gorm.DB) error {
	return e.HashCode()
}

// HashCode 验证码以 bcrypt 哈希落库，禁止明文存储
func (e *MsgCode) HashCode() error {
	if e.Code == "" || isBcryptHash(e.Code) {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(e.Code), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	e.Code = string(hash)
	return nil
}

// CheckCode 校验验证码（哈希比对）
func (e *MsgCode) CheckCode(code string) bool {
	if e.Code == "" || code == "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(e.Code), []byte(code)) == nil
}

func isBcryptHash(s string) bool {
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}
