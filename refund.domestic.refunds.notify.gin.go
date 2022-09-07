package wechatpayopen

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ResponseRefundDomesticRefundsNoNotifyGin struct {
	Id           string `form:"id" json:"status" xml:"id" uri:"id" binding:"required"`                                         // 通知ID
	CreateTime   string `form:"create_time" json:"create_time" xml:"create_time" uri:"create_time" binding:"required"`         // 通知创建时间
	EventType    string `form:"event_type" json:"event_type" xml:"event_type" uri:"event_type" binding:"required"`             // 通知类型
	Summary      string `form:"summary" json:"summary" xml:"summary" uri:"summary" binding:"required"`                         // 通知简要说明
	ResourceType string `form:"resource_type" json:"resource_type" xml:"resource_type" uri:"resource_type" binding:"required"` // 通知数据类型
	Resource     struct {
		Algorithm      string `form:"algorithm" json:"algorithm" xml:"algorithm" uri:"algorithm" binding:"required"`                          // 加密算法类型
		Ciphertext     string `form:"ciphertext" json:"ciphertext" xml:"ciphertext" uri:"ciphertext" binding:"required"`                      // 数据密文
		AssociatedData string `form:"associated_data" json:"associated_data" xml:"associated_data" uri:"associated_data" binding:"omitempty"` // 附加数据
		OriginalType   string `form:"original_type" json:"original_type" xml:"original_type" uri:"original_type" binding:"required"`          // 原始类型
		Nonce          string `form:"nonce" json:"nonce" xml:"nonce" uri:"nonce" binding:"required"`                                          // 随机串
	} `form:"resource" json:"resource" xml:"resource" uri:"resource" binding:"required"` // 通知数据
}

// RefundDomesticRefundsNoNotifyGin 申请退款API - 回调通知
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_11.shtml
func (c *Client) RefundDomesticRefundsNoNotifyGin(ctx context.Context, ginCtx *gin.Context) (ResponseRefundDomesticRefundsNoNotifyGin, error) {

	// 声明接收的变量
	var validateJson struct {
		Id           string `form:"id" json:"status" xml:"id" uri:"id" binding:"required"`                                         // 通知ID
		CreateTime   string `form:"create_time" json:"create_time" xml:"create_time" uri:"create_time" binding:"required"`         // 通知创建时间
		EventType    string `form:"event_type" json:"event_type" xml:"event_type" uri:"event_type" binding:"required"`             // 通知类型
		Summary      string `form:"summary" json:"summary" xml:"summary" uri:"summary" binding:"required"`                         // 通知简要说明
		ResourceType string `form:"resource_type" json:"resource_type" xml:"resource_type" uri:"resource_type" binding:"required"` // 通知数据类型
		Resource     struct {
			Algorithm      string `form:"algorithm" json:"algorithm" xml:"algorithm" uri:"algorithm" binding:"required"`                          // 加密算法类型
			Ciphertext     string `form:"ciphertext" json:"ciphertext" xml:"ciphertext" uri:"ciphertext" binding:"required"`                      // 数据密文
			AssociatedData string `form:"associated_data" json:"associated_data" xml:"associated_data" uri:"associated_data" binding:"omitempty"` // 附加数据
			OriginalType   string `form:"original_type" json:"original_type" xml:"original_type" uri:"original_type" binding:"required"`          // 原始类型
			Nonce          string `form:"nonce" json:"nonce" xml:"nonce" uri:"nonce" binding:"required"`                                          // 随机串
		} `form:"resource" json:"resource" xml:"resource" uri:"resource" binding:"required"` // 通知数据
	}

	err := ginCtx.ShouldBind(&validateJson)

	return validateJson, err
}
