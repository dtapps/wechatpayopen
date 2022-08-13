package wechatpayopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	SpAppid        string           // 服务商应用ID
	SpMchId        string           // 服务商户号
	SubAppid       string           // 子商户应用ID
	SubMchId       string           // 子商户号
	ApiV2          string           // APIv2密钥
	ApiV3          string           // APIv3密钥
	SerialNo       string           // 序列号
	MchSslSerialNo string           // pem 证书号
	MchSslCer      string           // pem 内容
	MchSslKey      string           // pem key 内容
	GormClient     *dorm.GormClient // 日志数据库
	LogClient      *golog.GoLog     // 日志驱动
	LogDebug       bool             // 日志开关
}

// Client 微信支付服务
type Client struct {
	requestClient *gorequest.App   // 请求服务
	logClient     *golog.ApiClient // 日志服务
	config        *ConfigClient    // 配置
}

// NewClient 实例化
func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
