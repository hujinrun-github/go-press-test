package statistic

type IStatistic interface {
	AllStart()
	AllEnd()
	PerStart()
	PerEnd()
	Parse()
}

type BaseStatistic struct {
}

func (bs *BaseStatistic) AllStart() {}
func (bs *BaseStatistic) AllEnd()   {}
func (bs *BaseStatistic) PerStart() {}
func (bs *BaseStatistic) PerEnd()   {}
func (bs *BaseStatistic) Parse()    {}
