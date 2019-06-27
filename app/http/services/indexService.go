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

	return &indexService{
		kr:kernel.NewRedis().Start(),
		mConfigs:models.NewConfigsModel(db1.Engine()),
	}
}

func (s *indexService) FetchAll() []sys.Configs {
	//var db1 kernel.DB = kernel.NewDB(0)
	//s.mConfigs = models.NewConfigsModel(db1.Engine())
	//v := s.mConfigs.FetchAll()
	//db1.Engine().Close()
	//return v
	return s.mConfigs.FetchAll()

	//driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	"root", "root", "127.0.0.1", 3306, "sys")
	//engine, err := xorm.NewEngine("mysql", driverSource)
	//defer engine.Close()
	//
	//if err != nil {
	//	log.Fatalf("db.DbInstanceMaster,", err)
	//	return nil
	//}
	//
	//test := models.NewConfigsModel(engine)
	//
	//return test.FetchAll()
}

func (s *indexService) GetRedis() string {
	return s.kr.Get("test")
}


