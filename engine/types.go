package engine

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

//TODO 	缺少data部分
type Request struct {
	Url         string
	Retry       int               //重试次数
	Timeout     time.Duration     //超时时间
	Interval    time.Duration     //间隔时间
	Method      string            //请求方式
	Header      map[string]string //每个请求自带的header
	VerifyProxy bool              //是否设置代理
	VerifyTLS   bool              //http false  or https true
	ParserFunc  func([]byte) ParseResult
}

//type ParseResult struct {
//	Requests []Request //标记是那个类型的数据(比如列表 详情数据存不同的表)
//	Items    ItemList
//	//Category string
//}

type ParseResult struct {
	Requests []Request //标记是那个类型的数据(比如列表 详情数据存不同的表)
	Items    ItemList
	//Category string
}

type ItemList struct {
	Items    []interface{}
	Category string
}

/**
type ParseResult struct {
	Requests []Request //标记是那个类型的数据(比如列表 详情数据存不同的表)
	Items    []interface{}
	Category string
}
*/
func NilParse([]byte) ParseResult {
	return ParseResult{}
}

//设置一个种子随机获取Ua和代理
func proxy(request *http.Request) (i *url.URL, e error) {
	return url.Parse("")
}

func (r *Request) Fetch() ([]byte, error) {
	time.Sleep(r.Interval)
	tr := &http.Transport{}
	if r.VerifyTLS == true {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	if r.VerifyProxy == true {
		tr.Proxy = proxy
	}

	client := &http.Client{Timeout: r.Timeout, Transport: tr}
	req, err := http.NewRequest(r.Method, r.Url, nil)
	if err != nil {
		log.Printf("http request url error:(%v)", err)
		return nil, err
	}
	if r.Header != nil {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}

	req.Header.Set("User-Agent", UserAgent())
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http response url error:(%v)", err)
		return nil, errors.New("http response error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wrong status code")
	}

	return ioutil.ReadAll(resp.Body)
}
