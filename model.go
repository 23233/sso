package sso

import "time"

// 用户无敏感信息
type UserInfoBase struct {
	Id        uint64  `json:"id"`
	NickName  string  `json:"nick_name"`
	Uid       string  `json:"uid"`
	Href      string  `json:"href"`
	AvatarUid string  `json:"avatar_uid"`
	Lat       float64 `json:"lat,omitempty"`
	Lng       float64 `json:"lng,omitempty"`
}

// 用户信息
type UserInfo struct {
	UserInfoBase
	UpdateTime      time.Time `json:"update_time"`
	CreateTime      time.Time `json:"create_time"`
	WeOpenId        string    `json:"we_open_id"`
	WeUnionId       string    `json:"we_union_id"`
	TelPhone        uint64    `json:"tel_phone"`
	Credit          uint8     `json:"credit"`
	PromoteBaseLine uint8     `json:"promote_base_line"`
	Balance         uint64    `json:"balance"`
}

// 商品支付
type ProductPay struct {
	ProductUid   string `json:"product_uid" form:"product_uid"`     // 商品uid
	ProductName  string `json:"product_name" form:"product_name"`   // 商品名
	ProductUrl   string `json:"product_url" form:"product_url"`     // 商品url
	ProductPrice uint64 `json:"product_price" form:"product_price"` // 商品价格
}

// 商品支付返回
type ProductResp struct {
	Uid string `json:"uid"`
}

// 商品退费
type ProductUndoReq struct {
	Uid   string `json:"uid"`
	Types int    `json:"types"` // 1是全额 0是部分
	Price uint64 `json:"price"`
	Msg   string `json:"msg"`
}

// 商品退费 用户uid
type ProductUndoUserReq struct {
	UserUid string `json:"user_uid"`
	ProductUndoReq
}

// 用户之间余额交易
type ProductRewardReq struct {
	ToUserUid string `json:"to_user_uid"`
	ProductPay
	Msg string `json:"msg" form:"msg" `
}

// 批量获取用户信息
type BulkGetUserInfoReq struct {
	UidList string `json:"uid_list"`
}

// 验证ticket是否有效
type ValidTicketReq struct {
	Ticket string `json:"ticket"`
}

// 验证结果
type ValidTicketResp struct {
	Token string `json:"token"`
}

// 用户头像
type UserAvatar struct {
	Id         uint64    `json:"id" form:"id"`
	Uid        string    `json:"uid" form:"uid"`
	Name       string    `json:"name" form:"name"`
	Href       string    `json:"href" form:"href"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time"`
}
