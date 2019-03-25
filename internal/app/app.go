package app

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"log"
	"os"
	"sync"
)

type app struct {
	Cfg *Config
}

func (t *app) InitLog() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if !t.debug {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	//log.Debug().Msg("init redis")
	//t.redis = redis.NewClient(utils.ParseRedisUrl(t.RedisUrl))
	//utils.MustNotError(t.redis.Ping().Err())

	t.id = utils.IpAddress()
	if t.id == "" {
		panic("获取不到ip地址")
	}

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service_name", "mworker").
		Str("service_ip", t.id).
		Str("service_id", t.id).
		Bool("is_debug", t.debug).
		Caller().
		Logger()
}

var _app *app
var once sync.Once

func DefaultApp() *app {

	once.Do(func() {
		_app = &app{}
	})
	return _app
}
