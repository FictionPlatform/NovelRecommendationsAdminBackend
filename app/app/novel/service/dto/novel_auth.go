package dto

// NovelAuthLoginReq app-读者登录请求
type NovelAuthLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// NovelAuthRegisterReq app-读者注册请求
type NovelAuthRegisterReq struct {
	Username string `json:"username" binding:"required,min=2,max=64"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Avatar   string `json:"avatar"`
}
