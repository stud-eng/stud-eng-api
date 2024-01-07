package usecase

import (
	"context"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

type (
	DBRepository interface {
		GetTest(context context.Context, id uint32) (*entity.Test, error)
		InsertTest(context context.Context, test *entity.Test) (*entity.Test, error)
		UpdateTest(context context.Context, test *entity.Test) (*entity.Test, error)
		DeleteTest(context context.Context, id uint32) error
	}

	ScrapeRepository interface {
		GetMeaning(context context.Context) string
	}
)
