package cart

import (
	"fmt"
	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/utils/log"
	"net/http"
	"time"
)

var cartClient = &http.Client{}

type CommonResp struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

// 设置请求头
func setCommonRequestHeader(req *http.Request) {
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E217 MicroMessenger/6.8.0(0x16080000) NetType/WIFI Language/en Branch/Br_trunk MiniProgramEnv/Mac")
	req.Header.Add("Accept-Encoding", "gzip,compress,br,deflate")
	req.Header.Add("Referer", "https://servicewechat.com/wx1e113254eda17715/422/page-frame.html")
	req.Header.Add("Cookie", config.Cookie)

	req.Header.Add("ddmc-device-id", config.DeviceID)
	req.Header.Add("ddmc-city-number", config.CityID)
	req.Header.Add("ddmc-build-version", config.AppBuildVersion)
	req.Header.Add("ddmc-station-id", config.StationID)
	req.Header.Add("ddmc-channel", config.Channel)
	req.Header.Add("ddmc-os-version", "[object Undefined]")
	req.Header.Add("ddmc-app-client-id", config.AppClientID)
	req.Header.Add("ddmc-ip", "")
	req.Header.Add("ddmc-longitude", config.Longitude)
	req.Header.Add("ddmc-latitude", config.Latitude)
	req.Header.Add("ddmc-api-version", config.ApiVersion)
	req.Header.Add("ddmc-uid", config.UID)
	req.Header.Add("ddmc-time", fmt.Sprintf("%d", time.Now().Unix()))

	log.Debug("%+v", req.Header)
}
