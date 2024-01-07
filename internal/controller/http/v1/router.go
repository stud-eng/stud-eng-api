package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stud-eng/stud-eng-api/internal/usecase"
)

func NewRouter(
	handler *gin.Engine,
	testUC *usecase.TestUseCase,
	scrapeUC *usecase.GetMeaningUseCase,
) error {
	h := handler.Group("/v1")
	{
		newTestRoute(h, testUC)
		newScrapeRoute(h, scrapeUC)

	}
	return nil
}
