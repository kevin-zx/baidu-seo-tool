package search

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func GetPCRecordFromDomain(domain string) (int, error) {
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+domain, 1, 20)
	if err != nil {
		return 0, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return 0, err
	}
	recordContainer := doc.Find("div.op_site_domain.c-row div span b")

	if recordContainer != nil && recordContainer.Size() > 0 {
		recordStr := strings.Replace(recordContainer.Text(), ",", "", -1)
		record, err := strconv.Atoi(recordStr)
		return record, err
	}
	siteTipsRecord := doc.Find("div.c-border.c-row.site_tip b")
	if siteTipsRecord != nil && siteTipsRecord.Size() > 0 {
		recordStr := strings.Replace(siteTipsRecord.Text(), "找到相关结果数约", "", 1)[0:1]
		recordStr = strings.Replace(recordStr, ",", "", -1)
		record, err := strconv.Atoi(recordStr)
		return record, err
	}

	return 0, nil
}

type KeywordRecordInfo struct {
	RecordInfo
	Keyword string
}

func GetPCKeywordSiteRecordInfo(keyword string, domain string) (kri *KeywordRecordInfo, err error) {
	kri = &KeywordRecordInfo{Keyword: keyword}
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+strings.Replace(domain, "www.", "", 1)+" "+keyword, 1, 20)
	if err != nil {
		return
	}
	srs, err := ParseBaiduPCSearchResultHtml(pageData)
	if err != nil {
		return
	}
	kri.SearchResults = srs
	kri.HomePageRank = GetFirstHomePageRank(srs, domain)
	for _, sr := range *srs {
		if sr.SiteName != "" {
			kri.SiteName = sr.SiteName
			break
		}
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return
	}
	size := doc.Find("#page > a").Size()
	if size == 0 {
		kri.Record = len(*srs)
	} else {
		kri.Record = size * 10
	}

	return
}

func GetMobileKeywordSiteRecordInfo(keyword string, domain string) (kri *KeywordRecordInfo, err error) {
	kri = &KeywordRecordInfo{Keyword: keyword}
	rs, err := GetBaiduMobileResultsByKeyword("site:"+strings.Replace(domain, "www.", "", 1)+" "+keyword, 1)

	if rs == nil || len(*rs) == 0 {
		return
	}
	kri.HomePageRank = GetFirstHomePageRank(rs, domain)
	kri.SearchResults = rs
	for _, r := range *rs {
		if r.SiteName != "" && r.SiteName != "总收录量：" {
			kri.SiteName = r.SiteName
			break
		}
	}

	if len(*rs) >= 1 && kri.Record == 0 {
		if (*rs)[len(*rs)-1].Rank != 10 {
			kri.Record = len(*rs)
		} else {
			var sRs *[]SearchResult
			sRs, err = GetBaiduMobileResultsByKeyword("site:"+domain, 2)
			if err != nil {
				return
			}
			if len(*sRs) < 10 {
				kri.Record = 10 + len(*sRs)
			} else {
				kri.Record = 21
			}
		}
	}

	return
}

type RecordInfo struct {
	Record        int
	HomePageRank  int
	SearchResults *[]SearchResult
	SiteName      string
}

// 根据网站 域名 名称（熊掌号名称） 主页title 获取收录信息  siteName homePageTitle 为空则可以不填
func GetPCRecordInfo(domain string) (rci *RecordInfo, err error) {
	rci = &RecordInfo{}
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+strings.Replace(domain, "www.", "", 1), 1, 20)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return
	}

	// 获取首页位置
	srs, err := ParseBaiduPCSearchResultHtml(pageData)
	if err != nil {
		return
	}
	rci.HomePageRank = GetFirstHomePageRank(srs, domain)
	rci.SearchResults = srs

	// 获取siteName
	for _, sr := range *srs {
		if sr.SiteName != "" {
			rci.SiteName = sr.SiteName
			break
		}
	}

	// 百度正常显示收录的方式
	recordContainer := doc.Find("div.op_site_domain.c-row div span b")
	if recordContainer != nil && recordContainer.Size() > 0 {
		rci.Record = recordStr2Record(recordContainer.Text())
		if rci.HomePageRank > 0 {
			rci.HomePageRank--
		}
		if err != nil {
			return
		}
	}

	// 百度简略显示收录的方式
	siteTipsRecord := doc.Find("div.c-border.c-row.site_tip b")
	if siteTipsRecord != nil && siteTipsRecord.Size() > 0 {
		recordStr := strings.Replace(siteTipsRecord.Text(), "找到相关结果数约", "", 1)[0:]
		recordStr = strings.Replace(recordStr, ",", "", -1)
		recordStr = strings.Replace(recordStr, "个", "", -1)
		rci.Record, err = strconv.Atoi(recordStr)
		if err != nil {
			return
		}
	}
	return
}

func GetMobileRecordInfo(domain string) (rci *RecordInfo, err error) {
	rci = &RecordInfo{}
	webCon, err := GetBaiduMobileSearchHtml("site:"+domain, 1)
	if err != nil {
		return
	}
	rs, err := ParseBaiduMobileSearchResultHtml(webCon, 1)
	if err != nil {
		return
	}
	if len(*rs) == 0 {
		return
	}
	if len(*rs) >= 0 && (*rs)[0].Type == "2H5" {
		var doc *goquery.Document
		doc, err = goquery.NewDocumentFromReader(strings.NewReader(webCon))
		if err != nil {
			return
		}
		recordEle := doc.Find("#results>div.c-result span.c-color-orange").First()
		if recordEle != nil {
			rci.Record = recordStr2Record(recordEle.Text())
		}
	}

	if len(*rs) >= 1 && rci.Record == 0 {
		if (*rs)[len(*rs)-1].Rank != 10 {
			rci.Record = len(*rs)
		} else {
			var sRs *[]SearchResult
			sRs, err = GetBaiduMobileResultsByKeyword("site:"+domain, 2)
			if err != nil {
				return
			}
			if len(*sRs) < 10 {
				rci.Record = 10 + len(*sRs)
			} else {
				rci.Record = 21
			}
		}
	}

	rci.HomePageRank = GetFirstHomePageRank(rs, domain)
	rci.SearchResults = rs
	for _, r := range *rs {
		if r.SiteName != "" && r.SiteName != "总收录量：" {
			rci.SiteName = r.SiteName
			break
		}
	}

	return
}

func recordStr2Record(recordStr string) (record int) {
	recordStr = strings.Replace(recordStr, "亿", "", -1)
	recordStr = strings.Replace(recordStr, "个", "", -1)
	recordStr = strings.Replace(recordStr, ",", "", -1)
	//是否有万
	wMatchFlag := strings.Index(recordStr, "万") >= 0
	recordStr = strings.Replace(recordStr, "万", "", -1)
	record, err := strconv.Atoi(recordStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	if wMatchFlag {
		return record * 10000
	}
	return
}
