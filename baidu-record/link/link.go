package link

import (
	"errors"
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/baidu"
	"net/url"
	"strings"
)

const baiduSearchAvailableLen = 76

func IsRecord(link string) (recordFlag bool, err error) {
	queryStr := ""
	if len(link) > baiduSearchAvailableLen {
		queryStr, err = handlerTooLongURL(link)
		if err != nil {
			return false, err
		}
	} else {
		queryStr = link
	}
	res, err := crawlerRecord2(queryStr)
	if strings.Contains(res, "没有找到该URL。您可以直接访问") || strings.Contains(res, "很抱歉，没有找到与") {
		return false, nil
	} else {
		return true, nil
	}
}

func crawlerRecord2(query string) (string, error) {
	return baidu.GetBaiduPCSearchHtml(query, 1)
}

const siteTemplate = "site:%s inurl:%s"

// site: inurl: 所用的字符串数
const siteTemplateLen = 12

func handlerTooLongURL(rawURL string) (string, error) {
	queryUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	domain := queryUrl.Host
	domainLen := len(domain)
	restLen := baiduSearchAvailableLen - siteTemplateLen - domainLen
	if restLen <= 0 {
		return "", errors.New(fmt.Sprintf("主域就已经过长了 查不出收录 url是：%s", rawURL))
	}
	//fmt.Println(len([]rune(rawURL))-restLen)
	//fmt.Println(len(rawURL) - restLen)
	restUrl := ""
	if len(rawURL) == len([]rune(rawURL)) {
		restUrl = string([]rune(rawURL)[len([]rune(rawURL))-restLen:])
	} else {
		restUrl = string([]rune(rawURL)[len([]rune(rawURL))-(restLen/2):])
	}

	return fmt.Sprintf(siteTemplate, domain, restUrl), nil
}
