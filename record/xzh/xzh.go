package xzh

import (
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/http-util"
	"strings"
	"time"
)

type PushMessage struct {
	Remain          int64 `json:"remain"`
	Success         int64 `json:"success"`
	SuccessRealtime int64 `json:"success_realtime"`
	RemainRealtime  int64 `json:"remain_realtime"`
}

func DayPush(urls []string, appID string, token string) (pu PushMessage, err error) {
	re, err := http_util.SendRequest(fmt.Sprintf("http://data.zz.baidu.com/urls?appid=%s&token=%s&type=realtime", appID, token), map[string]string{"Content-Type": "text/plain"},
		"POST", []byte(strings.Join(urls, "\n")), 10*time.Second)
	if err != nil {
		return
	}
	re.StatusCode = 200
	rlt, err := http_util.ReadContentFromResponse(re, "UTF-8")
	if err != nil {
		return
	}
	pu = PushMessage{}
	err = json.Unmarshal([]byte(rlt), &pu)
	if err != nil {
		return
	}
	return
}

func WeekPush(urls []string, appID string, token string) (pu PushMessage, err error) {
	re, err := http_util.SendRequest(fmt.Sprintf("http://data.zz.baidu.com/urls?appid=%s&token=%s&type=batch", appID, token), map[string]string{"Content-Type": "text/plain"},
		"POST", []byte(strings.Join(urls, "\n")), 10*time.Second)
	if err != nil {
		return
	}
	rlt, err := http_util.ReadContentFromResponse(re, "UTF-8")
	if err != nil {
		return
	}
	pu = PushMessage{}
	err = json.Unmarshal([]byte(rlt), &pu)
	if err != nil {
		return
	}
	return
}
