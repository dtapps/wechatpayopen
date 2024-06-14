package wechatpayopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
	"time"
)

type CertificatesResponse struct {
	Data []struct {
		EffectiveTime      time.Time `json:"effective_time"` // 过期时间
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
			Nonce          string `json:"nonce"`
		} `json:"encrypt_certificate"` // 加密证书
		ExpireTime time.Time `json:"expire_time"` // 有效时间
		SerialNo   string    `json:"serial_no"`   // 序列号
	} `json:"data"`
}

type CertificatesResult struct {
	Result CertificatesResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newCertificatesResult(result CertificatesResponse, body []byte, http gorequest.Response) *CertificatesResult {
	return &CertificatesResult{Result: result, Body: body, Http: http}
}

// Certificates 获取平台证书列表
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func (c *Client) Certificates(ctx context.Context) (*CertificatesResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "pay")
	defer c.TraceEndSpan()

	// 参数
	request, err := c.request(ctx, "v3/certificates", nil, http.MethodGet, &response, nil)
	if err != nil {
		return newCertificatesResult(CertificatesResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CertificatesResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCertificatesResult(response, request.ResponseBody, request), err
}
