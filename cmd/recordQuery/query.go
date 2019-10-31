package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var fileNamef = flag.String("f", "domain.csv", "域名文件, 域名按行分割")

func main() {
	flag.Parse()
	fileName := *fileNamef
	domainFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("对应文件未找到")
		panic(err)
	}
	defer domainFile.Close()
	dr := csv.NewReader(domainFile)

	ext := path.Ext(fileName)
	fileNamePrefix := strings.Replace(fileName, ext, "", -1)
	rFileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, time.Now().Format("2006_01_02_15_04_05"))
	rf, err := os.Create(rFileName)
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	rfw := csv.NewWriter(rf)

	for {
		r, err := dr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if len(r) > 0 && len(r[0]) > 3 {
			domain := r[0]
			if strings.Contains(domain, "http") {
				domain = strings.Replace(domain, "http://", "", -1)
				domain = strings.Replace(domain, "https://", "", -1)
			}
			record, err := search.GetPCRecordFromDomain(domain)
			if err != nil {
				fmt.Printf("%s出现错误,错误信息:%s\v", domain, err.Error())
				continue
			}
			fmt.Println(domain, record)
			recordStr := strconv.Itoa(record)
			err = rfw.Write([]string{domain, recordStr})
			if err != nil {
				panic(err)
			}
		}
		rfw.Flush()
	}
}
