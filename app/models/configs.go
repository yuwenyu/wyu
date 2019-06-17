package models

import(
	"log"

	"github.com/go-xorm/xorm"
	"wyu/app/repositories/sys"
)

type ConfigsModel struct {
	engine *xorm.Engine
}

func NewConfigsModel(engine *xorm.Engine) *ConfigsModel {
	return &ConfigsModel{engine:engine}
}

func (m *ConfigsModel) FetchAll() []sys.Configs {
	ds := make([]sys.Configs, 0)

	if err := m.engine.Find(&ds); err != nil {
		log.Println("Error DB: %s", err.Error())
		return nil
	} else {
		return ds
	}
}
