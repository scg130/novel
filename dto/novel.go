package dto

type NovelRequest struct {
	Page      int32  `json:"page" form:"page"`
	Size      int32  `json:"size" form:"size"`
	CateId    int32  `json:"cate_id" form:"cate_id"`
	NovelId   int32  `json:"novel_id" form:"novel_id"`
	ChapterId int32  `json:"chapter_id" form:"chapter_id"`
	Type      string `json:"type" form:"type"`
	Num       int32  `json:"num" form:"num"`
	Name      string `json:"name" form:"name"`
	Words     int32  `json:"words" form:"words"`
	IsEnd     int32  `json:"is_end" form:"is_end"`
}

type BuyRequest struct {
	Page      int32 `json:"page" form:"page"`
	Size      int32 `json:"size" form:"size"`
	ChapterId int64 `json:"chapter_id" form:"chapter_id"`
	NovelId   int32 `json:"novel_id" form:"novel_id"`
	Num       int32 `json:"num" form:"num"`
}

type IndexReq struct {
	Reqs []NovelListReq `json:"reqs" form:"reqs"`
}

type NovelListReq struct {
	Page   int32  `json:"page" form:"page"`
	Size   int32  `json:"size" form:"size"`
	CateId int32  `json:"cate_id" form:"cate_id"`
	NType  string `json:"n_type" form:"n_type"`
}
