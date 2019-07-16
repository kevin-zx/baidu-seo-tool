package zhanzhang

import (
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/http-util"
	"log"
	"strings"
	"time"
)

const zhanzhangBaseUrl = "http://data.zz.baidu.com/urls?site=%s&token=%s"

type CommitResult struct {
	Remain      int      `json:"remain"`
	Success     int      `json:"success"`
	NotSameSite []string `json:"not_same_site"`
	NotValid    []string `json:"not_valid"`
}

func Commit(zhanZhangToken string, domain string, urls []string) (*CommitResult, error) {
	commitUrl := fmt.Sprintf(zhanzhangBaseUrl, domain, zhanZhangToken)
	log.Println(commitUrl)
	header := map[string]string{}
	zhanzhangReBack, err := http_util.GetWebConFromUrlWithAllArgs(commitUrl, header, "POST", []byte(strings.Join(urls, "\n")), time.Second*100)
	if err != nil {
		return nil, err
	}
	var zhangZhangResult CommitResult
	log.Println(zhanzhangReBack)
	err = json.Unmarshal([]byte(zhanzhangReBack), &zhangZhangResult)
	if err != nil {
		return nil, err
	}
	//log.Printf("commit %d urls to site_base %s, success %d, faild %d, remain %d/n", len(urls), domain, zhangZhangResult.Success, len(urls)-zhangZhangResult.Success, zhangZhangResult.Remain)
	return &zhangZhangResult, nil
}
