package zhanzhang

import (
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/http-util"
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

func Commit(zhanzhangToken string, domain string, urls []string) (*CommitResult, error) {
	commitUrl := fmt.Sprintf(zhanzhangBaseUrl, domain, zhanzhangToken)
	//log.Println(commitUrl)
	header := map[string]string{}
	zhanzhangReBack, err := http_util.GetWebConFromUrlWithAllArgs(commitUrl, header, "POST", []byte(strings.Join(urls, "\n")), time.Second*100)
	if err != nil {
		return nil, err
	}
	var zhanZhangResult CommitResult
	//log.Println(zhanzhangReBack)
	err = json.Unmarshal([]byte(zhanzhangReBack), &zhanZhangResult)
	if err != nil {
		return nil, err
	}
	//log.Printf("commit %d urls to site_base %s, success %d, faild %d, remain %d/n", len(urls), domain, zhanZhangResult.Success, len(urls)-zhanZhangResult.Success, zhanZhangResult.Remain)
	return &zhanZhangResult, nil
}
