package examples

import (
	"sync"
	"time"

	statistic "github.com/hujinrun-github/go-press-test/statistic/interface"
)

type ResultInfo struct {
	requestID  uint64
	isSuccess  bool
	errMessage error
	cost       int
	extra      map[string][]byte
}

type SuccessStatistic struct {
	statistic.BaseStatistic
	resultData []ResultInfo
	costMap    sync.Map
}

func (s *SuccessStatistic) PerStart(id uint64) {
	s.costMap.Store(id, time.Now().Unix())
}

func (s *SuccessStatistic) PerEnd(meta statistic.IStatisticMeta) {
	istartTime, _ := s.costMap.Load(meta.GetRequestID())
	startTime, ok := istartTime.(int64)
	if !ok {
		s.resultData = append(s.resultData, ResultInfo{
			requestID:  meta.GetRequestID(),
			cost:       -1,
			errMessage: meta.GetError(),
			extra:      meta.GetExtra(),
		})
	} else {
		s.resultData = append(s.resultData, ResultInfo{
			requestID:  meta.GetRequestID(),
			cost:       int(time.Now().Unix()) - int(startTime),
			errMessage: meta.GetError(),
			extra:      meta.GetExtra(),
		})
	}
}

func (s *SuccessStatistic) Parse() {

}
