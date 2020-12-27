package sso

// 商品收款返回
type ProductPayResp struct {
	Order  string `json:"order" form:"order"`
	Detail string `json:"detail" form:"detail"`
}
