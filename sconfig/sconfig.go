package sconfig

import "github.com/spf13/viper"

type Sconfig struct {
	cfg *viper.Viper
}

func (s *Sconfig) Init(cfgName string) (err error) {
	s.cfg = viper.NewWithOptions()
	s.cfg.SetConfigName(cfgName)
	s.cfg.AddConfigPath(".")
	s.cfg.AddConfigPath("./config")
	err = s.cfg.ReadInConfig()
	return
}

func (s *Sconfig) GetString(key string) string {
	return s.cfg.GetString(key)
}

func (s *Sconfig) GetInt(key string) int {
	return s.cfg.GetInt(key)
}

func (s *Sconfig) GetInt32(key string) int32 {
	return s.cfg.GetInt32(key)
}

func (s *Sconfig) GetInt64(key string) int64 {
	return s.cfg.GetInt64(key)
}

func (s *Sconfig) GetStringMapString(key string) map[string]string {
	return s.cfg.GetStringMapString(key)
}
