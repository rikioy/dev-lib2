package senv

import (
	"log"

	flag "github.com/spf13/pflag"
)

const (
	ENV_DEV    = "dev"
	ENV_TEST   = "test"
	ENV_GRAY   = "gray"
	ENV_ONLINE = "online"
)

var env *Senv

func init() {
	env = &Senv{}
	flag.StringVar(&env.Env, "env", "dev", "current env from flag")
	flag.Parse()
	if env.Env != ENV_DEV && env.Env != ENV_TEST && env.Env != ENV_GRAY && env.Env != ENV_ONLINE {
		log.Fatal("--env must one of dev/test/gray/online")
	} else {
		log.Printf("current env is %s\n", env.Env)
	}
}

func Env() *Senv {
	return env
}
