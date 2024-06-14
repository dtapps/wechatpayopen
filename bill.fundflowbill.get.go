package wechatpayopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type BillFundFlowBillGetResponse struct {
	DownloadUrl string `json:"download_url"` // 哈希类型
	HashType    string `json:"hash_type"`    // 哈希值
	HashValue   string `json:"hash_value"`   // 账单下载地址
}

type BillFundFlowBillGetResult struct {
	Result BillFundFlowBillGetResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newBillFundFlowBillGetResult(result BillFundFlowBillGetResponse, body []byte, http gorequest.Response) *BillFundFlowBillGetResult {
	return &BillFundFlowBillGetResult{Result: result, Body: body, Http: http}
}

// BillFundFlowBillGet 申请资金账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_7.shtml
func (c *Client) BillFundFlowBillGet(ctx context.Context, notMustParams ...gorequest.Params) (*BillFundFlowBillGetResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "v3/bill/fundflowbill")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response BillFundFlowBillGetResponse
	var apiError ApiError
	request, err := c.request(ctx, "v3/bill/fundflowbill", params, http.MethodGet, &response, &apiError)
	return newBillFundFlowBillGetResult(response, request.ResponseBody, request), apiError, err
}
