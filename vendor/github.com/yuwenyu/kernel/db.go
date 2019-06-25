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

type DB interface {
	Engine() *xorm.Engine
	Test() int
}

type db struct {
	id int
	engine *xorm.Engine
	mx sync.Mutex
}

var _ DB = &db{}

func NewDB(src int) *db {
	if src <= 0 {
		src = 1
	}

	return &db{id:src}
}

func (db *db) Engine() *xorm.Engine {
	if db.engine == nil {
		db.instanceMaster()
		fmt.Println("--- Initialized Engine DB Source New Initialized! ---")
	}

	return db.engine
}

func (db *db) Test() int {
	return db.id
}

func (db *db) instanceMaster() *db {
	db.mx.Lock()
	defer db.mx.Unlock()

	if db.engine != nil {
		fmt.Println("--- Initialized Engine DB Source New Exist ---")
		return db
	}

	ds := db.initialized()
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		ds.username, ds.password, ds.host, ds.port, ds.table)
	engine, err := xorm.NewEngine(ds.dn, driverSource)

	if err != nil {
		log.Fatalf("db.DbInstanceMaster,", err)
		return nil
	}

	engine.ShowSQL(false)
	engine.SetTZDatabase(SysTimeLocation)

	// SQL Cache
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	// engine.SetDefaultCacher(cacher)

	fmt.Println("--- Initialized Engine DB Source New ---")
	db.engine = engine

	return db
}

type dataSource struct {
	dn string
	host string
	port int
	table string
	username string
	password string
}

func (db *db) initialized() dataSource {
	var src string = strconv.Itoa(db.id)
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