// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/bottos-project/bottos/api/chain.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Chain service

type ChainService interface {
	SendTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*SendTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...client.CallOption) (*GetTransactionResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...client.CallOption) (*GetBlockResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error)
	GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...client.CallOption) (*GetKeyValueResponse, error)
	GetAbi(ctx context.Context, in *GetAbiRequest, opts ...client.CallOption) (*GetAbiResponse, error)
	GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, opts ...client.CallOption) (*GetTransferCreditResponse, error)
}

type chainService struct {
	c    client.Client
	name string
}

func NewChainService(name string, c client.Client) ChainService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &chainService{
		c:    c,
		name: name,
	}
}

func (c *chainService) SendTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*SendTransactionResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.SendTransaction", in)
	out := new(SendTransactionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...client.CallOption) (*GetTransactionResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetTransaction", in)
	out := new(GetTransactionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...client.CallOption) (*GetBlockResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetBlock", in)
	out := new(GetBlockResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetInfo", in)
	out := new(GetInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetAccount", in)
	out := new(GetAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...client.CallOption) (*GetKeyValueResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetKeyValue", in)
	out := new(GetKeyValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetAbi(ctx context.Context, in *GetAbiRequest, opts ...client.CallOption) (*GetAbiResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetAbi", in)
	out := new(GetAbiResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainService) GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, opts ...client.CallOption) (*GetTransferCreditResponse, error) {
	req := c.c.NewRequest(c.name, "Chain.GetTransferCredit", in)
	out := new(GetTransferCreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chain service

type ChainHandler interface {
	SendTransaction(context.Context, *Transaction, *SendTransactionResponse) error
	GetTransaction(context.Context, *GetTransactionRequest, *GetTransactionResponse) error
	GetBlock(context.Context, *GetBlockRequest, *GetBlockResponse) error
	GetInfo(context.Context, *GetInfoRequest, *GetInfoResponse) error
	GetAccount(context.Context, *GetAccountRequest, *GetAccountResponse) error
	GetKeyValue(context.Context, *GetKeyValueRequest, *GetKeyValueResponse) error
	GetAbi(context.Context, *GetAbiRequest, *GetAbiResponse) error
	GetTransferCredit(context.Context, *GetTransferCreditRequest, *GetTransferCreditResponse) error
}

func RegisterChainHandler(s server.Server, hdlr ChainHandler, opts ...server.HandlerOption) {
	type chain interface {
		SendTransaction(ctx context.Context, in *Transaction, out *SendTransactionResponse) error
		GetTransaction(ctx context.Context, in *GetTransactionRequest, out *GetTransactionResponse) error
		GetBlock(ctx context.Context, in *GetBlockRequest, out *GetBlockResponse) error
		GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error
		GetAccount(ctx context.Context, in *GetAccountRequest, out *GetAccountResponse) error
		GetKeyValue(ctx context.Context, in *GetKeyValueRequest, out *GetKeyValueResponse) error
		GetAbi(ctx context.Context, in *GetAbiRequest, out *GetAbiResponse) error
		GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, out *GetTransferCreditResponse) error
	}
	type Chain struct {
		chain
	}
	h := &chainHandler{hdlr}
	s.Handle(s.NewHandler(&Chain{h}, opts...))
}

type chainHandler struct {
	ChainHandler
}

func (h *chainHandler) SendTransaction(ctx context.Context, in *Transaction, out *SendTransactionResponse) error {
	return h.ChainHandler.SendTransaction(ctx, in, out)
}

func (h *chainHandler) GetTransaction(ctx context.Context, in *GetTransactionRequest, out *GetTransactionResponse) error {
	return h.ChainHandler.GetTransaction(ctx, in, out)
}

func (h *chainHandler) GetBlock(ctx context.Context, in *GetBlockRequest, out *GetBlockResponse) error {
	return h.ChainHandler.GetBlock(ctx, in, out)
}

func (h *chainHandler) GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error {
	return h.ChainHandler.GetInfo(ctx, in, out)
}

func (h *chainHandler) GetAccount(ctx context.Context, in *GetAccountRequest, out *GetAccountResponse) error {
	return h.ChainHandler.GetAccount(ctx, in, out)
}

func (h *chainHandler) GetKeyValue(ctx context.Context, in *GetKeyValueRequest, out *GetKeyValueResponse) error {
	return h.ChainHandler.GetKeyValue(ctx, in, out)
}

func (h *chainHandler) GetAbi(ctx context.Context, in *GetAbiRequest, out *GetAbiResponse) error {
	return h.ChainHandler.GetAbi(ctx, in, out)
}

func (h *chainHandler) GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, out *GetTransferCreditResponse) error {
	return h.ChainHandler.GetTransferCredit(ctx, in, out)
}
