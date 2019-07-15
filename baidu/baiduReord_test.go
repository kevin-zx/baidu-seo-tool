package baidu

import (
	"fmt"
	"testing"
)

func TestGetRecordFromDomain(t *testing.T) {
	record, err := GetPCRecordFromDomain("centek.com.cn")
	if err != nil {
		t.Error(err)
	}
	t.Log(record)
}

func TestGetRecordInfo(t *testing.T) {
	testInstances := []string{
		//"www.centek.com.cn",
		//"www.cqjyxzs.com",
		"www.cqzkwx.com"}
	for _, ti := range testInstances {
		rci, err := GetPCRecordInfo(ti)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\n", *rci)
	}
}

func TestGetKeywordSiteRecordInfo(t *testing.T) {
	testInstances := [][]string{
		//{"www.cqjyxzs.com", "家装设计"},
		//{"www.centek.com.cn", "养老"},
		{"www.cqzkwx.com", "重庆自考查询"},
	}

	for _, ti := range testInstances {
		kri, err := GetPCKeywordSiteRecordInfo(ti[1], ti[0])
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\n", kri)
	}
}

func TestGetMobileRecordInfo(t *testing.T) {
	rci, err := GetMobileRecordInfo("scxinhai.top")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rci)
}

func TestGetMobileKeywordSiteRecordInfo(t *testing.T) {
	kri, err := GetMobileKeywordSiteRecordInfo("阀门", "028twt.cn")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", kri)
}
