package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	FetchCity()
}

func FetchCity() string {
	res, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {

	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		// 判断数据编码格式
		e := DetermineEncoding(res.Body)
		// 转码（将GBK转码成UTF8）
		utf8Reader := transform.NewReader(res.Body /*simplifiedchinese.GBK.NewDecoder()*/, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {

		}
		fmt.Printf("%s", all)
	}
	return ""
}

// 判断数据编码格式
func DetermineEncoding(reader io.Reader) encoding.Encoding {
	// 读取数据流前1024个字节
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {

	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
