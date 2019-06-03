package fetcher

import (
	"crawer/engine"
	"errors"
	"io/ioutil"
	"net/http"
)

/**
这里的添加代理等信息
*/
func Fetch(r engine.Request) ([]byte, error) {
	resp, err := http.Get(r.Url)
	//resp, err := http.Get("https://mp.weixin.qq.com/s/Ee_A0cmciXVo0gDSbSX8kw")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wrong status code")
	}

	//reader :=   transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//ParseContents(resp.Body)
	return ioutil.ReadAll(resp.Body)

}
