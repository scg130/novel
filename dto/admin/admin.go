package admin

type Resp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Pagnation `json:"pagnation"`
}
