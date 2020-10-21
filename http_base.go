package XMPushSDK

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func PostResult(requestPath string, fields map[string]interface{}, retries int, appSecret string) (*Result, error) {
	result, err := postReq(requestPath, fields, appSecret)
	if result.Code == Success && err == nil {
		return result, nil
	}
	// 重试
	for i := 0; i < retries; i++ {
		result, err = postReq(requestPath, fields, appSecret)
		if result.Code == Success && err == nil {
			break
		}
	}
	return result, err
}

func postReq(requestPath string, fields map[string]interface{}, appSecret string) (*Result, error) {
	form := url.Values{}
	for k, v := range fields {
		form.Add(k, v.(string))
	}
	param := form.Encode()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s", baseHost()+requestPath), strings.NewReader(param))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", appSecret))

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
	}

	var result = &Result{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func baseHost() string {
	return ProductionHost
}
