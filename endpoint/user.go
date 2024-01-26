package endpoint

import (
	"net/http"
	"novel/env"
	go_micro_service_user "novel/proto/user"
	go_micro_service_wallet "novel/proto/wallet"
	selfwrappers "novel/wrappers"
	"strconv"
	"time"

	"novel/dto"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/scg130/tools"
	"github.com/scg130/tools/wrappers"
	"github.com/sirupsen/logrus"
)

type User struct {
	UserCli   go_micro_service_user.UserCenterService
	WalletCli go_micro_service_wallet.WalletService
}

var userSrv *User

func NewUserSrv() *User {
	if userSrv == nil {
		userSrv = &User{
			UserCli: go_micro_service_user.NewUserCenterService("go.micro.service.user", tools.GetMicroClient(
				"go.micro.service.user",
				wrappers.NewTracerWrapper(),
				wrappers.NewRateLimitClientWrapper(100),
				selfwrappers.NewHystrixWrapper(),
			)),
			WalletCli: go_micro_service_wallet.NewWalletService("go.micro.service.wallet", tools.GetMicroClient(
				"go.micro.service.wallet",
				wrappers.NewTracerWrapper(),
				wrappers.NewRateLimitClientWrapper(100),
				selfwrappers.NewHystrixWrapper(),
			)),
		}
	}
	return userSrv
}

func (u *User) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", 0, "/", "/", false, false)
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
	})
	return
}

// @Summary 登录
// @Description 登录
// @Tags 用户中心
// @Produce json
// @Param body body dto.UserRequest true "body参数"
// @Success 200 {object}  dto.Resp{data=dto.LoginResp}
// @Router /login [post]
func (u *User) Login(ctx *gin.Context) {
	var req dto.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	resp, err := u.UserCli.Login(ctx, &go_micro_service_user.Request{Phone: req.Phone, Passwd: req.Passwd, Code: req.Code})
	if err != nil || resp.Code != 0 || resp.Data == nil {
		logrus.Error(err)
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	token, err := tools.GenerateToken(env.JwtConf.Secret, resp.Data, time.Duration(time.Second*env.TokenExpire))
	if err != nil {
		log.Error(err)
		ctx.String(http.StatusOK, "login fail")
		return
	}
	ctx.SetCookie("token", token, env.TokenExpire, "/", "/", false, false)
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: dto.LoginResp{Token: token},
	})
}

// @Summary 注册
// @Description 注册
// @Tags 用户中心
// @Produce json
// @Param body body dto.UserRequest true "body参数"
// @Success 200 {object}  dto.Resp{}
// @Failure 500 {string} string "服务异常"
// @Router /register [post]
func (u *User) Register(ctx *gin.Context) {
	var req dto.RegUserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	if req.Phone < 10000000000 || req.Phone > 99999999999 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "phone is valid",
		})
		return
	}

	if len(req.Passwd) < 6 || len(req.Passwd) > 20 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "password length is between 6 and 20",
		})
		return
	}
	if req.Passwd != req.PasswdConfirm {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "confirm_passwd isn't same of password",
		})
		return
	}
	if !captcha.VerifyString(req.Id, req.Code) {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "code is invalid",
		})
		return
	}
	uRsp, err := u.UserCli.Find(ctx, &go_micro_service_user.Request{Phone: req.Phone})
	if err != nil || uRsp == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "register failure",
		})
		return
	}
	if uRsp.Code == 0 && uRsp.Data.UserId > 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "phone 已经被注册",
		})
		return
	}
	rsp, err := u.UserCli.Register(ctx, &go_micro_service_user.Request{Phone: req.Phone, Passwd: req.Passwd, Code: req.Code})
	if err != nil || rsp == nil || rsp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "register failure",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: nil,
	})

}

// @Summary 查找用户
// @Description 通过手机号查找用户
// @Tags 用户中心
// @Param phone path int true "手机号"
// @Success 200 {object}  dto.Resp{data=go_micro_service_user.UserInfo}
// @Failure 500 {string} string "服务异常"
// @Router /find/{phone} [get]
func (u *User) Find(ctx *gin.Context) {
	phone, err := strconv.ParseInt(ctx.Param("phone"), 10, 64)
	if err != nil {
		ctx.String(http.StatusOK, "phone err:%s", err.Error())
		return
	}
	if phone < 10000000000 || phone > 99999999999 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "phone is valid",
		})
		return
	}

	rsp, err := u.UserCli.Find(ctx, &go_micro_service_user.Request{Phone: phone})
	if err != nil || rsp.Code != 0 || rsp.Data == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: rsp.Data,
	})
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户中心
// @Produce json
// @Success 200 {object}  dto.Resp{}
// @Failure 500 {string} string "服务异常"
// @Router /user_info [post]
func (u *User) UserInfo(ctx *gin.Context) {
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.JSON(http.StatusUnauthorized, dto.Resp{
			Code: -1,
			Msg:  "get failure",
		})
		return
	}

	userInfo := authData.(map[string]interface{})
	wallet, err := u.WalletCli.GetOne(ctx, &go_micro_service_wallet.WalletReq{
		Uid: int64(userInfo["user_id"].(float64)),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "get failure",
		})
		return
	}
	userInfo["coins"] = wallet.AvailableBalance
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: userInfo,
	})
}
