// 这个类是获取百度内容用的
package search

import (
	"fmt"
	"github.com/kevin-zx/http-util"
	"net/url"
	"strings"
)

// 百度pc端
func GetBaiduPCSearchHtml(keyword string, page int) (string, error) {
	return GetBaiduPCSearchHtmlWithRN(keyword, page, 50)
}

func GetBaiduPCSearchHtmlWithRN(keyword string, page int, rn int) (string, error) {
	sUrl := combinePcSearchUrl(keyword, rn, page)
	webCon, err := http_util.GetWebConFromUrlWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
	if err != nil {
		return "", err
	}
	return webCon, nil
}

const PCSearchUrlBase = "https://www.baidu.com/s?wd=%s&rn=%d&pn=%d"

func combinePcSearchUrl(wd string, rn int, pageNum int) string {
	wd = escapeKeyword(wd)
	pn := rn * (pageNum - 1)
	PcSearchUrl := fmt.Sprintf(PCSearchUrlBase, wd, rn, pn)
	return PcSearchUrl
}

// 百度移动端
func GetBaiduMobileSearchHtml(keyword string, page int) (string, error) {
	sUrl := combineMobileUrl(keyword, page)
	//.GetWebConFromUrlWithHeader()
	webResponse, err := http_util.GetWebResponseFromUrlWithHeader(sUrl, map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Cookie":     "BAIDUID=EB9A2C07DAED944FBF46F430DA5A1064:FG=1; H_WISE_SIDS=130611_126126_132919_114552_136285_128070_126062_120209_134721_132910_131246_132378_131518_118881_118870_118852_118833_118795_107315_133351_129652_136193_132250_131861_128967_135308_135813_132552_135433_135873_129645_131423_134614_135100_135594_110085_136144_134154_127969_131951_135839_135457_127416_135866_135036_132468_135836_136260; rsv_i=b27bVq5oJ81lehNpAXYjqbjzTPRwQAxb5xRsJKJ0NxgoyqbHZ8jEeUycb1Ts%2Fs3pPylsc7XIB3iNHheJTBv6IfThsXJ7y9o; BDSVRTM=77",
	},
	)
	if err != nil {
		return "", err
	}
	webCon, err := http_util.ReadContentFromResponse(webResponse, "UTF-8")
	if err != nil {
		return "", err
	}
	return webCon, nil
}

const mobileSearchUrlBase = "https://www.baidu.com/from=844b/s?pn=%d&word=%s&ms=1"

func combineMobileUrl(keyword string, page int) string {
	keyword = escapeKeyword(keyword)
	pn := (page - 1) * 10
	mobileSearchUrl := fmt.Sprintf(mobileSearchUrlBase, pn, keyword)
	return mobileSearchUrl
}

func escapeKeyword(keyword string) string {
	keyword = url.QueryEscape(keyword)
	keyword = strings.Replace(keyword, "+", "%20", -1)
	return keyword
}

func GetBaiduPcResultsByKeyword(keyword string, page int, rn int) (baiduResults *[]SearchResult, err error) {
	webCon, err := GetBaiduPCSearchHtmlWithRN(keyword, page, rn)
	if err != nil {
		return
	}
	baiduResults, err = ParseBaiduPCSearchResultHtml(webCon)
	return
}

func GetBaiduMobileResultsByKeyword(keyword string, page int) (baiduResults *[]SearchResult, err error) {
	webCon, err := GetBaiduMobileSearchHtml(keyword, page)
	if err != nil {
		return
	}
	baiduResults, err = ParseBaiduMobileSearchResultHtml(webCon, page)
	return
}
