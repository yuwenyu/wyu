package kernel

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	krc *redis.Client
	mxr	 sync.Mutex
)

type Kr interface {
	Get(key string) string
	Close()
}

type r struct {
	kr *redis.Client
}

var _ Kr = &r{}

type redisOptions struct {
	Addr 		string
	Password 	string
	DB 			int
	PoolSize	int
}

func initialized() *redisOptions {
	return &redisOptions{
		Addr:		"127.0.0.1:6379",
		Password:	"",
		DB:			0,
		PoolSize:	10,
	}
}

func NewRedis() *r {
	if krc == nil {
		ros := initialized()
		krc  = redis.NewClient(&redis.Options{
			Addr:     ros.Addr,
			Password: ros.Password,
			DB:       ros.DB,
			PoolSize: ros.PoolSize,
		})

		fmt.Println("--- Initialized Engine Redis ---")
	}

	_, err := krc.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	mxr.Lock()
	defer mxr.Unlock()

	return &r{kr:krc}
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