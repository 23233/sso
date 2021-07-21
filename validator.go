package sso_sdk

// ProductPayResp 商品收款返回
type ProductPayResp struct {
	OrderNo string `json:"order_no" form:"order_no"`
	Detail  string `json:"detail" form:"detail"`
}

// PreOrderResp 预下单返回
type PreOrderResp struct {
	PreOrderId string `json:"pre_order_id" form:"pre_order_id"`
}
