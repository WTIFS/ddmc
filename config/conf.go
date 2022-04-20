package config

const (
	BarkKey = "zARZ7G8bQP9nRnxXHH3cj3" // bark通知用的token

	// 固定值
	ApiVersion      = "9.50.0"
	AppBuildVersion = "2.83.0"
	AppClientID     = "4"
	Channel         = "applet"

	// 小程序添加商品到购物车 - 结算，抓包 getMultiReserveTime 这个接口，获取以下参数
	Cookie                     = "DDXQSESSID=1e7af1fcb3f754c50c511658038810b1"
	DeviceID                   = "osP8I0cYRAQOeNFNxoheO_yPyV2Q"
	Longitude                  = "121.5325"
	Latitude                   = "31.25229"
	UID                        = "625d0d66289afd0001562b55"
	StationID                  = "5bf2907a716de100468c6ca3"
	CityID                     = "0101"
	GetMultiReserveTimeRawBody = "uid=625d0d66289afd0001562b55&longitude=121.5325&latitude=31.25229&station_id=5bf2907a716de100468c6ca3&city_number=0101&api_version=9.50.0&app_version=2.83.0&applet_source=&channel=applet&app_client_id=4&sharer_uid=&s_id=1e7af1fcb3f754c50c511658038810b1&openid=osP8I0cYRAQOeNFNxoheO_yPyV2Q&h5_source=&time=1650444965&device_token=WHJMrwNw1k%2FFKPjcOOgRd%2BNCP82pZEcA9roozj%2BhXwXH7LnR1QsTHBHg02DD97l0tQKN7SmGKu%2BSdxTzn6QTJ1dlTA7iL2F1ydCW1tldyDzmauSxIJm5Txg%3D%3D1487582755342&address_id=625fba7711d35b0001970860&group_config_id=&products=%5B%5B%7B%22type%22%3A1%2C%22id%22%3A%225a2e53096ec2c403528b45aa%22%2C%22price%22%3A%2229.90%22%2C%22count%22%3A1%2C%22description%22%3A%22%22%2C%22sizes%22%3A%5B%5D%2C%22cart_id%22%3A%225a2e53096ec2c403528b45aa%22%2C%22parent_id%22%3A%22%22%2C%22parent_batch_type%22%3A-1%2C%22category_path%22%3A%2258f9e514936edfe3568b572e%2C5b0ffd8e06752e9e288bf1f2%22%2C%22manage_category_path%22%3A%22738%2C1058%2C1922%22%2C%22activity_id%22%3A%22%22%2C%22sku_activity_id%22%3A%22%22%2C%22conditions_num%22%3A%22%22%2C%22product_name%22%3A%22%E9%87%91%E5%AD%97%E9%87%91%E5%8D%8E%E9%A6%99%E8%82%A0%20208g%22%2C%22product_type%22%3A0%2C%22small_image%22%3A%22https%3A%2F%2Fimg.ddimg.mobi%2FformData%2F04d8d7c8ddaec1588817677259.jpg!deliver.product.list%22%2C%22total_price%22%3A%2229.90%22%2C%22origin_price%22%3A%2229.90%22%2C%22total_origin_price%22%3A%2229.90%22%2C%22no_supplementary_price%22%3A%2229.90%22%2C%22no_supplementary_total_price%22%3A%2229.90%22%2C%22size_price%22%3A%220.00%22%2C%22buy_limit%22%3A0%2C%22price_type%22%3A0%2C%22promotion_num%22%3A0%2C%22instant_rebate_money%22%3A%220.00%22%2C%22is_invoice%22%3A1%2C%22sub_list%22%3A%5B%5D%2C%22is_booking%22%3A0%2C%22is_bulk%22%3A0%2C%22view_total_weight%22%3A%22%E8%A2%8B%22%2C%22net_weight%22%3A%22208%22%2C%22net_weight_unit%22%3A%22g%22%2C%22storage_value_id%22%3A0%2C%22temperature_layer%22%3A%22%22%2C%22sale_batches%22%3A%7B%22batch_type%22%3A-1%7D%2C%22is_shared_station_product%22%3A0%2C%22is_gift%22%3A0%2C%22supplementary_list%22%3A%5B%5D%2C%22order_sort%22%3A4%2C%22is_presale%22%3A0%7D%5D%5D&isBridge=false&nars=541bcf0c1720d5486709f96ac46fe741&sesi=LQrgRgL6ffc9994155e7f577874f5d7ba313963"

	// 购物车接口
	SID         = "1e7af1fcb3f754c50c511658038810b1"
	DeviceToken = "WHJMrwNw1k/FKPjcOOgRd+NCP82pZEcA9roozj+hXwXH7LnR1QsTHBHg02DD97l0tQKN7SmGKu+SdxTzn6QTJ1dlTA7iL2F1ydCW1tldyDzmauSxIJm5Txg==1487582755342"
	ABConfig    = `{"key_onion":"D","key_cart_discount_price":"C"}`
)
