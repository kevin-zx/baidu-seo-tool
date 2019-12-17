package main

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"sort"
	"strings"
)

func main() {
	keyword := "python"
	brs, err := search.GetBaiduPcResultsByKeyword(keyword, 1, 50)
	if err != nil {
		panic(err)
	}
	km := GetKeywords(brs, keyword)
	for k, v := range km {
		fmt.Println(k, v)
	}
	//fmt.Println(rks)
}

func GetKeywords(brs *[]search.SearchResult, rootKey string) map[string]int {
	var rks []string
	for _, br := range *brs {
		rks = append(rks, br.TitleMatchWords...)
		rks = append(rks, br.BaiduDescriptionMatchWords...)
	}
	rks = RemoveDuplicatesAndEmpty(rks)
	var wordCount = make(map[string]int)
	wordCount[rootKey] = 1
	sort.Slice(rks, func(i, j int) bool {
		if len(rks[i]) < len(rks[j]) {
			return false
		} else {
			return true
		}
	})

	for _, br := range *brs {
		for _, rk := range rks {
			GetSingleRkKeys(rk, br, wordCount, brs)
		}
	}
	//for k, v := range wordCount {
	//	fmt.Println(k,v)
	//}
	return wordCount

}

func GetSingleRkKeys(rk string, br search.SearchResult, km map[string]int, brs *[]search.SearchResult) {
	title := br.Title + br.BaiduDescription
	title = clearHoleText(title)
	tCount := 1
	for {
		nrk := rk
		if !strings.Contains(title, rk) {
			break
		}
		forwardExist := true
		tnrk := nrk
		for forwardExist {
			tnrk, forwardExist = GetNextKeyPart(nrk, title, true)
			if !forwardExist {
				break
			}
			c := CountWord(tnrk, brs)
			//if !ExistInMap(km, tnrk) && c >= tCount && c>2 {
			if !ExistInMap(km, tnrk) && c > 2 {
				tCount = c
				nrk = tnrk
			} else {
				break
			}
		}
		backWardExist := true
		for backWardExist {
			tnrk, backWardExist = GetNextKeyPart(nrk, title, false)
			if !backWardExist {
				break
			}
			c := CountWord(tnrk, brs)
			//if !ExistInMap(km, tnrk) && c >= tCount  && c>2{
			if !ExistInMap(km, tnrk) && c > 2 {
				tCount = c
				nrk = tnrk
			} else {
				break
			}
		}

		if len(nrk) > len(rk) && tCount > 2 {
			km[nrk] = tCount
			title = strings.Replace(title, nrk, "", 1)
		} else {
			title = strings.Replace(title, nrk, "", 1)
		}

	}

}

func ExistInMap(km map[string]int, key string) bool {
	for k := range km {
		if strings.Contains(k, key) {
			return true
		}
	}
	return false
}

var punctuations = []string{"（", "）", "(", ")", "·", "=", "-", "，", "。", "、", "；", "’", "【", "】", "、", "`", "！", "@", "#", "￥", "%", "…", "…", "&", "×", "—", "—", "《", "》", "？", "：", "”", "“", "{", "}", "‘", "|", "～", "+", ",", ".", "/", ";", "'", "[", "]", "\\", "`", "!", "@", "#", "$", "%", "^", "&", "*", "_", "+", "<", ">", "?", ":", "\"", "{", "}", "~"}
var stopWords = []string{
	"在",
	"于",
	"的",
	"了",
	"和",
	"是",
	"就", "都", "而", "及", "与", "着", "或",
}
var emptybrackets = []string{"((", "（（", "()", "（）", "||", " ", " "}

func clearHoleText(wholeText string) string {
	//for _, l := range letters {
	//	wholeText = strings.Replace(wholeText,l,"",-1)
	//}
	//for _, n := range nums {
	//	wholeText = strings.Replace(wholeText,n,"",-1)
	//}
	for _, p := range punctuations {
		wholeText = strings.Replace(wholeText, p, "|", -1)

	}
	for _, p := range stopWords {
		wholeText = strings.Replace(wholeText, p, "|", -1)

	}
	//wholeText = strings.Replace(wholeText,"名称",punctuations[rand.Intn(10)]+"名称"+punctuations[rand.Intn(10)],-1)
	//wholeText = strings.Replace(wholeText,"地址",punctuations[rand.Intn(10)]+"地址"+punctuations[rand.Intn(10)],-1)
	for _, eb := range emptybrackets {
		for strings.Contains(wholeText, eb) {
			wholeText = strings.Replace(wholeText, eb, "", -1)
		}
	}

	return wholeText
}
func GetNextKeyPart(key string, parentString string, forward bool) (nextKey string, exist bool) {
	keyParts := strings.Split(key, "")
	parentPs := strings.Split(parentString, "")
	pl := len(parentPs)
	kl := len(keyParts)
	if forward {
		for i, _ := range parentPs {
			if i+kl+1 > pl {
				return key, false
			}
			//if keyParts[0] == s {
			//
			//}
			s := ""
			match := true
			for ki, kp := range keyParts {

				s = parentPs[i+ki+1]
				if kp != parentPs[i+ki] {
					match = false
					break
				}
			}
			if match && s != "|" {
				return key + s, true
			}
		}
	} else {
		for i, _ := range parentPs {
			if i >= pl {
				return key, false
			}
			//if keyParts[0] == s {
			//
			//}
			s := ""
			match := true
			if i-kl < 1 {
				continue
			}
			s = parentPs[i-kl-1]
			for ki := kl - 1; ki >= 0; ki-- {
				kp := keyParts[ki]
				//fmt.Println(kp)
				//fmt.Println(parentPs[i-(kl-ki)])
				if kp != parentPs[i-(kl-ki)] {
					match = false
					break
				}
			}

			if match && s != "|" {
				return s + key, true
			}
		}

	}
	return
}

func CountWord(key string, brs *[]search.SearchResult) int {
	var count = 0
	for _, br := range *brs {
		if strings.Contains(br.Title, key) || strings.Contains(br.BaiduDescription, key) {
			count++
		}
	}
	return count
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	var keywordCount = make(map[string]int)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		duFlag := false
		for _, re := range ret {
			if len(a[i]) == 0 {
				duFlag = true
				break
			}
			if re == a[i] {
				if _, ok := keywordCount[re]; !ok {
					keywordCount[re] = 1
				}
				duFlag = true
				break
			}
		}
		if !duFlag {
			ret = append(ret, a[i])
		}
	}
	return
}
