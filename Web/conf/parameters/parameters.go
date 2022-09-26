package parameters

import (
	"gopkg.in/ini.v1"
)

var ParamsService Params

type Params struct {
	PasswordAES string

	NatsHost string
	NatsPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func (params *Params) LoadINI(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	params.PasswordAES = cfg.Section("AES").Key("PasswordAES").String()

	params.DBHost = cfg.Section("DB").Key("Host").String()
	params.DBPort = cfg.Section("DB").Key("Port").String()
	params.DBUser = cfg.Section("DB").Key("User").String()
	params.DBPassword = cfg.Section("DB").Key("Password").String()
	params.DBName = cfg.Section("DB").Key("Name").String()

	params.NatsHost = cfg.Section("NATS").Key("Host").String()
	params.NatsPort = cfg.Section("NATS").Key("Port").String()
	return nil
}
