package model

import (
	"sync"

	"github.com/hujinrun-github/go-press-test/constant"
)

type IModel interface {
	LoadData(any) constant.ErrorCode
	Request(chan IRequestResult, *sync.WaitGroup) constant.ErrorCode
}

type IRequestResult interface {
	GetID() string
	GetTime() uint64
	GetStatus() int
	GetErrorCode() constant.ErrorCode
	GetErrorMsg() string
	GetData() map[string][]byte
}

type BaseModel struct {
}

func (b *BaseModel) Request(chan IRequestResult, *sync.WaitGroup) (any, constant.ErrorCode) {
	return nil, constant.ERR_CODE_SUCCESS
}

func (b *BaseModel) LoadData(any) constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}
