package wechatpayopen

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	SpAppid          string             // 服务商应用ID
	SpMchId          string             // 服务商户号
	ApiV2            string             // APIv2密钥
	ApiV3            string             // APIv3密钥
	SerialNo         string             // 序列号
	MchSslSerialNo   string             // pem 证书号
	MchSslCer        string             // pem 内容
	MchSslKey        string             // pem key 内容
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
	CurrentIp        string             // 当前ip
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	currentIp     string         // 当前ip
	config        struct {
		spAppid        string // 服务商应用ID
		spMchId        string // 服务商户号
		subAppid       string // 子商户应用ID
		subMchId       string // 子商户号
		apiV2          string // APIv2密钥
		apiV3          string // APIv3密钥
		serialNo       string // 序列号
		mchSslSerialNo string // pem 证书号
		mchSslCer      string // pem 内容
		mchSslKey      string // pem key 内容
	}
	log struct {
		gorm   bool             // 日志开关
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.currentIp = config.CurrentIp

	c.config.spAppid = config.SpAppid
	c.config.spMchId = config.SpMchId
	c.config.apiV2 = config.ApiV2
	c.config.apiV3 = config.ApiV3
	c.config.serialNo = config.SerialNo
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslCer = config.MchSslCer
	c.config.mchSslKey = config.MchSslKey

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.gorm = true
	}

	return c, nil
}
