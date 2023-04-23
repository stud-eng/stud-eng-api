package app

import (
	"fmt"

	"github.com/stud-eng/stud-eng-api/config"
)

func Run(conf *config.Config) error {
	fmt.Printf("DB-host :%s", conf.Database.Host)
	return nil
}
