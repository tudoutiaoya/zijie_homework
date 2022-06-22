package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictResponse struct {
	ErrorCode       int `json:"errorCode"`
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	Type        string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
}

func queryYoudao(word string) {
	body := "i=" + word + "&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=16559056823361&sign=bb7071e062de6da3848d43f83694f8a2&lts=1655905682336&bv=bdc0570a34c12469d01bfac66273680d&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME"
	client := &http.Client{}
	var data = strings.NewReader(body)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=-1137252044@10.110.96.154; OUTFOX_SEARCH_USER_ID_NCOO=597363497.2865735; UM_distinctid=180abe8c56f903-03168fc4a5ffaf-17333273-11442c-180abe8c5703d8; fanyi-ad-id=306808; fanyi-ad-closed=1; ___rl__test__cookies=1655905682331")
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dicResponse DictResponse
	err = json.Unmarshal(bodyText, &dicResponse)
	if err != nil {
		log.Fatal(err)
	}
	result := dicResponse.SmartResult.Entries

	for _, v := range result {
		fmt.Println(v)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("参数个数不对")
		os.Exit(1)
	}
	wold := os.Args[1]
	fmt.Println("有道翻译")
	queryYoudao(wold)
	fmt.Println("===================")
}
