package repo

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Novel struct {
	Id             int64     `xorm:" pk autoincr INT(11)" json:"id"`
	Name           string    `xorm:"not null unique(name) comment('分类名称') VARCHAR(255)" json:"name"`
	Author         string    `xorm:"not null default '' comment('排序') VARCHAR(255)" json:"author"`
	ChapterTotal   int32     `xorm:"not null default 0 comment('总章节数') int(11)" json:"chapter_total"`
	ChapterCurrent int32     `xorm:"not null default 0 comment('最新章节数') int(11)" json:"chapter_current"`
	Img            string    `xorm:"not null default '' comment('小说封面图') VARCHAR(300)" json:"img"`
	Intro          string    `xorm:"not null default '' comment('小说简介') VARCHAR(500)" json:"intro"`
	CateId         int       `xorm:"not null index(idx_cate_id_sort) default 0 comment('所属分类id') int(11)" json:"cate_id"`
	Words          int32     `xorm:"not null default 0 comment('总字数') int(11)" json:"words"`
	Likes          int32     `xorm:"not null default 0 comment('点赞数') int(11)" json:"likes"`
	UnLikes        int32     `xorm:"not null default 0 comment('不喜欢数') int(11)" json:"un_likes"`
	Sort           int       `xorm:"not null index(idx_cate_id_sort) default 0 comment('排序') int(11)" json:"sort"`
	UpdateTime     time.Time `xorm:"updated_at not null  DEFAULT CURRENT_TIMESTAMP  timestamp" json:"update_time"`
	CreateTime     time.Time `xorm:"created_at not null DEFAULT CURRENT_TIMESTAMP timestamp" json:"create_time"`

	NewChapter string `xorm:"not null  comment('分类名称') VARCHAR(255)" json:"new_chapter"`
	Category   string `json:"category" xorm:"-"`
}

func init() {
	novel := new(Novel)
	if isExist, _ := x.IsTableExist(novel); !isExist {
		if err := x.Sync2(novel); err != nil {
			log.Fatal(fmt.Sprintf("sync tables err:%v", err))
		}
	}
}

func (c *Novel) UpdateNovel(novel *Novel) error {
	_, err := x.Id(novel.Id).Update(novel)
	return err
}

func (c *Novel) GetByCateId(name string, cateId, page, size, words int) (data []Novel, total int64, err error) {
	wordCond := ""
	switch {
	case words == 0:
		wordCond = ""
	case (words <= 3000000):
		wordCond = "novel.words <= 3000000"
	case (words > 3000000 && words <= 5000000):
		wordCond = "novel.words>3000000 and novel.words<=5000000"
	case (words > 5000000 && words <= 10000000):
		wordCond = "novel.words>5000000 and novel.words<=10000000"
	case (words >= 10000001):
		wordCond = "novel.words>=10000001"
	}
	nameCond := ""
	if name != "" {
		nameCond = "novel.name like " + name + "%"
	}
	n := new(Novel)
	query := x.Table("novel").Join("inner", "chapter", "novel.id = chapter.novel_id and chapter.num = novel.chapter_current")
	if cateId != 0 {
		err = query.Where("novel.cate_id=?", cateId).
			And(nameCond).And(wordCond).
			And("novel.chapter_current>0").
			Select(
				"novel.id as id,novel.name as name,novel.author as author,novel.chapter_total as chapter_total,novel.chapter_current as chapter_current,novel.img as img,novel.intro as intro,novel.words as words,chapter.title as new_chapter,novel.updated_at as updated_at",
			).
			Limit(size, size*(page-1)).
			OrderBy("novel.sort").
			Find(&data)
		total, err = x.Where("novel.cate_id=?", cateId).And(nameCond).And(wordCond).And("novel.chapter_current>0").Count(n)
	} else {
		err = query.Where("novel.chapter_current>0").
			And(nameCond).And(wordCond).
			Select(
				"novel.id as id,novel.name as name,novel.author as author,novel.chapter_total as chapter_total,novel.chapter_current as chapter_current,novel.img as img,novel.intro as intro,novel.words as words,chapter.title as new_chapter,novel.updated_at as updated_at",
			).
			Limit(size, size*(page-1)).
			OrderBy("novel.sort").
			Find(&data)
		total, err = x.Where("novel.chapter_current>0").And(nameCond).And(wordCond).Count(n)
	}
	return
}

func (c *Novel) GetByName(name string, page, size int) (data []Novel, err error) {
	if name == "" {
		return
	}
	err = x.Where("name like ?", name+"%").Limit(size, size*(page-1)).Find(&data)

	return data, err
}

func (c *Novel) GetOne(novelId int) (Novel, error) {
	var novel Novel
	has, err := x.Id(novelId).Get(&novel)
	if !has {
		return novel, errors.New("no has")
	}
	return novel, err
}

func (c *Novel) FindList(page, pageSize int, name string, cateId int, author string) ([]Novel, int64, error) {
	list := make([]Novel, 0)
	query := x.Table("novel").Join("left", "category", "novel.cate_id = category.id")
	if name != "" {
		query = query.Where("novel.name like ?", name+"%")
	}
	if author != "" {
		query = query.Where("novel.author = ?", author)
	}
	if cateId != 0 {
		query = query.Where("novel.cate_id = ?", cateId)
	}

	total, err := query.Limit(pageSize, pageSize*(page-1)).Select("novel.id as id,novel.name as name,novel.author as author,novel.chapter_total as chapter_total,novel.chapter_current,novel.img as img,novel.cate_id as cate_id,novel.created_at as created_at,category.name as category,novel.sort as sort").OrderBy("novel.cate_id,novel.sort").FindAndCount(&list)
	return list, total, err
}
