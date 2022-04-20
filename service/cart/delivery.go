package cart

import (
	"encoding/json"
	"fmt"
	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/constants"
	"github.com/wtifs/ddmc/service/bark"
	"github.com/wtifs/ddmc/utils/log"
	"io/ioutil"
	"net/http"
	"strings"
)

// 检查配送时间

type GetMultipleReserveTimeResp struct {
	CommonResp
	Data []struct {
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
func GetMultiReserveTime() (GetMultipleReserveTimeResp, error) {
	ctx := "get_multi_reserve_time"
	var res GetMultipleReserveTimeResp

	//postBody := make(url.Values)
	//postBody.Set("uid", config.UID)
	//postBody.Set("longitude", config.Longitude)
	//postBody.Set("latitude", config.Latitude)
	//postBody.Set("station_id", config.StationID)
	//postBody.Set("city_number", config.CityID)
	//postBody.Set("api_version", config.ApiVersion)
	//postBody.Set("app_version", config.AppBuildVersion)
	//postBody.Set("channel", config.Channel)
	//postBody.Set("app_client_id", config.AppClientID)
	//postBody.Set("sid", config.SID)
	//postBody.Set("openid", config.DeviceID)
	//postBody.Set("time", fmt.Sprintf("%d", time.Now().Unix()))
	//postBody.Set("device_token", config.DeviceToken)
	//req, err := http.NewRequest(http.MethodPost, constants.URLGetMultiReserveTime, strings.NewReader(postBody.Encode()))

	req, err := http.NewRequest(http.MethodPost, constants.URLGetMultiReserveTime, strings.NewReader(config.GetMultiReserveTimeRawBody))
	if err != nil {
		return res, fmt.Errorf("make new request err: %s", ctx, err.Error())
	}

	setCommonRequestHeader(req)

	resp, err := cartClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("post: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	//reader, err := gzip.NewReader(resp.Body)
	//if err != nil {
	//	return res, fmt.Errorf("gzip decode: %s", err.Error())
	//}
	//body, err := ioutil.ReadAll(reader)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read: %s", err.Error())
	}

	log.Debug("resp: %s", body)
	err = json.Unmarshal(body, &res)
	return res, err
}

// 检查配送时间
// 有配送时间表示有运力
func CheckDeliverTime() {
	ctx := "check_cart"
	res, err := GetMultiReserveTime()
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
			bark.Bark("叮咚买菜购物车里的商品可以配送了！最早可预约时间：%s", t.SelectMsg)
			return
		}
	}
	log.Info("无可预约时间")
}
