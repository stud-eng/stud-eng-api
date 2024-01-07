package scrape

import (
	"github.com/stud-eng/stud-eng-api/pkg/scrape"
)

type ScrapeRepository struct {
	s *scrape.ScrapeHandler
}

func NewScrape(s *scrape.ScrapeHandler) *ScrapeRepository {
	return &ScrapeRepository{
		s: s,
	}
}
