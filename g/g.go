package g

import (
	"github.com/rikioy/dev-lib/sconfig"
	"github.com/rikioy/dev-lib/senv"
)

func Config() *sconfig.Sconfig {
	return sconfig.Cfg()
}

func Cfg() *sconfig.Sconfig {
	return Config()
}

func Env() *senv.Senv {
	return senv.Env()
}
