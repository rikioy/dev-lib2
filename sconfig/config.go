package sconfig

import (
	"github.com/rikioy/dev-lib/senv"

	"github.com/cloudwego/kitex/pkg/klog"
)

var cfg *Sconfig

func init() {
	cfg = &Sconfig{}
	cfgName := "config"
	curEnv := senv.Env().Env
	if curEnv != senv.ENV_ONLINE {
		cfgName = "config_" + senv.Env().CurEnv()
	}
	klog.Info("set config name ", cfgName)
	cfg.Init(cfgName)
}

func Cfg() *Sconfig {
	return cfg
}
