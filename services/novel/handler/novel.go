package handler

import (
	"context"
	"encoding/json"
	"fmt"
	novel "novel/proto/novel"
	"novel/repo"
	"time"

	"github.com/scg130/tools/bigcache"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

type NovelSrv struct {
	Cli     client.Client
	Cate    repo.Category
	Novel   repo.Novel
	Chapter repo.Chapter
	Notes   repo.Notes
}

func (this *NovelSrv) DelNote(ctx context.Context, req *novel.DelNoteReq, rsp *novel.Response) error {
	err := this.Notes.DelNote(req.NovelId, req.Uid)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (this *NovelSrv) SetVipChapter(ctx context.Context, req *novel.SetVipChapterReq, rsp *novel.Response) error {
	err := this.Chapter.SetVipChapter(int(req.NovelId), int(req.MinChapter), int(req.MaxChapter), int(req.IsVip))
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (this *NovelSrv) UpdateNovel(ctx context.Context, req *novel.Novel, rsp *novel.Response) error {
	err := this.Novel.UpdateNovel(&repo.Novel{
		Id:     int64(req.NovelId),
		Name:   req.Name,
		Author: req.Author,
		Img:    req.Img,
		CateId: int(req.CateId),
		Sort:   int(req.Sort),
	})
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (n *NovelSrv) GetNovelList(ctx context.Context, req *novel.NovelListReq, rsp *novel.NovelListResp) error {
	novels := make([]*novel.NovelData, 0)
	datas, total, err := n.Novel.FindList(int(req.Page), int(req.PageSize), req.Name, int(req.CateId), req.Author)
	if err != nil {
		rsp.Code = -1
		return err
	}
	rsp.Code = 0
	for _, data := range datas {
		novels = append(novels, &novel.NovelData{
			Id:             data.Id,
			Name:           data.Name,
			Author:         data.Author,
			ChapterTotal:   int64(data.ChapterTotal),
			ChapterCurrent: int64(data.ChapterCurrent),
			Img:            data.Img,
			CateName:       data.Category,
			CateId:         int64(data.CateId),
			CreateAt:       data.CreateTime.Format("2006-01-02 15:04:05"),
			Sort:           int32(data.Sort),
		})
	}
	rsp.Novels = novels
	rsp.Pagnation = &novel.Pagnation{
		Total:    int64(total),
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	}
	return nil
}

func (n *NovelSrv) DelCategory(ctx context.Context, req *novel.DelCategoryReq, rsp *novel.CommonResponse) error {
	err := n.Cate.Del(req.CategoryId)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (n *NovelSrv) UpdateCategory(ctx context.Context, req *novel.Category, rsp *novel.CommonResponse) error {
	err := n.Cate.Update(&repo.Category{
		Id:         int64(req.CateId),
		Name:       req.Name,
		Channel:    req.Channel,
		Sort:       req.Sort,
		IsShow:     req.IsShow,
		UpdateTime: time.Now(),
	})
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (n *NovelSrv) AddCateGory(ctx context.Context, req *novel.AddCateRequest, rsp *novel.CommonResponse) error {
	_, err := n.Cate.Create(repo.Category{
		Name:       req.Name,
		Channel:    req.Channel,
		Sort:       req.Sort,
		IsShow:     req.IsShow,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	})
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}

func (n *NovelSrv) GetCateGories(ctx context.Context, req *novel.Request, rsp *novel.CateResponse) error {
	cates := make([]*novel.Category, 0)
	datas, total, err := n.Cate.Get(int(req.Page), int(req.Size_), int(req.IsShow), req.Name)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	for _, data := range datas {
		cates = append(cates, &novel.Category{
			CateId:  int32(data.Id),
			Name:    data.Name,
			Sort:    int32(data.Sort),
			Channel: data.Channel,
			IsShow:  data.IsShow,
		})
	}
	rsp.Categories = cates
	rsp.Pagnation = &novel.Pagnation{
		Total:    int64(total),
		Page:     int64(req.Page),
		PageSize: int64(req.Size_),
	}
	return nil
}

func (n *NovelSrv) JoinNote(ctx context.Context, req *novel.Request, rsp *novel.NoteResponse) error {
	note, err := n.Notes.GetLastNote(req.UserId, req.NovelId)
	if err == nil && note.Id > 0 {
		return nil
	}
	err = n.Notes.CreateNote(req.UserId, req.NovelId, 1)
	if err != nil {
		return err
	}
	return nil
}

func (n *NovelSrv) GetNotes(ctx context.Context, req *novel.NoteRequest, rsp *novel.NoteResponse) error {
	notes := make([]*novel.Note, 0)
	novels, err := n.Notes.GetNotes(req.Name, int(req.UserId), int(req.Page), int(req.Size_), int(req.IsEnd))
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		return err
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	for _, data := range novels {
		newChapter, err := n.Chapter.GetOne(int(data.NovelId), int(data.NewNum))
		if err != nil {
			continue
		}
		prevChapter, err := n.Chapter.GetOne(int(data.NovelId), int(data.PreNum))
		if err != nil {
			continue
		}
		notes = append(notes, &novel.Note{
			NovelId:     data.NovelId,
			PrevNum:     data.PreNum,
			NewChapter:  newChapter.Title,
			PrevChapter: prevChapter.Title,
			NovelName:   data.Name,
			NewNum:      data.NewNum,
		})
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	rsp.Notes = notes
	return nil
}

func (n *NovelSrv) GetNovelsByCateId(ctx context.Context, req *novel.Request, rsp *novel.NovelsResponse) error {
	key := fmt.Sprintf("ns:%s:%d:%d:%d:%d", req.Name, int(req.CateId), int(req.Page), req.Size_, int(req.Words))
	totalKey := fmt.Sprintf("total:%s:%d:%d:%d:%d", req.Name, int(req.CateId), int(req.Page), req.Size_, int(req.Words))
	novData, err := bigcache.BigCache.Get(key)
	novs := make([]*novel.Novel, 0)
	if err != nil {
		novels, total, err := n.Novel.GetByCateId(req.Name, int(req.CateId), int(req.Page), int(req.Size_), int(req.Words))
		if err != nil {
			rsp.Code = -1
			rsp.Msg = "failure"
			return err
		}
		rsp.Code = 0
		rsp.Msg = "ok"
		rsp.Total = int32(total)
		totalBytes, _ := json.Marshal(total)
		bigcache.BigCache.Set(totalKey, totalBytes)
		for _, data := range novels {
			novs = append(novs, &novel.Novel{
				NovelId:        int32(data.Id),
				Name:           data.Name,
				Author:         data.Author,
				ChapterTotal:   data.ChapterTotal,
				ChapterCurrent: data.ChapterCurrent,
				Img:            data.Img,
				Intro:          data.Intro,
				Words:          data.Words,
				NewChapter:     data.NewChapter,
				Likes:          int32(data.Likes),
				UnLikes:        int32(data.UnLikes),
				UpdatedAt:      data.UpdateTime.Format("2006-01-02 15:04:05"),
			})
		}
		bytes, _ := json.Marshal(novs)
		bigcache.BigCache.Set(key, bytes)
		rsp.Novels = novs
		return nil
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	json.Unmarshal(novData, &novs)
	totalBytes, _ := bigcache.BigCache.Get(totalKey)
	var total int32
	json.Unmarshal(totalBytes, &total)
	rsp.Total = total
	rsp.Novels = novs
	return nil
}

func (n *NovelSrv) GetNovelsByName(ctx context.Context, req *novel.Request, rsp *novel.NovelsResponse) error {
	novelsKey := fmt.Sprintf("novels:%s:%d:%d", req.Name, int(req.Page), req.Size_)
	novelsData, err := bigcache.BigCache.Get(novelsKey)
	novs := make([]*novel.Novel, 0)
	if err != nil {
		novels, err := n.Novel.GetByName(req.Name, int(req.Page), int(req.Size_))
		if err != nil {
			rsp.Code = -1
			rsp.Msg = "failure"
			return err
		}
		rsp.Code = 0
		rsp.Msg = "ok"
		for _, data := range novels {
			novs = append(novs, &novel.Novel{
				NovelId:        int32(data.Id),
				Name:           data.Name,
				Author:         data.Author,
				ChapterTotal:   data.ChapterTotal,
				ChapterCurrent: data.ChapterCurrent,
				Img:            data.Img,
				Intro:          data.Intro,
				Words:          data.Words,
				Likes:          int32(data.Likes),
				UnLikes:        int32(data.UnLikes),
			})
		}
		bytes, _ := json.Marshal(novs)
		bigcache.BigCache.Set(novelsKey, bytes)
		rsp.Novels = novs
		return nil
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	json.Unmarshal(novelsData, &novs)
	rsp.Novels = novs
	return nil
}

func (n *NovelSrv) GetNovelById(ctx context.Context, req *novel.Request, rsp *novel.NovelResponse) error {
	novelKey := fmt.Sprintf("novel:%d", int(req.NovelId))
	novelData, err := bigcache.BigCache.Get(novelKey)
	var nov novel.Novel
	if err != nil {
		data, err := n.Novel.GetOne(int(req.NovelId))
		if err != nil {
			rsp.Code = -1
			rsp.Msg = "failure"
			return err
		}
		rsp.Code = 0
		rsp.Msg = "ok"
		nov = novel.Novel{
			NovelId:        int32(data.Id),
			Name:           data.Name,
			Author:         data.Author,
			ChapterTotal:   data.ChapterTotal,
			ChapterCurrent: data.ChapterCurrent,
			Img:            data.Img,
			Intro:          data.Intro,
			Words:          data.Words,
			Likes:          data.Likes,
			UnLikes:        data.UnLikes,
			UpdatedAt:      data.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		bytes, _ := json.Marshal(nov)
		bigcache.BigCache.Set(novelKey, bytes)
		rsp.Novel = &nov
		return nil
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	json.Unmarshal(novelData, &nov)
	rsp.Novel = &nov
	return nil
}

func (n *NovelSrv) GetChaptersByNovelId(ctx context.Context, req *novel.Request, rsp *novel.ChaptersResponse) error {
	chaptersKey := fmt.Sprintf("chapters:%d:%d:%d:%d", int(req.NovelId), int(req.Page), req.Size_, req.Type)
	chaptersData, err := bigcache.BigCache.Get(chaptersKey)
	chapters := make([]*novel.Chapter, 0)
	if err != nil {
		repoChapters, err := n.Chapter.GetByNovelId(int(req.NovelId), int(req.Page), int(req.Size_), req.Type)
		if err != nil {
			rsp.Code = -1
			rsp.Msg = err.Error()
			return err
		}

		rsp.Code = 0
		rsp.Msg = "ok"
		for _, data := range repoChapters {
			chapters = append(chapters, &novel.Chapter{
				ChapterId: int32(data.Id),
				Title:     data.Title,
				// Content:   data.Content,
				Words:   int32(data.Words),
				NovelId: int32(data.NovelId),
				IsVip:   novel.VipType(data.IsVip),
				Num:     int32(data.Num),
			})
		}
		rsp.Chapters = chapters
		bytes, _ := json.Marshal(chapters)
		bigcache.BigCache.Set(chaptersKey, bytes)
		return nil
	}
	rsp.Code = 0
	rsp.Msg = "ok"
	json.Unmarshal(chaptersData, &chapters)
	rsp.Chapters = chapters
	return nil
}

func (n *NovelSrv) GetChapterById(ctx context.Context, req *novel.Request, rsp *novel.ChapterResponse) error {
	var chapter novel.Chapter

	data, err := n.Chapter.GetOne(int(req.NovelId), int(req.Num))
	if err != nil {
		rsp.Code = -1
		rsp.Msg = "failure"
		return err
	}
	pub := micro.NewEvent("novel.read", n.Cli)
	err = pub.Publish(ctx, &novel.ReadRequest{
		NovelId:    req.NovelId,
		ChapterNum: req.Num,
		UserId:     req.UserId,
	})

	chapter = novel.Chapter{
		ChapterId: int32(data.Id),
		Title:     data.Title,
		Content:   data.Content,
		Words:     int32(data.Words),
		NovelId:   int32(data.NovelId),
		IsVip:     novel.VipType(data.IsVip),
		Num:       int32(data.Num),
	}

	rsp.Code = 0
	rsp.Msg = "ok"
	rsp.Chapter = &chapter
	return nil
}
