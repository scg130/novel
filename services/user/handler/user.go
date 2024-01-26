package handler

import (
	"context"
	"time"

	"github.com/scg130/tools"

	"github.com/google/uuid"

	user "user/proto/user"
	"user/repo"
	"errors"
)

type UserSrv struct {
	U repo.User
}

func (e *UserSrv) Login(ctx context.Context, req *user.Request, rsp *user.Response) error {
	userInfo, err := e.U.FindByPhone(req.Phone)
	if err != nil {
		return err
	}
	if userInfo.Id == 0 {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return nil
	}
	if tools.CompareHashAndPasswd(req.Passwd, userInfo.Password) {
		rsp.Code = 0
		rsp.Msg = "ok"
		rsp.Data = &user.UserInfo{
			UserId: userInfo.Id,
			Phone:  userInfo.Phone,
		}
		return nil
	}
	rsp.Code = -1
	rsp.Msg = "login failure"
	return nil
}

func (e *UserSrv) Register(ctx context.Context, req *user.Request, rsp *user.Response) error {
	userInfo, _ := e.U.FindByPhone(req.Phone)
	if userInfo.Id > 0 {
		rsp.Code = -1
		rsp.Msg = "user exist"
		return errors.New("user exist")
	}

	passwd, _ := tools.GeneratePasswd(req.Passwd)
	flag, err := e.U.Create(repo.User{
		Phone:      req.Phone,
		Password:   passwd,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Version:    uuid.New().ID(),
	})
	if !flag || err != nil {
		rsp.Code = -1
		rsp.Msg = "create user fail"
		return errors.New("fail")
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	return nil
}

func (e *UserSrv) Find(ctx context.Context, req *user.Request, rsp *user.Response) error {
	userInfo, err := e.U.FindByPhone(req.Phone)
	if userInfo.Id == 0 {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return nil
	}
	rsp.Code = 0
	rsp.Data = &user.UserInfo{
		UserId: userInfo.Id,
		Phone:  userInfo.Phone,
	}
	return nil
}
