# 叮咚买菜

1. 手机先用叮咚买菜小程序完成用户注册、加购物车流程
2. mac上使用fiddler抓包
3. mac微信小程序搜索叮咚买菜，点击购物车 - 结算按钮，找 `getMultiReserveTime` 这个接口，复制参数到 `config/conf.go` 文件里
4. 运行 `main.go` 文件
