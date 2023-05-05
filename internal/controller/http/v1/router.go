package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stud-eng/stud-eng-api/internal/usecase"
)

func NewRouter(
	handler *gin.Engine,
	testUC *usecase.TestUseCase,
) error {
	h := handler.Group("/v1")
	{
		newTestRoute(h, testUC)

	}
	return nil
}
