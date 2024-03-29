package sso

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"
)

var (
	Sdk Instance
)

type Instance struct {
	Host      string
	PublicKey string
	SecretKey string
	Prefix    string
}

func New(publicKey, secretKey string) Instance {
	Sdk = Instance{
		Host:      "https://www.resok.cn",
		PublicKey: publicKey,
		SecretKey: secretKey,
		Prefix:    "/o",
	}
	return Sdk
}

func (c *Instance) SetHost(host string) {
	c.Host = host
}

// 生成基本的public_key url参数
func (c *Instance) getParam() map[string]string {
	return map[string]string{
		"public_key": c.PublicKey,
	}
}

// 生成一个请求req
func (c *Instance) getReq() *req.Request {
	return req.R()
}
func (c *Instance) getJsonReq(body interface{}) *req.Request {
	r := c.getReq()
	r.SetBodyJsonMarshal(body)
	return r
}

// 生成随机字符串
func (c *Instance) randomStr(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// 获取当前时间戳文本
func (c *Instance) getTimeUnixStr() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// Sign 生成一次加密
func (c *Instance) Sign() (string, string, string) {
	rs := c.randomStr(16)
	us := c.getTimeUnixStr()
	return c.sign(rs, us), rs, us
}

// 加密方法
func (c *Instance) sign(randomStr, timeUnix string) string {
	h := md5.New()
	h.Write([]byte(randomStr))
	h.Write([]byte(c.SecretKey))
	h.Write([]byte(timeUnix))
	return hex.EncodeToString(h.Sum(nil))
}

// CheckSign 验证加密
func (c *Instance) CheckSign(sign, randomStr, timeUnix string) bool {
	nowSign := c.sign(randomStr, timeUnix)
	return sign == nowSign
}

// UrlGen 请求url路径生成
func (c *Instance) UrlGen(prefix string, p string) string {
	return c.Host + prefix + p + "?public_key=" + c.PublicKey
}

// RunTr 发起交易 receipt 是否是商品收款
func (c *Instance) RunTr(data ProductReceipt, receipt bool) (ProductPayResp, error, int) {
	data.GenSign()
	var d ProductPayResp
	var msg string
	var url string
	if receipt {
		url = c.UrlGen(c.Prefix, "/receipt")
		msg = "商品收款"
	} else {
		url = c.UrlGen(c.Prefix, "/payment")
		msg = "转账"
	}

	resp, err := c.getJsonReq(data).SetQueryParams(c.getParam()).Post(url)
	if err != nil {
		return d, errors.Wrap(err, "发起交易出错"), 0
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		// 余额不足
		if code == http.StatusUpgradeRequired {
			return d, errors.New("余额不足"), code
		}
		return d, errors.New(fmt.Sprintf("%s响应错误 %d %s", msg, code, resp.String())), code
	}
	err = resp.UnmarshalJson(&d)
	if err != nil {
		return d, errors.Wrap(err, fmt.Sprintf("%s解析返回信息出错", msg)), code
	}
	return d, nil, code
}

// ProductPreOrder 预下单
func (c *Instance) ProductPreOrder(data PreOrder) (PreOrderResp, error) {
	data.GenSign()
	var d PreOrderResp
	url := c.UrlGen(c.Prefix, "/pre_order")
	resp, err := c.getJsonReq(data).Post(url)
	if err != nil {
		return d, errors.Wrap(err, "预下单出错")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("预下单相应失败 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析预下单返回失败")
	}
	return d, nil
}

// UidGetUserInfo 通过uid获取用户信息
func (c *Instance) UidGetUserInfo(uid string) (UidGetUserResp, error) {
	var d UidGetUserResp
	url := c.UrlGen(c.Prefix, "/get_user")
	var p UidGetUserReq
	p.Uid = uid
	p.GenSign()
	resp, err := c.getJsonReq(p).Post(url)
	if err != nil {
		return d, errors.Wrap(err, "获取用户信息请求出错")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("获取用户信息请求错误 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析用户信息出错")
	}
	return d, nil
}

// ChangeUserPower 主动变更用户能力
func (c *Instance) ChangeUserPower(data PowerChangeReq) (bool, error) {
	data.GenSign()
	url := c.UrlGen(c.Prefix, "/power_change")
	resp, err := c.getJsonReq(data).Post(url)
	if err != nil {
		return false, errors.Wrap(err, "变更用户能力出错")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return false, errors.New(fmt.Sprintf("变更用户能力失败 %d %s", code, resp.String()))
	}

	return true, nil

}

// UidGetUserPowerSetting 获取用户能力设置
func (c *Instance) UidGetUserPowerSetting(uid string, eng string) (PowerSettingResp, error) {
	var p PowerSettingReq
	p.Uid = uid
	p.Eng = eng
	p.GenSign()
	url := c.UrlGen(c.Prefix, "/power_setting_new")
	var d PowerSettingResp
	resp, err := c.getJsonReq(p).Post(url)
	if err != nil {
		return d, errors.Wrap(err, "获取能力设置失败")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("获取用户能力设置请求错误 %d %s ", code, resp.String()))
	}
	err = resp.UnmarshalJson(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析用户能力设置出错")
	}
	return d, nil

}

// GetUploadKey 获取上传凭据
func (c *Instance) GetUploadKey() (UploadKeyResp, error) {
	var d UploadKeyResp
	url := c.UrlGen(c.Prefix, "/upload_key")
	resp, err := c.getReq().Get(url)
	if err != nil {
		return d, errors.Wrap(err, "获取上传凭据请求出错")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("获取上传凭据请求出错 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析上传凭据请求出错")
	}
	return d, nil
}

// PreOrderIdGetSuccessList 通过预下单ID获取成交列表
func (c *Instance) PreOrderIdGetSuccessList(preOrderId string, page, pageSize uint64) (*BalanceChangeHistoryResp, error) {
	var r = new(BalanceChangeHistoryResp)
	url := c.UrlGen(c.Prefix, "/pre_order_id")
	params := map[string]interface{}{"pre_order_id": preOrderId, "page": page, "page_size": pageSize}
	sign, st, t := Sdk.Sign()
	params["sign"] = sign
	params["random_str"] = st
	params["t"] = t

	resp, err := c.getReq().SetQueryParamsAnyType(params).Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "获取成交列表失败")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("获取成交列表请求出错 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&r)
	if err != nil {
		return nil, errors.Wrap(err, "解析成交记录失败")
	}
	return r, nil
}

// OrderIdGetInfo 通过orderId获取成交记录
func (c *Instance) OrderIdGetInfo(orderId string) (GetOrderInfoResp, error) {
	var r GetOrderInfoResp
	url := c.UrlGen(c.Prefix, "/order_id")
	params := map[string]string{"order_id": orderId}
	sign, st, t := Sdk.Sign()
	params["sign"] = sign
	params["random_str"] = st
	params["t"] = t
	resp, err := c.getReq().SetQueryParams(params).Get(url)
	if err != nil {
		return r, errors.Wrap(err, "获取成交列表失败")
	}
	code := resp.StatusCode
	if code != http.StatusOK {
		return r, errors.New(fmt.Sprintf("获取成交列表请求出错 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&r)
	if err != nil {
		return r, errors.Wrap(err, "解析成交记录失败")
	}
	return r, nil
}

// UploadImage 上传图片
func (c *Instance) UploadImage(imgPath string, maxWidth int) (UploadImageResp, error) {
	var r UploadImageResp
	url := c.UrlGen(c.Prefix, "/img_upload")
	params := make(map[string]interface{})
	sign, st, t := Sdk.Sign()
	params["sign"] = sign
	params["random_str"] = st
	params["t"] = t
	params["max_width"] = maxWidth
	resp, err := c.getReq().SetQueryParamsAnyType(params).SetFile("file", imgPath).SetResult(&r).Post(url)
	if err != nil {
		return r, errors.Wrap(err, "发起图像上传失败")
	}

	code := resp.StatusCode
	if code != http.StatusOK {
		return r, errors.New(fmt.Sprintf("上传图像失败 %d %s", code, resp.String()))
	}
	err = resp.UnmarshalJson(&r)
	if err != nil {
		return r, errors.Wrap(err, "解析上传图像结果失败")
	}
	return r, nil
}

// HitText 检测文字是否违规
func (c *Instance) HitText(content string) (*HitTextResp, error) {
	url := c.UrlGen("", "/hit_text")
	body := map[string]string{
		"content": content,
	}
	var r = new(HitTextResp)
	resp, err := c.getReq().SetBodyJsonMarshal(body).SetResult(r).Post(url)
	if err != nil && !resp.IsSuccess() {
		return r, errors.Wrap(err, "校验文本失败")
	}
	return r, err
}

// HitImage 检测图像是否违规
func (c *Instance) HitImage(imageUrl string) (*HitImgResp, error) {
	url := c.UrlGen("", "/hit_image")
	body := map[string]string{
		"uri": imageUrl,
	}
	var r = new(HitImgResp)
	resp, err := c.getReq().SetBodyJsonMarshal(body).SetResult(r).Post(url)
	if err != nil && !resp.IsSuccess() {
		return r, errors.Wrap(err, "校验图片合规失败")
	}
	return r, err
}
