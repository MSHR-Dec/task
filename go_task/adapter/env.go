package adapter

import (
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	Env              string   `default:"local"`
	CorsAllowOrigins []string `default:"http://localhost:3000,http://frontend:3000" split_words:"true"`
	MysqlUser        string   `default:"task" split_words:"true"`
	MysqlPassword    string   `default:"task" split_words:"true"`
	MysqlDatabase    string   `default:"task" split_words:"true"`
	MysqlHost        string   `default:"mysql:3306" split_words:"true"`
	RedisHost        string   `default:"redis:6379" split_words:"true"`
}

var Environment environment

func init() {
	if err := envconfig.Process("", &Environment); err != nil {
		panic("failed to read environment variables")
	}
}
