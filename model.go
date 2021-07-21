package sso_sdk

import "time"

// SignBase 签名请求基础
type SignBase struct {
	Sign      string `json:"sign" form:"sign" url:"sign"`
	RandomStr string `json:"random_str" form:"random_str" url:"random_str"`
	T         string `json:"t" form:"t" url:"t"` // unix时间戳字符串
}

// ProductPayBase 商品支付基础
type ProductPayBase struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`      // 商品uid
	ProductName  string `json:"product_name" form:"product_name" `   // 商品名 必填
	ProductUrl   string `json:"product_url" form:"product_url" `     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price" ` // 商品价格 必填
	Remark       string `json:"remark" form:"remark" `               // 备注
	UserUid      string `json:"user_uid" form:"user_uid" `           // 用户标识符
}

// ProductReceipt 商品收款
type ProductReceipt struct {
	Uid string `json:"uid"`
	ProductPayBase
	SignBase
}

// PreOrder 预下单
type PreOrder struct {
	Name       string    `json:"name"`
	Price      uint64    `json:"price"` // 不允许有免费的出现
	Desc       string    `json:"desc"`
	ImgUrl     string    `json:"img_url"`
	PreviewUrl []string  `json:"preview_url"`
	Extra      string    `json:"extra"`
	ExpireTime time.Time `json:"expire_time"`
	SignBase
}

// BaseUserInfo 基础用户信息
type BaseUserInfo struct {
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

// UserInfo 用户信息
type UserInfo struct {
	BaseUserInfo
	Id            string `json:"id"`
	UniqueId      string `json:"unique_id,omitempty"`
	TelPhone      string `json:"tel_phone"`       // 手机号
	Balance       string `json:"balance"`         // 余额
	PromoteUserId string `json:"promote_user_id"` // 推荐人
}

type UidGetUserResp struct {
	User UserInfo     `json:"user"`
	Info BaseUserInfo `json:"info"`
}

// 商品支付
type ProductPay struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`     // 商品uid
	ProductName  string `json:"product_name" form:"product_name"`   // 商品名
	ProductUrl   string `json:"product_url" form:"product_url"`     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price"` // 商品价格
	Remark       string `json:"remark" form:"remark"`
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
