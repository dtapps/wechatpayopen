package wechatpayopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

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

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// NewClient 创建实例化
// spAppid 服务商应用ID
// spMchId 服务商户号
// apiV2 APIv2密钥
// apiV3 APIv3密钥
// serialNo 序列号
// mchSslSerialNo pem 证书号
// mchSslCer pem 内容
// mchSslKey pem key 内容
func NewClient(spAppid, spMchId, apiV2, apiV3, serialNo, mchSslSerialNo, mchSslCer, mchSslKey string, gormClientFun gormClientFun, mongoClientFun mongoClientFun, debug bool) (*Client, error) {

	var err error
	c := &Client{}

	c.config.spAppid = spAppid
	c.config.spMchId = spMchId
	c.config.apiV2 = apiV2
	c.config.apiV3 = apiV3
	c.config.serialNo = serialNo
	c.config.mchSslSerialNo = mchSslSerialNo
	c.config.mchSslCer = mchSslCer
	c.config.mchSslKey = mchSslKey

	c.requestClient = gorequest.NewHttp()

	gormClient := gormClientFun()
	if gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (*dorm.GormClient, string) {
			return gormClient, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
	}
	c.log.gormClient = gormClient

	mongoClient, databaseName := mongoClientFun()
	if mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
	}
	c.log.mongoClient = mongoClient

	return c, nil
}
