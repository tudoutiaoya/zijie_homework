package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

//有道
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
	defer wg.Done()
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
	fmt.Println("有道翻译")
	for _, v := range result {
		fmt.Println(v)
	}
	fmt.Println("======================")
}

//彩云
type DicRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DicResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   []interface{} `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func queryCaiyun(word string) {
	defer wg.Done()
	client := &http.Client{}
	dicRequest := DicRequest{TransType: "en2zh", Source: word}
	body, _ := json.Marshal(dicRequest)
	var data = bytes.NewReader(body)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
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
	var dicResponse DicResponse
	err = json.Unmarshal(bodyText, &dicResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("彩云翻译")
	for _, v := range dicResponse.Dictionary.Explanations {
		fmt.Println(v)
	}
	fmt.Println("====================")
}

var wg sync.WaitGroup

func main() {
	if len(os.Args) != 2 {
		fmt.Println("参数个数不对")
		os.Exit(1)
	}
	wold := os.Args[1]
	wg.Add(1)
	go queryYoudao(wold)
	wg.Add(1)
	go queryCaiyun(wold)

	wg.Wait()
}
