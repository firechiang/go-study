package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	getMethodTest1()
	getMethodTest2()
}

// 简单的get请求
func getMethodTest1() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}

func getMethodTest2() {
	// 创建请求体
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// 配置头信息
	request.Header.Add("token", "0xzxdsdsx1vsddas342dfgdgd55dfgdf5")
	// 发起请求
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}

func getMethodTest3() {
	// 创建请求体
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// 配置头信息
	request.Header.Add("token", "0xzxdsdsx1vsddas342dfgdgd55dfgdf5")
	// 自定义Http Client
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("服务器说要重定向")
		return nil
	}}
	// 发起请求
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}
