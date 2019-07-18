package middleware

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"

	"wyu/config"
)

var mapMSession map[string]interface{}

func init() {
	if mapMSession == nil {
		initMSessionParams()
	}
}

type Session interface {
	SetKeyPrefix(keyPrefix string) *session
	Start() redis.Store
}

type session struct {
	keyPrefix 	string
}

var _ Session = &session{}

func NewSession() *session {
	return &session{}
}

func (s *session) SetKeyPrefix(keyPrefix string) *session {
	s.keyPrefix = keyPrefix
	return s
}

func (s *session) Start() redis.Store {

	store, err := redis.NewStore(
		mapMSession["redisPool"].(int),
		mapMSession["redisNetwork"].(string),
		mapMSession["redisAddr"].(string),
		mapMSession["redisPassword"].(string),
		[]byte(mapMSession["redisKeyPairs"].(string)),
	)

	if err != nil {
		panic(err.Error())
	}

	store.Options(sessions.Options{
		MaxAge: int(mapMSession["storeMaxAge"].(time.Duration) * time.Minute),
		Path:   mapMSession["storePath"].(string),
	})

	if s.keyPrefix != "" {
		redis.SetKeyPrefix(store, s.keyPrefix)
	}

	return store
}

func MSession() gin.HandlerFunc {
	var s Session = NewSession()
	s.SetKeyPrefix(mapMSession["keyPrefix"].(string))

	return sessions.Sessions(mapMSession["keySID"].(string), s.Start())
}

type SessionInitialized interface {
	Default() sessions.Session
}
type sessionInitialized struct {
	c *gin.Context
}

var _ SessionInitialized = &sessionInitialized{}

func NewSessionInitialized(c *gin.Context) *sessionInitialized {
	return &sessionInitialized{c:c}
}

func (s *sessionInitialized) Default() sessions.Session {
	return sessions.Default(s.c)
}

func initMSessionParams() {
	var c kernel.INI = kernel.NewIni().LoadByFN(config.ConfSession)

	keySID 			:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][0],
	).String()
	if keySID == "" {keySID = config.MSessionKeySID}

	keyPrefix		:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][1],
	).String()
	if keyPrefix == "" {keyPrefix = config.MSessionKeyPrefix}

	redisNetwork	:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][3],
	).String()
	if redisNetwork == "" {redisNetwork = config.MSessionRedisNetwork}

	redisAddr		:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][4],
	).String()
	if redisAddr == "" {
		panic("Middleware Session Redis Address can't be empty")
	}

	redisPassword	:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][5],
	).String()

	redisKeyPairs	:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][6],
	).String()
	if redisKeyPairs == "" {redisKeyPairs = config.MSessionRedisKeyPairs}

	redisPool, err	:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][2],
	).Int()
	if err != nil {redisPool = config.MSessionRedisPool}

	storeMaxAge, err:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][7],
	).Duration()
	if err != nil {storeMaxAge = config.MSessionStoreMaxAge}

	storePath		:= c.K(
		config.MapConfLists[config.ConfSession][0],
		config.MapConfParam[config.MapConfLists[config.ConfSession][0]][8],
	).String()
	if storePath == "" {storePath = config.MSessionStorePath}

	mapMSession = map[string]interface{}{
		"keySID":keySID,
		"keyPrefix":keyPrefix,
		"redisPool":redisPool,
		"redisNetwork":redisNetwork,
		"redisAddr":redisAddr,
		"redisPassword":redisPassword,
		"redisKeyPairs":redisKeyPairs,
		"storeMaxAge":storeMaxAge,
		"storePath":storePath,
	}
}