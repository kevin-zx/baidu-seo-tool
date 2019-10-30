package main

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
)

func main() {
	srs, err := search.GetBaiduPcResultsByKeyword("电话", 1, 10)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}

	srs, err = search.GetBaiduMobileResultsByKeyword("电话2", 1)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}
}
