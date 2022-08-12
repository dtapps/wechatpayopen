package wechatpayopen

import (
	"context"
	"encoding/json"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PayPartnerTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayPartnerTransactionsJsapiResult struct {
	Result   PayPartnerTransactionsJsapiResponse // 结果
	Body     []byte                              // 内容
	Http     gorequest.Response                  // 请求
	Err      error                               // 错误
	ApiError ApiError                            // 接口错误
}

func newPayPartnerTransactionsJsapiResult(result PayPartnerTransactionsJsapiResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsJsapiResult {
	return &PayPartnerTransactionsJsapiResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
func (c *Client) PayPartnerTransactionsJsapi(ctx context.Context, notMustParams ...gorequest.Params) *PayPartnerTransactionsJsapiResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.config.SpAppid)   // 服务商应用ID
	params.Set("sp_mchid", c.config.SpMchId)   // 服务商户号
	params.Set("sub_appid", c.config.SubAppid) // 子商户应用ID
	params.Set("sub_mchid", c.config.SubMchId) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/jsapi", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsJsapiResult(PayPartnerTransactionsJsapiResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 结果
	var response PayPartnerTransactionsJsapiResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsJsapiResult(response, request.ResponseBody, request, err, apiError)
}
