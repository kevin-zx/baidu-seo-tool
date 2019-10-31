package main

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
)

func main() {
	srs, err := search.GetBaiduPcResultsByKeyword("site:www.ynhande.cn 云南安全鞋批发", 1, 10)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}

	srs, err = search.GetBaiduMobileResultsByKeyword("site:www.ynhande.cn 云南安全鞋批发", 1)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}
}
