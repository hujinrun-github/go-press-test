package server

import (
	"context"
	"sync"
	"time"

	"github.com/hujinrun-github/go-press-test/constant"
	model "github.com/hujinrun-github/go-press-test/model/interface"
	statistic "github.com/hujinrun-github/go-press-test/statistic/interface"
)

type Server struct {
	// c        client.IClient
	s        statistic.IStatistic
	m        model.IModel
	parallel uint64
}

type Option func(*Server)

// func WithClient(c client.IClient) Option {
// 	return func(s *Server) {
// 		s.c = c
// 	}
// }

func WithStatistic(st statistic.IStatistic) Option {
	return func(s *Server) {
		s.s = st
	}
}

func WithParallel(p uint64) Option {
	return func(s *Server) {
		s.parallel = p
	}
}

func WithModel(m model.IModel) Option {
	return func(s *Server) {
		s.m = m
	}
}

// NewServer 创建服务
func NewServer(opts ...Option) *Server {
	s := &Server{
		parallel: constant.DEFAULT_PARALLEL,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Run(ctx context.Context, datasource any) constant.ErrorCode {
	ch := make(chan model.IRequestResult, 1000)
	var (
		swg sync.WaitGroup //发送数据完成
		rwg sync.WaitGroup //接收数据完成
	)

	go s.s.ReceivingResults(ch, &rwg)

	code := s.m.LoadData(datasource)
	if code != constant.ERR_CODE_SUCCESS {
		return code
	}

	for i := uint64(0); i < uint64(s.parallel); i++ {
		swg.Add(1)
		go s.m.Request(ch, &swg)
	}

	swg.Wait()
	time.Sleep(1 * time.Millisecond)
	close(ch)
	rwg.Wait()
	return constant.ERR_CODE_SUCCESS
}
