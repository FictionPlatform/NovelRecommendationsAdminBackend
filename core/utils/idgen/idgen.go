package idgen

import (
	"crypto/rand"
	"math/big"

	"github.com/google/uuid"
)

const inviteChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func UUID() string {
	return uuid.New().String()
}

// InviteId 生成6位邀请码（crypto/rand 不可预测；原实现 base64Captcha.RandText 基于 math/rand 时间种子，可预测撞码）
func InviteId() string {
	b := make([]byte, 6)
	clen := big.NewInt(int64(len(inviteChars)))
	for i := range b {
		n, err := rand.Int(rand.Reader, clen)
		if err != nil {
			// crypto/rand 失败为致命错误（熵源不可用），直接暴露
			panic("idgen: crypto/rand failed: " + err.Error())
		}
		b[i] = inviteChars[n.Int64()]
	}
	return string(b)
}
