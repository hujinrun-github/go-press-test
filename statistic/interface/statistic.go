package statistic

import (
	"sync"

	"github.com/hujinrun-github/go-press-test/constant"
	model "github.com/hujinrun-github/go-press-test/model/interface"
)

type IStatistic interface {
	Init() constant.ErrorCode
	ReceivingResults(ch <-chan model.IRequestResult, wg *sync.WaitGroup)
	CalculateData(model.IRequestResult)
}

type BaseStatistic struct {
}

func (b *BaseStatistic) Init() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}

func (b *BaseStatistic) ReceivingResults(ch <-chan model.IRequestResult, wg *sync.WaitGroup) {}

func (b *BaseStatistic) CalculateData(r model.IRequestResult) {}
