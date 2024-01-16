package examples

import (
	client "github.com/hujinrun-github/go-press-test/client/interface"
)

type HttpClient struct {
	client.BaseClient
	model     *JsonModel
	statistic *SuccessStatistic
}
