package model

import "github.com/hujinrun-github/go-press-test/constant"

type IModel interface {
	LoadData(any) (any, constant.ErrorCode)
}

type BaseModel struct {
}

func (b *BaseModel) LoadData(any) (any, constant.ErrorCode) {
	return "", constant.ERR_CODE_SUCCESS
}
