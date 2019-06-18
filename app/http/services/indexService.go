package services

import (
	"github.com/yuwenyu/kernel"
	"wyu/app/models"
	"wyu/app/repositories/sys"
)

type IndexService interface {
	FetchAll() []sys.Configs
}

type indexService struct {
	mConfigs *models.ConfigsModel
}

var _ IndexService = &indexService{}

func NewIndexService() *indexService {
	return &indexService{
		mConfigs:models.NewConfigsModel(kernel.InstanceMaster(0)),
	}
}

func (this *indexService) FetchAll() []sys.Configs {
	return this.mConfigs.FetchAll()
}


