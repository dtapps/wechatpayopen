package wechatpayopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ProfitSharingReturnOrdersResponse struct {
	SubMchid    string `json:"sub_mchid"`     // 子商户号
	OrderId     string `json:"order_id"`      // 微信分账单号
	OutOrderNo  string `json:"out_order_no"`  // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
	ReturnId    string `json:"return_id"`     // 微信回退单号
	ReturnMchid string `json:"return_mchid"`  // 回退商户号
	Amount      int    `json:"amount"`        // 回退金额
	Description string `json:"description"`   // 回退描述
	Result      string `json:"result"`        // 回退结果
	FailReason  string `json:"fail_reason"`   // 失败原因
	CreateTime  string `json:"create_time"`   // 创建时间
	FinishTime  string `json:"finish_time"`   // 完成时间
}

type ProfitSharingReturnOrdersResult struct {
	Result ProfitSharingReturnOrdersResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newProfitSharingReturnOrdersResult(result ProfitSharingReturnOrdersResponse, body []byte, http gorequest.Response) *ProfitSharingReturnOrdersResult {
	return &ProfitSharingReturnOrdersResult{Result: result, Body: body, Http: http}
}

// ProfitSharingReturnOrders 请求分账回退API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_3.shtml
func (c *Client) ProfitSharingReturnOrders(ctx context.Context, notMustParams ...gorequest.Params) (*ProfitSharingReturnOrdersResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/return-orders", params, http.MethodPost)
	if err != nil {
		return newProfitSharingReturnOrdersResult(ProfitSharingReturnOrdersResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response ProfitSharingReturnOrdersResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingReturnOrdersResult(response, request.ResponseBody, request), apiError, err
}
