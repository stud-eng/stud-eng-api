package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stud-eng/stud-eng-api/internal/usecase"
)

type scrapeRoute struct {
	getMeaningUC *usecase.GetMeaningUseCase
}

func newScrapeRoute(handler *gin.RouterGroup, getMeaningUC *usecase.GetMeaningUseCase) {
	repo := &scrapeRoute{getMeaningUC: getMeaningUC}
	{
		handler.GET("/scrape", repo.getMeaning)
	}
}

func (repo *scrapeRoute) getMeaning(c *gin.Context) {
	meaning := repo.getMeaningUC.GetMeaning(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"message": meaning})
}
