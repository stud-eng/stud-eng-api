package model

import (

	//validation "github.com/go-ozzo/ozzo-validation"
	"fmt"

	"github.com/stud-eng/stud-eng-api/internal/entity"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}

func BindTest(c *gin.Context) (*Test, error) {
	var res Test

	if err := c.BindJSON(&res); err != nil {
		fmt.Println(res)
		fmt.Println("bindJSon error")
		return nil, err
	}
	return &res, nil

}

// func (t Test) Validate() error {
// 	err := validation.Errors{

// 	}
// }

func ValidateTest(c *gin.Context) (*Test, error) {
	tst, err := BindTest(c)
	if err != nil {
		return nil, err
	}
	return tst, nil
}

func (t *Test) ToEntity() *entity.Test {
	return &entity.Test{
		Name:     t.Name,
		Mail:     t.Name,
		Password: t.Password,
	}
}
