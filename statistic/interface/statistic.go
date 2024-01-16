package statistic

type IStatistic interface {
	AllStart()
	AllEnd()
	PerStart()
	PerEnd()
	Parse() string
}

type IStatisticMeta interface {
	GetCode() STATISTIC_RESULT
	GetExtra() map[string][]byte
	GetError() error
	GetOriginRequest() any
	GetRequestID() uint64
}

type BaseStatistic struct {
}

func (bs *BaseStatistic) AllStart()             {}
func (bs *BaseStatistic) AllEnd()               {}
func (bs *BaseStatistic) PerStart(uint64)       {}
func (bs *BaseStatistic) PerEnd(IStatisticMeta) {}
func (bs *BaseStatistic) Parse() string         { return "" }
