package usecase

import (
	"context"

	"github.com/stud-eng/stud-eng-api/internal/entity"
)

type TestUseCase struct {
	DBRepo DBRepository
}

func NewTestUseCase(dbRepo DBRepository) *TestUseCase {
	return &TestUseCase{
		DBRepo: dbRepo,
	}
}

func (uc *TestUseCase) PostTest(contex context.Context, tst *entity.Test) (*entity.Test, error) {
	return uc.DBRepo.InsertTest(contex, tst)
}
