/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	masterEngine *xorm.Engine
	slaverEngine *xorm.Engine
	mx		 	  sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	mx.Lock()
	defer mx.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	var c INI = &ini{}
	c.LoadByFN("db")

	host := c.K("db_master","host").String()
	port, _ := c.K("db_master","port").Int()
	table:= c.K("db_master","table").String()
	username := c.K("db_master","username").String()
	password := c.K("db_master","password").String()

	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, host, port, table)
	engine, err := xorm.NewEngine("mysql", driverSource)

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

func InstanceSlaver() *xorm.Engine {
	if slaverEngine != nil {
		return slaverEngine
	}

	mx.Lock()
	defer mx.Unlock()

	if slaverEngine != nil {
		return slaverEngine
	}

	var c INI = &ini{}
	c.LoadByFN("db")

	host := c.K("db_slaver","host").String()
	port, _ := c.K("db_slaver","port").Int()
	table:= c.K("db_slaver","table").String()
	username := c.K("db_slaver","username").String()
	password := c.K("db_slaver","password").String()

	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, host, port, table)
	engine, err := xorm.NewEngine("mysql", driverSource)

	if err != nil {
		log.Fatalf("db.DbInstanceSlaver", err)
	}

	engine.SetTZLocation(SysTimeLocation)
	slaverEngine = engine

	return engine
}