package kernel


//"log"
//"sync"
//"fmt"
//"github.com/go-xorm/xorm"
//_ "github.com/go-sql-driver/mysql"
//"salamander/configs"
//import (
//	"fmt"
//	"log"
//	"sync"
//)
//
//var (
//	masterEngine *xorm.Engine
//	slaverEngine *xorm.Engine
//	lock		 sync.Mutex
//)
//
//func InstanceMaster() *xorm.Engine {
//	if masterEngine != nil {
//		return masterEngine
//	}
//
//	lock.Lock()
//	defer lock.Unlock()
//
//	if masterEngine != nil {
//		return masterEngine
//	}
//
//	c := configs.MasterDbConfig
//	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
//		c.User, c.Pwd, c.Host, c.Port, c.Table)
//	engine, err := xorm.NewEngine(configs.DriverName, driverSource)
//
//	if err != nil {
//		log.Fatalf("dbhelper.DbInstanceMaster,", err)
//		return nil
//	}
//
//	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
//	engine.ShowSQL(true)
//	engine.SetTZDatabase(configs.SysTimeLocation)
//
//	//性能优化的时候才考虑，加上本机的SQL缓存 SQL Cache
//	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
//	// engine.SetDefaultCacher(cacher)
//
//	masterEngine = engine
//
//	return engine
//}
//
//func InstanceSlaver() *xorm.Engine {
//	if slaverEngine != nil {
//		return slaverEngine
//	}
//
//	lock.Lock()
//	defer lock.Unlock()
//
//	if slaverEngine != nil {
//		return slaverEngine
//	}
//
//	c := configs.SlaverDbConfig
//	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
//		c.User, c.Pwd, c.Host, c.Port, c.Table)
//	engine, err := xorm.NewEngine(configs.DriverName, driverSource)
//
//	if err  != nil {
//		log.Fatalf("dbhelper.DbInstanceMaster", err)
//	}
//
//	engine.SetTZLocation(configs.SysTimeLocation)
//
//	slaverEngine = engine
//
//	return engine
//}



//repository repositories
//import (
//"log"
//"github.com/go-xorm/xorm"
//"salamander/models"
//)
//
//type TestDao struct {
//	engine *xorm.Engine
//}
//
//func NewTestDao(engine *xorm.Engine) *TestDao {
//	return &TestDao{
//		engine: engine,
//	}
//}
//
//func (d *TestDao) GetAll() []models.Test {
//	dl := make([]models.Test, 0)
//	log.Println("-]",dl,"[-")
//	if err := d.engine.Find(&dl); err != nil {
//		log.Printf("Error DB: %s", err)
//		return nil
//	} else {
//		return dl
//	}
//}

//services
//package services
//
//import (
//"log"
//"salamander/dao"
//"salamander/datasource"
//"salamander/models"
//)
//
//type IndexService interface {
//	GetAll() []models.Test
//	FindAll() []models.Test
//}
//
//type indexService struct {
//	dao *dao.TestDao
//}
//
//func NewIndexService() *indexService {
//	return &indexService{
//		dao: dao.NewTestDao(datasource.InstanceMaster()),
//	}
//}
//
//func (s *indexService) GetAll() []models.Test {
//	log.Println(s.dao.GetAll())
//	return s.dao.GetAll()
//}
//
//func (s *indexService) FindAll() []models.Test {
//	return s.dao.GetAll()
//}