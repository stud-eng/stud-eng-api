package dbrepo_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stud-eng/stud-eng-api/config"

	repo "github.com/stud-eng/stud-eng-api/internal/usecase/dbrepo"
	"github.com/stud-eng/stud-eng-api/pkg/db"
)

var rdb *db.DBHandler
var dbrepo *repo.DBRepository

// var goenv struct {
// 	Host string `envconfig:"DBHostName" env-default:"0.0.0.0"`
// 	Port string `envconfig:"DBPort" env-default:"3306"`
// 	User string `envconfig:"DBUser" env-default:"test"`
// 	Name string `envconfig:"DBName" env-default:"test"`
// 	Pass string `envconfig:"DBPpass" env-default:"password"`
// }

func TestMain(m *testing.M) {
	os.Exit(func() int {
		//ctx := context.Background()

		//read config
		// if err = envconfig.Process("", &goenv); err != nil {
		// 	panic(err)
		// }
		conf, err := config.New()
		if err != nil {
			fmt.Printf("conf error: %s", err)
		}

		for i := 0; i < 10; i++ {
			rdb, err = db.New(
				conf.Database.Host,
				conf.Database.Port,
				conf.Database.User,
				conf.Database.Name,
				conf.Database.Pass,
			)
			if err != nil && i == 10 {
				fmt.Println("app dbh error!!!!!!!!!!!!")
			}
			time.Sleep(1 * time.Second)
		}

		//create instance of dbrepo
		dbrepo = repo.New(rdb)
		fmt.Println("test db handler created!!!!!!!!!!!!!!")
		return m.Run()

	}())
}
