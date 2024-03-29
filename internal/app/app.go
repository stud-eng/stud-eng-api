package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stud-eng/stud-eng-api/config"
	v1 "github.com/stud-eng/stud-eng-api/internal/controller/http/v1"
	"github.com/stud-eng/stud-eng-api/internal/usecase"
	dbRepo "github.com/stud-eng/stud-eng-api/internal/usecase/dbrepo"
	scrapeRepo "github.com/stud-eng/stud-eng-api/internal/usecase/scrape"
	"github.com/stud-eng/stud-eng-api/pkg/db"
	"github.com/stud-eng/stud-eng-api/pkg/scrape"
)

func Run(conf *config.Config) error {
	fmt.Printf("DB-host :%s", conf.Database.Host)
	dbh, err := db.New(
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Name,
		conf.Database.Pass,
	)
	if err != nil {
		fmt.Println("app dbh error!!!!!!!!!!!!")
		return err
	}
	fmt.Println("db handler created!!!!!!!!!!!!!!")

	scrapeh, err := scrape.New()

	if err != nil {
		fmt.Println("scrape handler error!!!!!!!!!!!!")
		return err
	}

	dbrepo := dbRepo.New(dbh)
	scrapeRepo := scrapeRepo.NewScrape(scrapeh)

	testUC := usecase.NewTestUseCase(dbrepo)
	getMeaningUC := usecase.NewGetMeaningUseCase(scrapeRepo)

	handler := gin.New()
	if err := v1.NewRouter(
		handler,
		testUC,
		getMeaningUC,
	); err != nil {
		return fmt.Errorf("init router failed!!")
	}

	if handlerErr := handler.Run(":8080"); handlerErr != nil {
		fmt.Println("handler run error !!!!")
	}
	return nil
}
