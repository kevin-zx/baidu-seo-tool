package search

import (
	"fmt"
	"testing"
)

func TestParseBaiduPcSearchInfoFromHtml(t *testing.T) {
	searhHTML, err := GetBaiduPCSearchHtml("承德众鼎致远电气安装有限公司", 1)
	if err != nil {
		panic(err)
	}
	bi, err := ParseBaiduPcSearchInfoFromHtml(searhHTML)
	if err != nil {
		panic(err)
	}
	for _, sr := range *bi.SearchResults {

		fmt.Println(sr.BaiduDescription)
	}
}
