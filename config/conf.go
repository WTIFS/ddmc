package config

var (
	BarkKey string // bark通知用的token

	// 固定值
	ApiVersion      = "9.50.0"
	AppBuildVersion = "2.83.0"
	AppClientID     = "4"
	Channel         = "applet"
	ABConfig        = `{"key_onion":"D","key_cart_discount_price":"C"}`

	// 小程序添加商品到购物车 - 结算，抓包 getMultiReserveTime 这个接口，获取以下参数
	Cookie                     string
	DeviceID                   string
	Longitude                  string
	Latitude                   string
	UID                        string
	StationID                  string
	CityID                     string
	GetMultiReserveTimeRawBody string

	// 购物车接口
	SID         string
	DeviceToken string
)
