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
}

type db struct {
	id int
	engine *xorm.Engine
	mx sync.Mutex
}

var _ DB = &db{}

func NewDB(src int) *db {
	if src <= 0 { src = 1 }

	var odb *db
	if v, ok := dbEngines[src]; ok {
		odb = &v
	} else {
		odb = &db{id:src}
	}

	return odb
}

func (odbc *db) Engine() *xorm.Engine {
	if odbc.engine == nil {
		odbc.instanceMaster()
	}

	return odbc.engine
}

func (odbc *db) instanceMaster() *db {
	odbc.mx.Lock()
	defer odbc.mx.Unlock()

	if odbc.engine != nil {
		return odbc
	}

	if len(dbEngines) == 0 {
		dbEngines = make(map[int]db)
	} else {
		if v, ok := dbEngines[odbc.id]; ok {
			if odbc.engine == nil {
				odbc.engine = v.engine
			}

			return odbc
		}
	}

	ds := odbc.initDataSource()
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

	odbc.engine = engine

	dbEngines[odbc.id] = *odbc

	return odbc
}

type dataSource struct {
	dn 		 string
	host 	 string
	port 	 int
	table 	 string
	username string
	password string
}

func (odbc *db) initDataSource() *dataSource {
	var key string = "db_engine" + StrUL + strconv.Itoa(odbc.id)
	var c INI = NewIni().LoadByFN("db")

	dn 		:= c.K(key, "driver").String()
	host 	:= c.K(key, "host").String()
	port, _ := c.K(key, "port").Int()
	table 	:= c.K(key, "table").String()
	username:= c.K(key, "username").String()
	password:= c.K(key, "password").String()

	return &dataSource{
		dn:			dn,
		host:		host,
		port:		port,
		table:		table,
		username:	username,
		password:	password,
	}
}