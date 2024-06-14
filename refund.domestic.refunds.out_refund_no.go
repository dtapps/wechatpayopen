package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
	"time"
)

type RefundDomesticRefundsOutRefundNoResponse struct {
	RefundId            string    `json:"refund_id"`
	OutRefundNo         string    `json:"out_refund_no"`
	TransactionId       string    `json:"transaction_id"`
	OutTradeNo          string    `json:"out_trade_no"`
	Channel             string    `json:"channel"`
	UserReceivedAccount string    `json:"user_received_account"`
	SuccessTime         string    `json:"success_time"`
	CreateTime          time.Time `json:"create_time"`
	Status              string    `json:"status"`
	FundsAccount        string    `json:"funds_account"`
	Amount              struct {
		Total  int `json:"total"`
		Refund int `json:"refund"`
		From   []struct {
			Account string `json:"account"`
			Amount  int    `json:"amount"`
		} `json:"from"`
		PayerTotal       int    `json:"payer_total"`
		PayerRefund      int    `json:"payer_refund"`
		SettlementRefund int    `json:"settlement_refund"`
		SettlementTotal  int    `json:"settlement_total"`
		DiscountRefund   int    `json:"discount_refund"`
		Currency         string `json:"currency"`
	} `json:"amount"`
	PromotionDetail []struct {
		PromotionId  string `json:"promotion_id"`
		Scope        string `json:"scope"`
		Type         string `json:"type"`
		Amount       int    `json:"amount"`
		RefundAmount int    `json:"refund_amount"`
		GoodsDetail  struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`
			WechatpayGoodsId string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			UnitPrice        int    `json:"unit_price"`
			RefundAmount     int    `json:"refund_amount"`
			RefundQuantity   int    `json:"refund_quantity"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
}

type RefundDomesticRefundsOutRefundNoResult struct {
	Result RefundDomesticRefundsOutRefundNoResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newRefundDomesticRefundsOutRefundNoResult(result RefundDomesticRefundsOutRefundNoResponse, body []byte, http gorequest.Response) *RefundDomesticRefundsOutRefundNoResult {
	return &RefundDomesticRefundsOutRefundNoResult{Result: result, Body: body, Http: http}
}

// RefundDomesticRefundsOutRefundNo 查询单笔退款API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_9.shtml
func (c *Client) RefundDomesticRefundsOutRefundNo(ctx context.Context, outRefundNo string) (*RefundDomesticRefundsOutRefundNoResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/refund/domestic/refunds/%s?sub_mchid=%s", outRefundNo, c.GetSubMchId()))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParams()

	// 请求
	var response RefundDomesticRefundsOutRefundNoResponse
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/refund/domestic/refunds/%s?sub_mchid=%s", outRefundNo, c.GetSubMchId()), params, http.MethodGet, &response, &apiError)
	return newRefundDomesticRefundsOutRefundNoResult(response, request.ResponseBody, request), apiError, err
}
