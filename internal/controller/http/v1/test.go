package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stud-eng/stud-eng-api/internal/controller/http/model"
	"github.com/stud-eng/stud-eng-api/internal/usecase"
)

type testRoute struct {
	testUC *usecase.TestUseCase
}

func newTestRoute(handler *gin.RouterGroup, testUC *usecase.TestUseCase) {
	repo := &testRoute{testUC: testUC}
	{
		handler.POST("/test", repo.postTest)
	}
}

func (repo *testRoute) postTest(c *gin.Context) {
	tst, appErr := model.ValidateTest(c)
	if appErr != nil {
		fmt.Println("validation 失敗")
	}

	res, appErr := repo.testUC.PostTest(c.Request.Context(), tst.ToEntity())
	if appErr != nil {
		fmt.Println("postTest Error!!!1")
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
