package scrape

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func (repo *ScrapeRepository) GetMeaning(context context.Context) string {
	var err error
	repo.s.Doc, err = goquery.NewDocument("https://ejje.weblio.jp/content/test")
	if err != nil {
		panic("HTMLの取得に失敗しました")
	}
	//span.content-explanation.ej
	text := repo.s.Doc.Find("span.content-explanation.ej")

	//text := repo.s.Find("span.content-explanation.ej")
	fmt.Println(text.Text())
	return text.Text()
}
