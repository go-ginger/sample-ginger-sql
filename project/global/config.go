package global

import (
	gh "github.com/go-ginger/helpers"
	gm "github.com/go-ginger/models"
	"github.com/go-ginger/sql"
)

type config struct {
	gm.IConfig

	Port string

	// sql config
	SqlDialect          string
	SqlConnectionString string
	//
}

var Config *config

func initializeConfig() {
	Config = &config{
		Port:       gh.GetEnv("Port", "8080"),
		SqlDialect: gh.GetEnv("SqlDialect", "mysql"),
		SqlConnectionString: gh.GetEnv("SqlConnectionString",
			"root:123@/test?charset=utf8&parseTime=True&loc=UTC"),
	}
	sql.InitializeConfig(Config)
}
