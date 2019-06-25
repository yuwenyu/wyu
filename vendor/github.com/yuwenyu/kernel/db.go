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

var dbEngines map[int]db

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

	var odb *db
	if v, ok := dbEngines[src]; ok {
		odb = &v
		fmt.Println("initialized global")
	} else {
		odb = &db{id:src}
		fmt.Println("initialized New")
	}

	return odb
}

func (tdb *db) Engine() *xorm.Engine {
	if tdb.engine == nil {
		tdb.instanceMaster()
	}

	return tdb.engine
}

func (tdb *db) Test() int {
	return tdb.id
}

func (tdb *db) instanceMaster() *db {
	tdb.mx.Lock()
	defer tdb.mx.Unlock()

	if tdb.engine != nil {
		return tdb
	}

	ds := tdb.initialized()
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

	tdb.engine = engine

	dbEngines := make(map[int]db)
	dbEngines[tdb.id] = *tdb

	return tdb
}

type dataSource struct {
	dn string
	host string
	port int
	table string
	username string
	password string
}

func (tdb *db) initialized() dataSource {
	var src string = strconv.Itoa(tdb.id)
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