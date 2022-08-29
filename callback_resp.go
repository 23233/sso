package sso

// 各类事件回调传入的参数

// TrSendBodyResp 收银台收款
type TrSendBodyResp struct {
	Event     string `json:"event,omitempty"` // cash_receive
	Name      string `json:"name,omitempty"`
	OrderNo   string `json:"order_no,omitempty"`
	ToUser    string `json:"to_user,omitempty"`
	PayUser   string `json:"pay_user,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Price     uint64 `json:"price,omitempty"`
	DetailUrl string `json:"detail_url,omitempty"`
	SignBase
}

// ChangeInfoSendBodyResp 变更app里的用户基本信息
type ChangeInfoSendBodyResp struct {
	Event     string `json:"event,omitempty"` //app_change_info
	AvatarUrl string `json:"avatar_url,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	ToUser    string `json:"to_user,omitempty"`
	SignBase
}

// ChangePhoneSendBodyResp 变更用户手机号
type ChangePhoneSendBodyResp struct {
	Event     string `json:"event,omitempty"` // user_change_phone
	PublicKey string `json:"public_key,omitempty"`
	ToUser    string `json:"to_user,omitempty"`
	Phone     string `json:"phone,omitempty"`
	SignBase
}

type ProductInfo struct {
	Uid        string   `json:"uid" comment:"唯一ID"`
	Count      uint64   `json:"count" comment:"对应数量"`
	Name       string   `json:"name"  comment:"商品名"`
	Price      uint64   `json:"price"  comment:"价格"`
	Desc       string   `json:"desc" comment:"描述"`
	ImgUrl     string   `json:"img_url"  comment:"主图"`
	PreviewUrl []string `json:"preview_url"  comment:"预览图"` // 其他图片
}

// ProductPaySendBodyResp 预下单商品支付完成
type ProductPaySendBodyResp struct {
	Event      string      `json:"event,omitempty"` // product_pay
	ToUser     string      `json:"to_user,omitempty"`
	Extra      string      `json:"extra,omitempty"`
	Substance  string      `json:"substance,omitempty"`
	Remark     string      `json:"remark,omitempty"`
	OrderNo    string      `json:"order_no,omitempty"`
	PreOrderId string      `json:"pre_order_id,omitempty"`
	Product    ProductInfo `json:"product"`
	SignBase
}

type PowerRespBase struct {
	Eng string `json:"eng"`
	Uid string `json:"uid"`
}

// PowerChangeBodyResp 用户能力变更通知
type PowerChangeBodyResp struct {
	SignBase
	PowerRespBase
	Open   bool   `json:"open"`
	Reason string `json:"reason,omitempty"`
}

// PowerNeedVerifyBodyResp 用户能力提交申请等待审核通知
type PowerNeedVerifyBodyResp struct {
	SignBase
	PowerRespBase
}

// PowerSettingChangeBodyResp 用户能力设置变更通知
type PowerSettingChangeBodyResp struct {
	SignBase
	PowerRespBase
	Body string `json:"body"`
}
