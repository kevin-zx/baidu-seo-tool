// 这个类是获取百度内容用的
package search

import (
	"fmt"
	"github.com/kevin-zx/http-util"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 百度pc端
func GetBaiduPCSearchHtml(keyword string, page int) (string, error) {
	return GetBaiduPCSearchHtmlWithRN(keyword, page, 50)
}

func GetBaiduPCSearchHtmlWithRN(keyword string, page int, rn int) (string, error) {
	sUrl := combinePcSearchUrl(keyword, rn, page)
	webCon, err := http_util.GetWebConFromUrlWithHeader(sUrl,
		map[string]string{
			"Connection":                " keep-alive",
			"Cache-Control":             " max-age=0",
			"Upgrade-Insecure-Requests": " 1",
			"User-Agent":                " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
			"Sec-Fetch-Mode":            " navigate",
			"Sec-Fetch-User":            " ?1",
			"Accept":                    " text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Sec-Fetch-Site":            " none",
			"Accept-Encoding":           " gzip, deflate, br",
			"Accept-Language":           " en,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,zh-TW;q=0.6",
			"Cookie":                    " BAIDUID=8829A73D6D0F4DF82E34A5B447352469:FG=1; BIDUPSID=8829A73D6D0F4DF82E34A5B447352469; PSTM=1567864240; MSA_WH=360_640; H_WISE_SIDS=136262_126008_100805_133333_134982_120141_136986_132910_136456_136620_131247_136722_132378_131517_118885_118875_118857_118823_118798_107313_132783_136800_136431_133351_136862_136818_137012_129652_136195_132250_124634_135308_133847_132552_129645_131423_136220_134318_136165_110085_134254_134153_127969_131754_131953_136614_135416_135458_127416_134935_136636_131545_134353_132467_136413_100457; BD_UPN=12314753; BDUSS=n5iNTVFT3Jrem1uWnRISjZDR2xQMXZZczc0VlpiZ2xCWmRSfjBoWWtkZDA1TkJkSVFBQUFBJCQAAAAAAAAAAAEAAACgp0IQYTUxOTc1NDgyMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHRXqV10V6lda; MCITY=-%3A; BD_HOME=1; H_PS_PSSID=1448_21084_29910_29567_29700_29220_22160; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; BD_CK_SAM=1; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; sug=3; sugstore=1; ORIGIN=0; bdime=0; H_PS_645EC=9da1yRjLjuL8qUx2Hxi4YkTtpfGjYzlNR%2Bf3jmrw3bcKnhEdy%2Bh4b1znz6k",
		},
	)
	if err != nil {
		return "", err
	}
	return webCon, nil
}

func GetBaiduPCSearchHtmlWithRNAndTimeDayInterval(keyword string, page int, rn int, startDate string) (string, error) {
	sUrl := combinePcSearchUrl(keyword, rn, page)
	startDateUnix, err := time.ParseInLocation("2006-01-02 15:04:05", startDate+" 00:00:00", time.Local)
	if err != nil {
		return "", err
	}
	endDateUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate+" 23:59:59", time.Local)

	sUrl += "&gpc=stf%3D" + strconv.Itoa(int(startDateUnix.Unix())) + "%2C" + strconv.Itoa(int(endDateUnix.Unix())) + "%7Cstftype%3D2"
	webCon, err := http_util.GetWebConFromUrlWithHeader(sUrl,
		map[string]string{
			"Connection":                " keep-alive",
			"Cache-Control":             " max-age=0",
			"Upgrade-Insecure-Requests": " 1",
			"User-Agent":                " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
			"Sec-Fetch-Mode":            " navigate",
			"Sec-Fetch-User":            " ?1",
			"Accept":                    " text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Sec-Fetch-Site":            " none",
			"Accept-Encoding":           " gzip, deflate, br",
			"Accept-Language":           " en,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,zh-TW;q=0.6",
			"Cookie":                    " BAIDUID=8829A73D6D0F4DF82E34A5B447352469:FG=1; BIDUPSID=8829A73D6D0F4DF82E34A5B447352469; PSTM=1567864240; MSA_WH=360_640; H_WISE_SIDS=136262_126008_100805_133333_134982_120141_136986_132910_136456_136620_131247_136722_132378_131517_118885_118875_118857_118823_118798_107313_132783_136800_136431_133351_136862_136818_137012_129652_136195_132250_124634_135308_133847_132552_129645_131423_136220_134318_136165_110085_134254_134153_127969_131754_131953_136614_135416_135458_127416_134935_136636_131545_134353_132467_136413_100457; BD_UPN=12314753; BDUSS=n5iNTVFT3Jrem1uWnRISjZDR2xQMXZZczc0VlpiZ2xCWmRSfjBoWWtkZDA1TkJkSVFBQUFBJCQAAAAAAAAAAAEAAACgp0IQYTUxOTc1NDgyMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHRXqV10V6lda; MCITY=-%3A; BD_HOME=1; H_PS_PSSID=1448_21084_29910_29567_29700_29220_22160; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; BD_CK_SAM=1; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; sug=3; sugstore=1; ORIGIN=0; bdime=0; H_PS_645EC=9da1yRjLjuL8qUx2Hxi4YkTtpfGjYzlNR%2Bf3jmrw3bcKnhEdy%2Bh4b1znz6k",
		},
	)
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

func GetBaiduPcResultsByKeywordAndSearchDay(keyword string, page int, rn int, startDate string) (baiduResults *[]SearchResult, err error) {
	webc, err := GetBaiduPCSearchHtmlWithRNAndTimeDayInterval(keyword, page, rn, startDate)
	if err != nil {
		return
	}
	baiduResults, err = ParseBaiduPCSearchResultHtml(webc)
	return
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
