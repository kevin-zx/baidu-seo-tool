package search

import (
	"fmt"
	"testing"
)

func TestParseBaiduPcSearchInfoFromHtml(t *testing.T) {
	searhHTML, err := GetBaiduPCSearchHtml("快餐陪送", 1)
	if err != nil {
		panic(err)
	}
	bi, err := ParseBaiduPcSearchInfoFromHtml(searhHTML)
	if err != nil {
		panic(err)
	}
	for _, sr := range *bi.SearchResults {

		fmt.Println(sr.CacheUrl)
	}
}
