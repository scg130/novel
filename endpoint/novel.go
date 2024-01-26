package endpoint

import (
	"fmt"
	"net/http"
	"novel/dto"
	go_micro_service_novel "novel/proto/novel"
	go_micro_service_user "novel/proto/user"
	go_micro_service_wallet "novel/proto/wallet"
	selfwrappers "novel/wrappers"
	"strconv"
	"sync"

	"github.com/scg130/tools"
	"github.com/scg130/tools/wrappers"

	"github.com/gin-gonic/gin"
)

type Novel struct {
	novelCli  go_micro_service_novel.NovelSrvService
	walletCli go_micro_service_wallet.WalletService
	userCli   go_micro_service_user.UserCenterService
}

var novelSrv *Novel

const (
	NOVEL_SRV_NAME  = "go.micro.service.novel"
	WALLET_SRV_NAME = "go.micro.service.wallet"
	USER_SRV_NAME   = "go.micro.service.user"
)

func NewNovelSrv() *Novel {
	if novelSrv == nil {
		novelSrv = &Novel{
			novelCli: go_micro_service_novel.NewNovelSrvService(NOVEL_SRV_NAME, tools.GetMicroClient(
				NOVEL_SRV_NAME,
				wrappers.NewTracerWrapper(),
				wrappers.NewRateLimitClientWrapper(100),
				selfwrappers.NewHystrixWrapper(),
			)),
			walletCli: go_micro_service_wallet.NewWalletService(WALLET_SRV_NAME, tools.GetMicroClient(
				WALLET_SRV_NAME,
				wrappers.NewTracerWrapper(),
				wrappers.NewRateLimitClientWrapper(100),
				selfwrappers.NewHystrixWrapper(),
			)),
			userCli: go_micro_service_user.NewUserCenterService(USER_SRV_NAME, tools.GetMicroClient(
				USER_SRV_NAME,
				wrappers.NewTracerWrapper(),
				wrappers.NewRateLimitClientWrapper(100),
				selfwrappers.NewHystrixWrapper(),
			)),
		}
	}
	return novelSrv
}

func (n *Novel) Index(ctx *gin.Context) {
	var r dto.IndexReq
	if err := ctx.Bind(&r); err != nil {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}
	novelsMap := make(map[string]interface{}, 0)
	wg := sync.WaitGroup{}
	for _, v := range r.Reqs {
		wg.Add(1)
		go func(req dto.NovelListReq) {
			defer wg.Done()
			resp, err := n.novelCli.GetNovelsByCateId(ctx, &go_micro_service_novel.Request{CateId: req.CateId, Page: req.Page, Size_: req.Size})
			fmt.Println(resp, err)
			if err != nil || resp.Code != 0 {
				return
			}
			novelsMap[req.NType] = resp.Novels
			return
		}(v)
	}
	wg.Wait()
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: novelsMap,
	})
}

// @Summary 分类列表
// @Description 分类列表
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Category}
// @Router /novel/cates [get]
func (n *Novel) Cates(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.Page == 0 || req.Size == 0 || req.Size > 100 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	resp, err := n.novelCli.GetCateGories(ctx, &go_micro_service_novel.Request{Page: req.Page, Size_: req.Size, IsShow: int32(1)})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Categories,
	})
}

// @Summary 搜索
// @Description 搜索
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Novel}
// @Router /novel/search [get]
func (n *Novel) SearchNovels(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.Page == 0 || req.Size == 0 || req.Size > 100 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	resp, err := n.novelCli.GetNovelsByName(ctx, &go_micro_service_novel.Request{Name: req.Name, Page: req.Page, Size_: req.Size})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Novels,
	})
}

// @Summary 小说列表
// @Description 小说列表
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Novel}
// @Router /novel/list [get]
func (n *Novel) Novels(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.Page == 0 || req.Size == 0 || req.Size > 10 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	resp, err := n.novelCli.GetNovelsByCateId(ctx, &go_micro_service_novel.Request{Name: req.Name, Words: req.Words, CateId: req.CateId, Page: req.Page, Size_: req.Size})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code:    0,
		Msg:     "ok",
		Data:    resp.Novels,
		Total:   resp.Total,
		CurPage: req.Page,
		CateId:  req.CateId,
	})
}

// @Summary 章节列表
// @Description 章节列表
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Chapter}
// @Router /novel/chapters [get]
func (n *Novel) Chapters(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.Page == 0 || req.Size == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	resp, err := n.novelCli.GetChaptersByNovelId(ctx, &go_micro_service_novel.Request{NovelId: req.NovelId, Page: req.Page, Size_: req.Size, Type: req.Type})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code:    0,
		Msg:     "ok",
		Data:    resp.Chapters,
		CurPage: req.Page,
	})
}

// @Summary 加入书架
// @Description 我的书架
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Note}
// @Router /novel/join-book [get]
func (n *Novel) JoinBook(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}
	userInfo := authData.(map[string]interface{})

	resp, err := n.novelCli.JoinNote(ctx, &go_micro_service_novel.Request{UserId: int32(userInfo["user_id"].(float64)), NovelId: req.NovelId})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: nil,
	})
}

// @Summary 删除阅读记录
// @Description 删除阅读记录
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{}
// @Router /novel/note/del [get]
func (n *Novel) DelNote(ctx *gin.Context) {
	novelId, _ := strconv.Atoi(ctx.Param("novel_id"))
	if novelId == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}
	userInfo := authData.(map[string]interface{})

	resp, err := n.novelCli.DelNote(ctx, &go_micro_service_novel.DelNoteReq{Uid: int32(userInfo["user_id"].(float64)), NovelId: int32(novelId)})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: nil,
	})
}

// @Summary 我的书架
// @Description 我的书架
// @Tags novel
// @Produce json
// @Param query query dto.NovelRequest true "query参数"
// @Success 200 {object}  dto.Resp{data=[]go_micro_service_novel.Note}
// @Router /novel/notes [get]
func (n *Novel) Notes(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.Page == 0 || req.Size == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}
	userInfo := authData.(map[string]interface{})

	resp, err := n.novelCli.GetNotes(ctx, &go_micro_service_novel.NoteRequest{UserId: int32(userInfo["user_id"].(float64)), Name: req.Name, Page: req.Page, Size_: req.Size, IsEnd: req.IsEnd})
	if err != nil || resp.Code != 0 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Notes,
	})
}

// @Summary 获取章节
// @Description 获取章节
// @Tags novel
// @Produce json
// @Param chapter_id query int true "query参数"
// @Success 200 {object}  dto.Resp{data=go_micro_service_novel.Chapter}
// @Router /novel/chapter [get]
func (n *Novel) Chapter(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.NovelId == 0 || req.Num == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}

	userInfo := authData.(map[string]interface{})

	resp, err := n.novelCli.GetChapterById(ctx, &go_micro_service_novel.Request{NovelId: req.NovelId, Num: req.Num, UserId: int32(userInfo["user_id"].(float64))})
	if err != nil || resp.Code != 0 || resp.Chapter == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	if resp.Chapter.IsVip == go_micro_service_novel.VipType_IS_VIP {
		rsp, err := n.walletCli.GetChapter(ctx, &go_micro_service_wallet.BuyChapterRequest{
			Uid:       int64(userInfo["user_id"].(float64)),
			ChapterId: int64(resp.Chapter.ChapterId),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		if rsp.State != 1 {
			ctx.JSON(http.StatusOK, dto.Resp{
				Code: 1,
				Msg:  "该章节属性vip章节,请先购买或开通vip。",
				Data: gin.H{
					"chapter_id": resp.Chapter.ChapterId,
				},
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Chapter,
	})
}

// @Summary 获取小说详情
// @Description 获取小说详情
// @Tags novel
// @Produce json
// @Param novel_id query int true "query参数"
// @Success 200 {object}  dto.Resp{data=go_micro_service_novel.Novel}
// @Router /novel/novel [get]
func (n *Novel) Novel(ctx *gin.Context) {
	var req dto.NovelRequest
	if err := ctx.BindQuery(&req); err != nil || req.NovelId == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}

	//span := opentracing.StartSpan("novel")
	//defer span.Finish()
	//c := opentracing.ContextWithSpan(ctx, span)

	resp, err := n.novelCli.GetNovelById(ctx, &go_micro_service_novel.Request{NovelId: req.NovelId})
	if err != nil || resp.Code != 0 || resp.Novel == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Novel,
	})
}

// @Summary 获取购买历史
// @Description 获取购买历史
// @Tags novel
// @Produce json
// @Param page query int true "query参数"
// @Param size query int true "query参数"
// @Success 200 {object}  dto.Resp{data=go_micro_service_wallet.Log}
// @Router /novel/buy_logs [get]
func (n *Novel) BuyLogs(ctx *gin.Context) {
	var req dto.BuyRequest
	if err := ctx.BindQuery(&req); err != nil || req.Size == 0 {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}

	userInfo := authData.(map[string]interface{})

	resp, err := n.walletCli.FindBuyLog(ctx, &go_micro_service_wallet.LogRequest{
		Uid:   int64(userInfo["user_id"].(float64)),
		Page:  req.Page,
		Size_: req.Size,
	})
	if err != nil || resp.State != 1 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "failure",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "ok",
		Data: resp.Log,
	})
}

// @Summary 购买章节
// @Description 购买章节
// @Tags novel
// @Produce json
// @Param chapter_id query int true "query参数"
// @Success 200 {object}  dto.Resp{}
// @Router /novel/buy_chapter?chapter_id=1 [get]
func (self *Novel) BuyChapter(ctx *gin.Context) {
	var req dto.BuyRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.String(http.StatusBadRequest, "bad params")
		return
	}
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return
	}

	userInfo := authData.(map[string]interface{})

	chapter, err := self.novelCli.GetChapterById(ctx, &go_micro_service_novel.Request{
		NovelId: req.NovelId,
		Num:     req.Num,
	})
	if err != nil || chapter.Code != 0 || chapter.Chapter == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "购买失败",
		})
		return
	}
	novel, err := self.novelCli.GetNovelById(ctx, &go_micro_service_novel.Request{
		NovelId: chapter.Chapter.NovelId,
	})
	if err != nil || novel.Code != 0 || novel.Novel == nil {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "购买失败",
		})
		return
	}
	wRsp, err := self.walletCli.GetOne(ctx, &go_micro_service_wallet.WalletReq{
		Uid: int64(userInfo["user_id"].(float64)),
	})
	if err != nil || wRsp.AvailableBalance < 100 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "余额不足,请先充值！",
		})
		return
	}
	resp, err := self.walletCli.BuyChapter(ctx, &go_micro_service_wallet.BuyChapterRequest{
		Uid:       int64(userInfo["user_id"].(float64)),
		ChapterId: req.ChapterId,
		NovelId:   int64(novel.Novel.NovelId),
		NovelName: novel.Novel.Name,
		Amount:    100,
	})
	if err != nil || resp.State != 1 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "购买失败",
		})
		return
	}
	changeRep, err := self.walletCli.Change(ctx, &go_micro_service_wallet.WalletReq{
		Uid:    int64(userInfo["user_id"].(float64)),
		Amount: 100,
		Type:   go_micro_service_wallet.Type_STATE_BUY_CHAPTER,
	})
	if err != nil || changeRep.State != 1 {
		ctx.JSON(http.StatusOK, dto.Resp{
			Code: -1,
			Msg:  "购买失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Resp{
		Code: 0,
		Msg:  "购买成功",
	})
	return
}
