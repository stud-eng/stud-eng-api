package usecase

import (
	"context"
)

type GetMeaningUseCase struct {
	repo ScrapeRepository
}

func NewGetMeaningUseCase(scrapeRepo ScrapeRepository) *GetMeaningUseCase {
	return &GetMeaningUseCase{
		repo: scrapeRepo,
	}
}

func (uc *GetMeaningUseCase) GetMeaning(contex context.Context) string {
	return uc.repo.GetMeaning(contex)
}
