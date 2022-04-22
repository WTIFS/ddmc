package cart

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/wtifs/ddmc/utils"
	"github.com/wtifs/ddmc/utils/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/constants"
	"github.com/wtifs/ddmc/service/bark"
)

type CartProduct struct {
	ProductName string `json:"product_name"`
}

// 购物车响应结构体
type CartIndexResp struct {
	CommonResp
	Data struct {
		Product struct {
			// 有库存的
			Effective []struct {
				Products []CartProduct `json:"products"`
			} `json:"effective"`

			// 无库存的
			Invalid []struct {
				Products []CartProduct `json:"products"`
			} `json:"invalid"`
		} `json:"product"`
	} `json:"data"`
}

// js签名结果结构体
type JsSign struct {
	Nars string `json:"nars"`
	Sesi string `json:"sesi"`
}

// 调用购物车库存接口
func GetCart() (CartIndexResp, error) {
	var res CartIndexResp

	// 构造请求参数
	queryMap := map[string]string{
		"uid":           config.UID,
		"longitude":     config.Longitude,
		"latitude":      config.Latitude,
		"station_id":    config.StationID,
		"city_number":   config.CityID,
		"api_version":   config.ApiVersion,
		"app_version":   config.AppBuildVersion,
		"channel":       config.Channel,
		"app_client_id": config.AppClientID,
		"s_id":          config.SID,
		"openid":        config.DeviceID,
		"time":          fmt.Sprintf("%d", time.Now().Unix()),
		"device_token":  config.DeviceToken,
		"is_load":       "1",
		"ab_config":     config.ABConfig,
	}

	// 构造请求
	requestURL := constants.URLGetCart + "?" + utils.Map2URLValues(queryMap).Encode()
	log.Debug("GET %s", requestURL)
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)
	setCommonRequestHeader(req)

	// 调用购物车 API，发起请求
	resp, err := cartClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("HTTP GET: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return res, fmt.Errorf("gzip decode: %s", err.Error())
	}
	body, err := ioutil.ReadAll(reader)
	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read: %s", err.Error())
	}

	log.Debug("resp: %s", body)
	err = json.Unmarshal(body, &res)
	return res, err
}

// 加密函数，经测试不加密也可以调通接口
func addSign(req *http.Request, queryMap map[string]string) error {
	// 读取js文件
	filePath := utils.GetCurrentAbPath() + "/sign.js"
	jsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("load sign.js: %s", err.Error())
	}

	log.Debug("%s", jsBytes)

	vm := otto.New()
	_, err = vm.Run(string(jsBytes))
	if err != nil {
		return fmt.Errorf("run otto: %s", err.Error())
	}

	// 调用js文件里的签名函数
	signParams, _ := json.Marshal(queryMap)
	signValue, err := vm.Call("sign", nil, string(signParams))
	if err != nil {
		return fmt.Errorf("otto call: %s", err.Error())
	}
	log.Debug("sign: %s", signValue)

	// 解析签名结果
	jsSign := JsSign{}
	if err := json.Unmarshal([]byte(signValue.String()), &jsSign); err != nil {
		return fmt.Errorf("parse sign: %s", err.Error())
	}

	// 将签名参数加到请求里
	queryMap["nars"] = jsSign.Nars
	queryMap["sesi"] = jsSign.Sesi

	return nil
}

// 检查购物车库存
func CheckCart() bool {
	ctx := "check_cart"
	res, err := GetCart()
	if err != nil {
		log.Err("%s: failed to get_cart: %s", ctx, err.Error())
		return false
	}

	if res.Code != 0 {
		log.Err("%s: resp code 0: %+v", ctx, res)
		return false
	}

	if len(res.Data.Product.Effective) == 0 || len(res.Data.Product.Effective[0].Products) == 0 {
		log.Info("购物车所有商品均无库存")
		return false
	}

	availableProducts := res.Data.Product.Effective[0].Products

	availableProductNames := make([]string, 0, len(availableProducts))
	for _, p := range availableProducts {
		availableProductNames = append(availableProductNames, p.ProductName)
	}
	bark.Bark("以下商品有库存了：%s", strings.Join(availableProductNames, ", "))
	return true
}
