package dto

type RegUserRequest struct {
	Phone         int64  `json:"phone" form:"phone" binding:"required"`
	Passwd        string `json:"passwd" form:"passwd" binding:"required"`
	PasswdConfirm string `json:"passwd_confirm" form:"passwd_confirm" binding:"required"`
	Captcha
}

type UserRequest struct {
	Phone  int64  `json:"phone" form:"phone" binding:"required"`
	Passwd string `json:"passwd" form:"passwd" binding:"required"`
	Code   string `json:"code" form:"code" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type Captcha struct {
	Id   string `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
}

type UserInfo struct {
	UserId int64
	Phone  int64
}
