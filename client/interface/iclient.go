package client

import (
	"github.com/hujinrun-github/go-press-test/constant"
	model "github.com/hujinrun-github/go-press-test/model/interface"
	"github.com/hujinrun-github/go-press-test/util"
)

type IClient interface {
	Init() constant.ErrorCode
	// 获取客户端的ID
	GetID() string
	// 获取客户端的名称
	GetName() string
	// 获取客户端的类型
	GetType() constant.ClientType
	// 获取客户端的地址
	GetAddress() util.Address
	// 获取客户端的连接状态
	GetStatus() constant.ErrorCode
	// 发送请求
	Request() constant.ErrorCode
}

type BaseClient struct {
	model.IModel
}

func (b *BaseClient) Init() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}

func (b *BaseClient) GetID() string {
	return constant.DEFAULT_CLIENT
}

func (b *BaseClient) GetName() string {
	return ""
}

func (b *BaseClient) GetType() string {
	return ""
}

func (b *BaseClient) GetAddress() util.Address {
	return util.Address{}
}

func (b *BaseClient) GetStatus() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}

func (b *BaseClient) Request() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}

func (b *BaseClient) LoadData() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}
