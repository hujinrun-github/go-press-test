package model

import "github.com/hujinrun-github/go-press-test/constant"

type IModel interface {
	LoadData(any) ([]any, constant.ErrorCode)
	Request(any) (any, constant.ErrorCode)
	ParseRsp(any) (any, constant.ErrorCode)
}

type BaseModel struct {
}

func (b *BaseModel) Request(any) (any, constant.ErrorCode) {
	return nil, constant.ERR_CODE_SUCCESS
}

func (b *BaseModel) LoadData(any) ([]any, constant.ErrorCode) {
	return nil, constant.ERR_CODE_SUCCESS
}

func (b *BaseModel) ParseRsp(any) (any, constant.ErrorCode) {
	return nil, constant.ERR_CODE_SUCCESS
}
