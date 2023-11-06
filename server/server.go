package server

import (
	client "github.com/hujinrun-github/go-press-test/client/interface"
	"github.com/hujinrun-github/go-press-test/constant"
	model "github.com/hujinrun-github/go-press-test/model/interface"
	statistic "github.com/hujinrun-github/go-press-test/statistic/interface"
)

type Server struct {
	c        client.IClient
	s        statistic.IStatistic
	m        model.IModel
	parallel int
}

type Option func(*Server)

func WithClient(c client.IClient) Option {
	return func(s *Server) {
		s.c = c
	}
}

func WithStatistic(st statistic.IStatistic) Option {
	return func(s *Server) {
		s.s = st
	}
}

func WithParallel(p int) Option {
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

func (s *Server) Run() constant.ErrorCode {
	return constant.ERR_CODE_SUCCESS
}
