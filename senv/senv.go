package senv

type Senv struct {
	Env string
}

func (s *Senv) CurEnv() string {
	return s.Env
}

func (s *Senv) IsOnline() bool {
	if s.Env == ENV_ONLINE {
		return true
	} else {
		return false
	}
}
