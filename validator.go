package sso_sdk

import "time"

type RespBase struct {
	Id       string    `json:"id,omitempty" url:"id" form:"id"`
	UpdateAt time.Time `json:"update_at" url:"update_at" form:"update_at"`
	CreateAt time.Time `json:"create_at" url:"create_at" form:"create_at"`
}

// ProductPayResp 商品收款返回
type ProductPayResp struct {
	OrderNo string `json:"order_no" form:"order_no"`
	Detail  string `json:"detail" form:"detail"`
}

// PreOrderResp 预下单返回
type PreOrderResp struct {
	PreOrderId string `json:"pre_order_id" form:"pre_order_id"`
}

type BalanceChangeHistoryResp struct {
	Page     uint64                     `json:"page"`
	PageSize uint64                     `json:"page_size"`
	Data     []BalanceChangeHistoryItem `json:"data"`
	Total    uint64                     `json:"total"` //
}

type PreOrderItem struct {
	RespBase
	AppId       string      `json:"app_id"`
	ProductInfo ProductInfo `json:"product_info"`
	Substance   string      `json:"substance"`
	Extra       string      `json:"extra"`
	ExpireTime  time.Time   `json:"expire_time"`
}

type GetOrderInfoResp struct {
	Pay      BalanceChangeHistoryItem `json:"pay" form:"pay"`
	PreOrder PreOrderItem             `json:"pre_order"`
}

// BalanceChangeHistoryItem 成交记录
type BalanceChangeHistoryItem struct {
	RespBase
	UserId      string `json:"user_id,omitempty" url:"user_id" form:"user_id"`
	PreOrderId  string `json:"pre_order_id,omitempty" url:"pre_order_id" form:"pre_order_id"`
	MapId       string `json:"map_id,omitempty" url:"map_id" form:"map_id"`
	Quantity    uint64 `json:"quantity,omitempty" url:"quantity" form:"quantity"` // 支付金额
	ProductUid  string `json:"product_uid,omitempty" url:"product_uid" form:"product_uid"`
	ProductName string `json:"product_name,omitempty" url:"product_name" form:"product_name"`
	ProductUrl  string `json:"product_url,omitempty" url:"product_url" form:"product_url"`
	Remark      string `json:"remark,omitempty" url:"remark" form:"remark"`
	OrderUid    string `json:"order_uid,omitempty" url:"order_uid" form:"order_uid"`
	PublicKey   string `json:"public_key,omitempty" url:"public_key" form:"public_key"`
	Extra       string `json:"extra" url:"extra" from:"extra"`
}

type UidGetUserReq struct {
	Uid string `json:"uid" form:"uid"`
	SignBase
}

type UidGetUserResp struct {
	User UserInfo     `json:"user"`
	Info BaseUserInfo `json:"info"`
}

func (c *UidGetUserResp) HasPower(name string) bool {
	for _, power := range c.Info.Powers {
		if name == power {
			return true
		}
	}
	return false
}
func (c *UidGetUserResp) HasManagePower(name string) bool {
	for _, power := range c.Info.ManagePowers {
		if name == power {
			return true
		}
	}
	return false
}

type PowerChangeReq struct {
	UidGetUserReq
	Eng    string `json:"eng" form:"eng"`
	Open   bool   `json:"open" form:"open" `    // 打开还是关闭
	Reason string `json:"reason" form:"reason"` // 理由
}

type PowerSettingReq struct {
	UidGetUserReq
	Eng string `json:"eng" form:"eng"`
}

type PowerSettingResp struct {
	UpdateAt string `json:"update_at"`
	Data     string `json:"data"`
}

// UploadKeyResp 上传key请求resp
type UploadKeyResp struct {
	SecretID     string
	SecretKey    string
	SessionToken string
	ExpiredTime  uint64
	Prefix       string
	Visit        string
}

type UploadImageResp struct {
	Origin    string `json:"origin"`
	Thumbnail string `json:"thumbnail"`
}

// JsonSchemaReq 用户填写了表单的回调
type JsonSchemaReq struct {
	SendUserInfo struct {
		Mid       string `json:"mid" form:"mid" comment:"用户id" validate:"required"`
		PublicKey string `json:"public_key,omitempty" form:"public_key,omitempty"`
	} `json:"send_user_info" form:"send_user_info" validate:"required"`
	SendSignInfo SignBase `json:"send_sign_info,omitempty" form:"send_sign_info,omitempty"`
	FormId       string   `json:"form_id" form:"form_id" comment:"表单ID" validate:"required"`
	FormEng      string   `json:"form_eng,omitempty" form:"form_eng,omitempty" comment:"表单英文唯一"`
	Data         string   `json:"data" form:"data" comment:"表单数据" validate:"required"`
	InjectData   string   `json:"inject_data,omitempty" form:"inject_data,omitempty" comment:"注入数据"`
}

type HitTextResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type HitImgResp struct {
	Success bool `json:"success"`
}
