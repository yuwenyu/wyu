package kernel

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	cacheKr *redis.Client
	redisMX	 sync.Mutex
)

type Kr interface {
	Get(key string) string
	Start() *r
	Engine() *redis.Client
	Close()
}

type r struct {
	kr *redis.Client
}

var _ Kr = &r{}

func NewRedis() *r {
	var object *r = &r{}
	if cacheKr != nil {
		object.kr = cacheKr
	}

	return object
}

func (thisKr *r) Start() *r {
	if thisKr.kr == nil {
		thisKr.instanceMaster()
	}

	return thisKr
}

func (thisKr *r) Engine() *redis.Client {
	if thisKr.kr == nil {
		thisKr.instanceMaster()
	}

	return thisKr.kr
}

func (thisKr *r) instanceMaster() *r {
	redisMX.Lock()
	defer redisMX.Unlock()

	if cacheKr != nil {
		thisKr.kr = cacheKr
		return thisKr
	}

	ros := thisKr.initialized()
	clientKr := redis.NewClient(&redis.Options{
		Addr:     ros.Addr,
		Password: ros.Password,
		DB:       ros.DB,
		PoolSize: ros.PoolSize,
	})

	_, err := clientKr.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	if cacheKr == nil {
		cacheKr = clientKr
	}

	thisKr.kr = clientKr

	return thisKr
}

func (thisKr *r) Get(key string) string {
	v, err := thisKr.kr.Get(key).Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	return v
}

func (thisKr *r) Close() {
	if thisKr.kr != nil {
		thisKr.kr.Close()
	}
}

type redisOptions struct {
	Addr 		string
	Password 	string
	DB 			int
	PoolSize	int
}

func (thisKr *r) initialized() *redisOptions {
	return &redisOptions{
		Addr:		"127.0.0.1:6379",
		Password:	"",
		DB:			0,
		PoolSize:	100,
	}
}