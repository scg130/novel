package alipay

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	apay "github.com/smartwalle/alipay/v3"
)

type AliPay struct {
	AppId         string
	RsaPrivateKey string
	NotifyUrl     string
	PublicKey     string
	Client        *apay.Client
}

var aliPayCli *AliPay

func NewAliPay(appId, rsaPrivateKey, publicKey, notifyUrl string) *AliPay {
	if aliPayCli == nil {
		client, _ := apay.New(appId, rsaPrivateKey, true)
		client.LoadAliPayPublicKey(publicKey)
		return &AliPay{
			AppId:         appId,
			RsaPrivateKey: rsaPrivateKey,
			NotifyUrl:     notifyUrl,
			PublicKey:     publicKey,
			Client:        client,
		}
	}
	return aliPayCli
}

func (self *AliPay) CreateOrder(orderNo, amount, subject string) (resp *CreateResp, err error) {
	p := apay.TradePreCreate{}
	p.TotalAmount = amount
	p.OutTradeNo = orderNo
	p.Subject = subject
	p.NotifyURL = self.NotifyUrl
	res, err := self.Client.TradePreCreate(p)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if res.IsSuccess() {
		return &CreateResp{
			OutTradeOrder: res.Content.OutTradeNo,
			Qrcode:        res.Content.QRCode,
		}, nil
	}
	log.Println(res.Content)

	return nil, errors.New(res.Content.Msg + res.Content.SubMsg)
}

func (self *AliPay) Query(outTradeNo string) {
	query := apay.TradeQuery{}
	query.OutTradeNo = outTradeNo
	self.Client.TradeQuery(query)

}

func (self *AliPay) Callback(req *http.Request, w http.ResponseWriter) *SuccessResp {
	req.ParseForm()
	rsp := &SuccessResp{Success: false}
	if flag, err := self.Client.VerifySign(req.Form); !flag || err != nil {
		return rsp
	}
	fmt.Println(req.Form)

	noti, err := self.Client.GetTradeNotification(req)
	if err != nil {
		log.Println(err.Error())
		return rsp
	}
	fmt.Println(noti.OutTradeNo, noti.TradeStatus, noti.TotalAmount, noti.BuyerPayAmount, noti.BuyerId, noti.Subject, noti.TradeNo)
	switch noti.TradeStatus {
	case "TRADE_SUCCESS":
		rsp.Success = true
		rsp.BuyerId, _ = strconv.Atoi(noti.BuyerId)
		rsp.OutTradeNo = noti.OutTradeNo
		rsp.Subject = noti.Subject
		rsp.PayAmount = noti.BuyerPayAmount
		rsp.TotalAmount = noti.TotalAmount
	default:
	}
	self.Client.AckNotification(w)
	return rsp
}
