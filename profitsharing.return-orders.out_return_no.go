package wechatpayopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ProfitSharingReturnOrdersOutReturnNoResponse struct {
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

type ProfitSharingReturnOrdersOutReturnNoResult struct {
	Result ProfitSharingReturnOrdersOutReturnNoResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
}

func newProfitSharingReturnOrdersOutReturnNoResult(result ProfitSharingReturnOrdersOutReturnNoResponse, body []byte, http gorequest.Response) *ProfitSharingReturnOrdersOutReturnNoResult {
	return &ProfitSharingReturnOrdersOutReturnNoResult{Result: result, Body: body, Http: http}
}

// ProfitSharingReturnOrdersOutReturnNo 查询分账回退结果API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_4.shtml
func (c *Client) ProfitSharingReturnOrdersOutReturnNo(ctx context.Context, outReturnNo, outOrderNo string) (*ProfitSharingReturnOrdersOutReturnNoResult, ApiError, error) {
	// 参数
	params := gorequest.NewParams()
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("out_return_no", outReturnNo) // 商户回退单号
	params.Set("out_order_no", outOrderNo)   // 商户分账单号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/return-orders/"+outReturnNo, params, http.MethodGet)
	if err != nil {
		return newProfitSharingReturnOrdersOutReturnNoResult(ProfitSharingReturnOrdersOutReturnNoResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response ProfitSharingReturnOrdersOutReturnNoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingReturnOrdersOutReturnNoResult(response, request.ResponseBody, request), apiError, err
}
