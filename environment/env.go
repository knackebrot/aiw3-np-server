package environment

import (
	"github.com/knackebrot/aiw3-np-server/config"
	"github.com/eaigner/jet"
	"github.com/go-redis/redis"
)

type Environment struct {
	Config   *config.Config
	Database *jet.Db
	Redis    *redis.Client
}

var Env *Environment

func SetEnvironment(env *Environment) {
	Env = env
}
