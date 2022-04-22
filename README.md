# 叮咚买菜
个人学习用

## 功能
1. 检查购物车里的商品是否有库存，在有库存时通过 Bark App 发送通知到手机（仅支持iOS）
2. 检查有库存的商品是否有运力配送

## 食用方法
1. 手机先用叮咚买菜小程序完成用户注册、加购物车流程
2. mac 上使用 fiddler 抓包（fiddler 自带 HTTPS 证书安装指引，可以直接解密 HTTPS 请求，比较方便）
3. mac 微信小程序搜索叮咚买菜
4. 点击购物车按钮，找 `cart/index` 和 `getMultiReserveTime` 这个接口，复制参数到 `config/conf.go` 文件里
5. 运行 `main.go` 文件

## 参考项目
- python: https://github.com/jozhn/ddmc.monitor
- JAVA: https://github.com/JannsenYang/dingdong-helper