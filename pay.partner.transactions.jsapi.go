package wechatpayopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PayPartnerTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayPartnerTransactionsJsapiResult struct {
	Result PayPartnerTransactionsJsapiResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
}

func newPayPartnerTransactionsJsapiResult(result PayPartnerTransactionsJsapiResponse, body []byte, http gorequest.Response) *PayPartnerTransactionsJsapiResult {
	return &PayPartnerTransactionsJsapiResult{Result: result, Body: body, Http: http}
}

// PayPartnerTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
func (c *Client) PayPartnerTransactionsJsapi(ctx context.Context, notMustParams ...gorequest.Params) (*PayPartnerTransactionsJsapiResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "v3/pay/partner/transactions/jsapi")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var response PayPartnerTransactionsJsapiResponse
	var apiError ApiError
	request, err := c.request(ctx, "v3/pay/partner/transactions/jsapi", params, http.MethodPost, &response, &apiError)
	return newPayPartnerTransactionsJsapiResult(response, request.ResponseBody, request), apiError, err
}
