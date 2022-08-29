package sso

import "time"

// SignBase 签名请求基础
type SignBase struct {
	Sign      string `json:"sign" form:"sign" url:"sign"`
	RandomStr string `json:"random_str" form:"random_str" url:"random_str"`
	T         string `json:"t" form:"t" url:"t"` // unix时间戳字符串
}

func (c *SignBase) GenSign() {
	sign, st, t := Sdk.Sign()
	c.Sign = sign
	c.RandomStr = st
	c.T = t
}

// ProductPayBase 商品支付基础
type ProductPayBase struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`      // 商品uid
	ProductName  string `json:"product_name" form:"product_name" `   // 商品名 必填
	ProductUrl   string `json:"product_url" form:"product_url" `     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price" ` // 商品价格 必填
	Remark       string `json:"remark" form:"remark" `               // 备注
}

// ProductReceipt 商品收款
type ProductReceipt struct {
	Uid string `json:"uid"` // 用户uid
	ProductPayBase
	SignBase
}

// PreOrder 预下单 除了name 和price之外都可以不传
type PreOrder struct {
	Uid        string    `json:"uid"`       // 对应商品UID 可不传
	Count      uint64    `json:"count"`     // 对应商品数量 可不传
	Substance  string    `json:"substance"` // 传什么吐什么
	Name       string    `json:"name"`
	Price      uint64    `json:"price"` // 不允许有免费的出现
	Desc       string    `json:"desc"`
	ImgUrl     string    `json:"img_url"`
	PreviewUrl []string  `json:"preview_url"`
	Extra      string    `json:"extra"` // 传什么吐什么
	ExpireTime time.Time `json:"expire_time"`
	SignBase
}

// BaseUserInfo 基础用户信息
type BaseUserInfo struct {
	NickName     string   `json:"nick_name"`
	AvatarUrl    string   `json:"avatar_url"`
	Powers       []string `json:"powers,omitempty"`
	ManagePowers []string `json:"manage_powers,omitempty"`
}

// UserInfo 用户信息
type UserInfo struct {
	BaseUserInfo
	Id            string `json:"id"`
	UniqueId      string `json:"unique_id,omitempty"`
	TelPhone      string `json:"tel_phone"`       // 手机号
	Balance       uint64 `json:"balance"`         // 余额
	PromoteUserId string `json:"promote_user_id"` // 推荐人
}
