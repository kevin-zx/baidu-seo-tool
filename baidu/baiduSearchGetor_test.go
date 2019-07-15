package baidu

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetBaiduMobileResultsByKeyword(t *testing.T) {
	rs, err := GetBaiduMobileResultsByKeyword("测试", 1)
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
