/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	masterEngine *xorm.Engine
	mx		 	  sync.Mutex
)

type dataSource struct {
	dn string
	host string
	port int
	table string
	username string
	password string
}

func initDataSource(intSrc int) dataSource {
	var src string = strconv.Itoa(intSrc)
	var c INI = &ini{}
	c.LoadByFN("db")
	port, _ := c.K("db_engine_" + src, "port").Int()

	return dataSource{
		dn:c.K("db_engine_" + src, "driver").String(),
		host:c.K("db_engine_" + src, "host").String(),
		port:port,
		table:c.K("db_engine_" + src, "table").String(),
		username:c.K("db_engine_" + src, "username").String(),
		password:c.K("db_engine_" + src, "password").String(),
	}
}

func InstanceMaster(src int) *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	mx.Lock()
	defer mx.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	if src <= 0 {
		src = 1
	}

	ds := initDataSource(src)
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		ds.username, ds.password, ds.host, ds.port, ds.table)
	engine, err := xorm.NewEngine(ds.dn, driverSource)

	if err != nil {
		log.Fatalf("db.DbInstanceMaster,", err)
		return nil
	}

	engine.ShowSQL(true)
	engine.SetTZDatabase(SysTimeLocation)

	// SQL Cache
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	// engine.SetDefaultCacher(cacher)

	masterEngine = engine

	return engine
}