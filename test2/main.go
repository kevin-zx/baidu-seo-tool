package main

import (
	"fmt"
	http_util "github.com/kevin-zx/http-util"
	"time"
)

func main() {
	con, err := http_util.GetWebConFromUrlWithAllArgs("http://www.szcppf.com/list-4-1.html", map[string]string{"User-Agent": "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"}, "GET", nil, time.Second*10)
	if err != nil {
		panic(err)
	}
	fmt.Println(con)
}
