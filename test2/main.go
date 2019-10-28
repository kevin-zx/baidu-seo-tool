package main

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
)

func main() {
	srs, err := search.GetBaiduPcResultsByKeyword("site:qichacha.com 电话 18852997668", 10, 10)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}

	srs, err = search.GetBaiduMobileResultsByKeyword("site:qichacha.com 18852997668", 1)
	if err != nil {
		panic(err)
	}
	for _, sr := range *srs {
		fmt.Printf("%v\n", sr)
	}
}
