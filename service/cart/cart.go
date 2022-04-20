package cart

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/constants"
	"github.com/wtifs/ddmc/service/bark"
	"github.com/wtifs/ddmc/service/log"
)

var cartClient = &http.Client{}

type GetMultipleReserveTimeResp struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    []struct {
		Time []struct {
			DateStr          string `json:"date_str"`
			DateStrTimeStamp string `json:"date_str_time_stamp"`
			Day              string `json:"day"`
			Times            []struct {
				Type      int    `json:"type"`
				FullFlag  bool   `json:"fullFlag"`
				SelectMsg string `json:"select_msg"`
			} `json:"times"`
		} `json:"time"`
	} `json:"data"`
}

// 调用 结算 - 获取配送时间接口
func GetMultiReserveTime(rawBody string) (GetMultipleReserveTimeResp, error) {
	var res GetMultipleReserveTimeResp

	queries, err := url.ParseQuery(rawBody)
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, constants.UrlGetMultiReserveTime, strings.NewReader(rawBody))
	if err != nil {
		return res, err
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "2215")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", config.UserAgent)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Referer", "https://servicewechat.com/wx1e113254eda17715/422/page-frame.html")
	req.Header.Add("Cookie", config.Cookie)

	req.Header.Add("ddmc-device-id", config.DeviceID)
	req.Header.Add("ddmc-city-number", queries.Get("city_number"))
	req.Header.Add("ddmc-build-version", queries.Get("build_version"))
	req.Header.Add("ddmc-station-id", queries.Get("station_id"))
	req.Header.Add("ddmc-channel", queries.Get("channel"))
	req.Header.Add("ddmc-os-version", "[object Undefined]")
	req.Header.Add("ddmc-app-client-id", queries.Get("app_client_id"))
	req.Header.Add("ddmc-ip", "")
	req.Header.Add("ddmc-longitude", queries.Get("longitude"))
	req.Header.Add("ddmc-latitude", queries.Get("latitude"))
	req.Header.Add("ddmc-api-version", queries.Get("api_version"))
	req.Header.Add("ddmc-uid", queries.Get("uid"))
	req.Header.Add("ddmc-time", fmt.Sprintf("%d", time.Now().Unix()))

	log.Debug("%+v", req.Header)

	resp, err := cartClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return res, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return res, err
	}

	log.Debug("resp: %s", body)
	err = json.Unmarshal(body, &res)
	return res, err
}

// 检查配送时间
// 有配送时间表示有运力
func CheckCart(rawBody string) {
	ctx := "check_cart"
	res, err := GetMultiReserveTime(rawBody)
	if err != nil {
		log.Err("%s: failed to get_multi_reserve_time: %s", ctx, err.Error())
		return
	}

	if res.Code != 0 {
		log.Err("%s: resp code 0: %+v", ctx, res)
		return
	}

	if len(res.Data) == 0 || len(res.Data[0].Time) == 0 {
		log.Err("%s: empty list: %+v", ctx, res)
		return
	}

	isFull := true

	time0 := res.Data[0].Time[0]
	for _, t := range time0.Times {
		isFull = isFull && t.FullFlag
		if !isFull {
			bark.Bark("叮咚买菜有货了！最早可预约时间：%s", t.SelectMsg)
			return
		}
	}
	log.Info("无可预约时间")
}
