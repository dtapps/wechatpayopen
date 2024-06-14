package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult struct {
	Body []byte             // 内容
	Http gorequest.Response // 请求
}

func newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(body []byte, http gorequest.Response) *PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult {
	return &PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult{Body: body, Http: http}
}

// PayPartnerTransactionsOutTradeNoOutTradeNoClosePost 关闭订单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_3.shtml
func (c *Client) PayPartnerTransactionsOutTradeNoOutTradeNoClosePost(ctx context.Context, outTradeNo string, notMustParams ...gorequest.Params) (*PayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/pay/partner/transactions/out-trade-no/%s/close", outTradeNo), params, http.MethodPost, nil, &apiError)
	return newPayPartnerTransactionsOutTradeNoOutTradeNoClosePostResult(request.ResponseBody, request), apiError, err
}
