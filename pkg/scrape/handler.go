package scrape

import (
	"github.com/PuerkitoBio/goquery"
)

type ScrapeHandler struct {
	Doc *goquery.Document
}

func New() (*ScrapeHandler, error) {
	var err error
	scrapeHandler := new(ScrapeHandler)
	if err != nil {
		return nil, err
	}
	return scrapeHandler, nil
}
