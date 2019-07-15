package search

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type BaiduSearchInfo struct {
	Port            string
	BaiduMatchCount int
	MainPageCount   int
	IsEscape        bool
	EscapeWord      string
	SearchResults   *[]SearchResult
}

func ParseBaiduPcSearchInfoFromHtml(html string) (bsi *BaiduSearchInfo, err error) {
	bsi = &BaiduSearchInfo{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}
	t := doc.Find("div.nums>span.nums_text").Text()
	if t == "" {
		bsi.BaiduMatchCount = -1
	}
	t = strings.Replace(t, "百度为您找到相关结果约", "", -1)
	t = strings.Replace(t, "个", "", -1)
	t = strings.Replace(t, ",", "", -1)
	bsi.BaiduMatchCount, err = strconv.Atoi(t)

	seTip := doc.Find("#super_se_tip").Text()
	if strings.Contains(seTip, "已显示") {
		doc.Find("#super_se_tip strong").Each(func(_ int, strong *goquery.Selection) {
			_, ok := strong.Attr("class")
			if !ok {
				bsi.IsEscape = true
				bsi.EscapeWord = strong.Text()
			}
		})
	}
	srs, err := ParseBaiduPCSearchResultHtml(html)
	if err != nil {
		return
	}
	bsi.SearchResults = srs
	for _, sr := range *bsi.SearchResults {
		if sr.IsHomePage() {
			bsi.MainPageCount++
		}
	}
	return
}
