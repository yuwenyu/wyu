package services

import (
	"github.com/yuwenyu/kernel"
	"wyu/app/models"
	"wyu/app/repositories/sys"
)

type IndexService interface {
	FetchAll() []sys.Configs
	GetRedis() string
}

type indexService struct {
	kr kernel.Kr
	mConfigs *models.ConfigsModel
}

var _ IndexService = &indexService{}

func NewIndexService() *indexService {
	return &indexService{
		kr:kernel.NewRedis(),
		mConfigs:models.NewConfigsModel(kernel.NewDB(0).Engine()),
	}
}

func (s *indexService) FetchAll() []sys.Configs {
	return s.mConfigs.FetchAll()
}

func (s *indexService) GetRedis() string {
	return s.kr.Get("test")
}


