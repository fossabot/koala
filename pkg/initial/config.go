package initial

import (
	"github.com/mkzilla/koala/pkg/types/config"
)

func InitConfig(filepath string) {
	cfg := new(config.Config)

	err := cfg.LoadFromJson(filepath)
	if err != nil {
		panic(err)
	}
	cfg.BasicRegister.Template = cfg.Read(cfg.BasicRegister.Template)
	config.Configs = cfg
}
