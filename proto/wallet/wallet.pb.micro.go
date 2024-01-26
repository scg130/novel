// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/wallet/wallet.proto

package go_micro_srv_charge

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Wallet service

func NewWalletEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Wallet service

type WalletService interface {
	Create(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error)
	Change(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error)
	GetOne(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error)
	FindBuyLog(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error)
	BuyChapter(ctx context.Context, in *BuyChapterRequest, opts ...client.CallOption) (*WalletResponse, error)
	GetChapter(ctx context.Context, in *BuyChapterRequest, opts ...client.CallOption) (*LogResponse, error)
}

type walletService struct {
	c    client.Client
	name string
}

func NewWalletService(name string, c client.Client) WalletService {
	return &walletService{
		c:    c,
		name: name,
	}
}

func (c *walletService) Create(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Create", in)
	out := new(WalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Change(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Change", in)
	out := new(WalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) GetOne(ctx context.Context, in *WalletReq, opts ...client.CallOption) (*WalletResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.GetOne", in)
	out := new(WalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) FindBuyLog(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.FindBuyLog", in)
	out := new(LogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) BuyChapter(ctx context.Context, in *BuyChapterRequest, opts ...client.CallOption) (*WalletResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.BuyChapter", in)
	out := new(WalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) GetChapter(ctx context.Context, in *BuyChapterRequest, opts ...client.CallOption) (*LogResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.GetChapter", in)
	out := new(LogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wallet service

type WalletHandler interface {
	Create(context.Context, *WalletReq, *WalletResponse) error
	Change(context.Context, *WalletReq, *WalletResponse) error
	GetOne(context.Context, *WalletReq, *WalletResponse) error
	FindBuyLog(context.Context, *LogRequest, *LogResponse) error
	BuyChapter(context.Context, *BuyChapterRequest, *WalletResponse) error
	GetChapter(context.Context, *BuyChapterRequest, *LogResponse) error
}

func RegisterWalletHandler(s server.Server, hdlr WalletHandler, opts ...server.HandlerOption) error {
	type wallet interface {
		Create(ctx context.Context, in *WalletReq, out *WalletResponse) error
		Change(ctx context.Context, in *WalletReq, out *WalletResponse) error
		GetOne(ctx context.Context, in *WalletReq, out *WalletResponse) error
		FindBuyLog(ctx context.Context, in *LogRequest, out *LogResponse) error
		BuyChapter(ctx context.Context, in *BuyChapterRequest, out *WalletResponse) error
		GetChapter(ctx context.Context, in *BuyChapterRequest, out *LogResponse) error
	}
	type Wallet struct {
		wallet
	}
	h := &walletHandler{hdlr}
	return s.Handle(s.NewHandler(&Wallet{h}, opts...))
}

type walletHandler struct {
	WalletHandler
}

func (h *walletHandler) Create(ctx context.Context, in *WalletReq, out *WalletResponse) error {
	return h.WalletHandler.Create(ctx, in, out)
}

func (h *walletHandler) Change(ctx context.Context, in *WalletReq, out *WalletResponse) error {
	return h.WalletHandler.Change(ctx, in, out)
}

func (h *walletHandler) GetOne(ctx context.Context, in *WalletReq, out *WalletResponse) error {
	return h.WalletHandler.GetOne(ctx, in, out)
}

func (h *walletHandler) FindBuyLog(ctx context.Context, in *LogRequest, out *LogResponse) error {
	return h.WalletHandler.FindBuyLog(ctx, in, out)
}

func (h *walletHandler) BuyChapter(ctx context.Context, in *BuyChapterRequest, out *WalletResponse) error {
	return h.WalletHandler.BuyChapter(ctx, in, out)
}

func (h *walletHandler) GetChapter(ctx context.Context, in *BuyChapterRequest, out *LogResponse) error {
	return h.WalletHandler.GetChapter(ctx, in, out)
}