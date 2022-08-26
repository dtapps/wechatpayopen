package wechatpayopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// ClientConfig 实例配置
type ClientConfig struct {
	SpAppid        string         // 服务商应用ID
	SpMchId        string         // 服务商户号
	ApiV2          string         // APIv2密钥
	ApiV3          string         // APIv3密钥
	SerialNo       string         // 序列号
	MchSslSerialNo string         // pem 证书号
	MchSslCer      string         // pem 内容
	MchSslKey      string         // pem key 内容
	GormClientFun  gormClientFun  // 日志配置
	MongoClientFun mongoClientFun // 日志配置
	Debug          bool           // 日志开关
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
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
		gormClient     *dorm.GormClient  // 日志数据库
		gorm           bool              // 日志开关
		logGormClient  *golog.ApiClient  // 日志服务
		mongoClient    *dorm.MongoClient // 日志数据库
		mongo          bool              // 日志开关
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.config.spAppid = config.SpAppid
	c.config.spMchId = config.SpMchId
	c.config.apiV2 = config.ApiV2
	c.config.apiV3 = config.ApiV3
	c.config.serialNo = config.SerialNo
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslCer = config.MchSslCer
	c.config.mchSslKey = config.MchSslKey

	c.requestClient = gorequest.NewHttp()

	gormClient := config.GormClientFun()
	if gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (*dorm.GormClient, string) {
			return gormClient, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
	}
	c.log.gormClient = gormClient

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
	}
	c.log.mongoClient = mongoClient

	return c, nil
}
