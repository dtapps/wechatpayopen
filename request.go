package wechatpayopen

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (gorequest.Response, error) {

	// 认证
	authorization, err := c.authorization(method, params, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 设置头部
	client.SetHeader("Authorization", authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.log.gorm {
		go c.log.logGormClient.GormMiddleware(ctx, request, Version)
	}
	if c.log.mongo {
		go c.log.logMongoClient.MongoMiddleware(ctx, request, Version)
	}

	return request, err
}
