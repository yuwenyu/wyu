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
	var db1 kernel.DB = kernel.NewDB(0)
	var db2 kernel.DB = kernel.NewDB(2)
	db1.Engine()
	db2.Engine()
	return &indexService{
		kr:kernel.NewRedis(),
		mConfigs:models.NewConfigsModel(db1.Engine()),
	}
}

func (s *indexService) FetchAll() []sys.Configs {
	return s.mConfigs.FetchAll()
}

func (s *indexService) GetRedis() string {
	return s.kr.Get("test")
}


