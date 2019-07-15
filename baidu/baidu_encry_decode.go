package baidu

import (
	"github.com/kevin-zx/http-util"
	"strings"
	"time"
)

func DecodeBaiduEncURL(baiduUrl string) string {

	response, err := http_util.SendRequest(baiduUrl, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"}, "GET", nil, 10*time.Second)
	if err != nil {
		return ""
	}
	//fmt.Println(baiduUrl)
	if response.StatusCode == 200 {

		h, _ := http_util.ReadContentFromResponse(response, "")
		if strings.Contains(h, "window.opener&&window.opener.bds&&window.opener.bds.pdc&&window.opener.bds.pdc.sendLinkLog") {

			part1 := strings.Split(h, "window.location.replace(\"")
			if len(part1) < 2 {
				//fmt.Println(baiduUrl)
				return baiduUrl
			} else {
				return strings.Split(part1[1], "\")},timeout")[0]
			}
		} else {
			return response.Request.URL.String()
		}
	}

	return baiduUrl

}
