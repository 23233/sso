package sso_sdk

// 签名请求基础
type SignBase struct {
	Sign      string `json:"sign" form:"sign" url:"sign"`
	RandomStr string `json:"random_str" form:"random_str" url:"random_str"`
	T         string `json:"t" form:"t" url:"t"` // unix时间戳字符串
}

// 商品支付基础
type ProductPayBase struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`      // 商品uid
	ProductName  string `json:"product_name" form:"product_name" `   // 商品名 必填
	ProductUrl   string `json:"product_url" form:"product_url" `     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price" ` // 商品价格 必填
	Remark       string `json:"remark" form:"remark" `               // 备注
	UserUid      string `json:"user_uid" form:"user_uid" `           // 用户标识符
}

// 商品收款
type ProductReceipt struct {
	ProductPayBase
	SignBase
}

// ticket获取用户
type TicketGetUserReq struct {
	Ticket string `json:"ticket"`
	SignBase
}

// 用户信息
type UserInfo struct {
	Uid       string `json:"uid"`
	UniqueId  string `json:"unique_id,omitempty"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

// 商品支付
type ProductPay struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`     // 商品uid
	ProductName  string `json:"product_name" form:"product_name"`   // 商品名
	ProductUrl   string `json:"product_url" form:"product_url"`     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price"` // 商品价格
	Remark       string `json:"remark" form:"remark"`
}

// 验证ticket是否有效
type ValidTicketReq struct {
	Ticket string `json:"ticket"`
}

// 上传key请求resp
type UploadKeyResp struct {
	SecretID     string
	SecretKey    string
	SessionToken string
	ExpiredTime  uint64
	Prefix       string
	Visit        string
}
