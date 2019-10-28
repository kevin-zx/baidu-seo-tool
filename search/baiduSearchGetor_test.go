package search

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetBaiduMobileResultsByKeyword(t *testing.T) {
	wec, err := GetBaiduPCSearchHtmlWithRNAndTimeDayInterval("石家庄", 1, 50, "2019-10-26")
	if err != nil {
		panic(err)
	}
	rs, err := ParseBaiduPCSearchResultHtml(wec)
	if err != nil {
		panic(err)
	}
	for _, r := range *rs {
		fmt.Println(r.BaiduURL)
		baiduUrl, err := url.Parse(r.BaiduURL)
		if err != nil {
			panic(err)
		}
		fmt.Println(baiduUrl.Query().Get("lid"))
	}

}
