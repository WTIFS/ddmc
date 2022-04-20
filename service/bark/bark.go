package bark

import (
	"encoding/json"
	"fmt"
	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/utils/log"
	"io/ioutil"
	"net/http"
)

const (
	barkUrl = "https://api.day.app/" + config.BarkKey + "/叮咚买菜/"
)

var (
	barkClient = &http.Client{}
)

// Bark API 响应结构体
type barkResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// 发送 HTTP GET 请求调用 Bark Api 发通知到 Bark App
func Bark(format string, msgs ...string) int {
	msg := fmt.Sprintf(format, msgs)
	resp, err := barkClient.Get(barkUrl + msg)
	if err != nil {
		log.Err("failed to bark get: %s", err.Error())
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err("failed to bark resp: %s. resp: %s", err.Error(), body)
		return 0
	}

	// 解析响应
	barkResp := &barkResp{}
	if err := json.Unmarshal(body, barkResp); err != nil {
		log.Err("failed to bark unmarshal: %s. resp: %s", err.Error(), body)
		return 0
	}

	return barkResp.Code
}
