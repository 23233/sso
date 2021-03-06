package sso_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"
)

var (
	Sdk Sso
)

type Sso struct {
	Host      string
	PublicKey string
	SecretKey string
	Prefix    string
}

func New(publicKey, secretKey string) Sso {
	Sdk = Sso{
		Host:      "https://www.resok.cn",
		PublicKey: publicKey,
		SecretKey: secretKey,
		Prefix:    "/o",
	}
	return Sdk
}

func (c *Sso) SetHost(host string) {
	c.Host = host
}

// 生成基本的public_key url参数
func (c *Sso) getParam() req.Param {
	return req.Param{
		"public_key": c.PublicKey,
	}
}

// 生成一个请求req
func (c *Sso) getReq() *req.Req {
	r := req.New()
	r.SetTimeout(10 * time.Second)
	return r
}

// 生成随机字符串
func (c *Sso) randomStr(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// 获取当前时间戳文本
func (c *Sso) getTimeUnixStr() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// Sign 生成一次加密
func (c *Sso) Sign() (string, string, string) {
	rs := c.randomStr(16)
	us := c.getTimeUnixStr()
	return c.s(rs, us), rs, us
}

// 加密方法
func (c *Sso) s(randomStr, timeUnix string) string {
	h := md5.New()
	h.Write([]byte(randomStr))
	h.Write([]byte(c.SecretKey))
	h.Write([]byte(timeUnix))
	return hex.EncodeToString(h.Sum(nil))
}

// CheckSign 验证加密
func (c *Sso) CheckSign(sign, randomStr, timeUnix string) bool {
	nowSign := c.s(randomStr, timeUnix)
	return sign == nowSign
}

// UrlGen 请求url路径生成
func (c *Sso) UrlGen(prefix string, p string) string {
	return c.Host + prefix + p + "?public_key=" + c.PublicKey
}

// RunTr 发起交易 receipt 是否是商品收款
func (c *Sso) RunTr(data ProductReceipt, receipt bool) (ProductPayResp, error, int) {
	var d ProductPayResp
	var msg string
	var url string
	if receipt {
		url = c.UrlGen(c.Prefix, "receipt")
		msg = "商品收款"
	} else {
		url = c.UrlGen(c.Prefix, "payment")
		msg = "转账"
	}

	resp, err := c.getReq().Post(url, c.getParam(), req.BodyJSON(data))
	if err != nil {
		return d, errors.Wrap(err, "发起交易出错"), 0
	}
	code := resp.Response().StatusCode
	if code != http.StatusOK {
		// 余额不足
		if code == http.StatusUpgradeRequired {
			return d, errors.New("余额不足"), code
		}
		return d, errors.New(fmt.Sprintf("%s响应错误 %d %s", msg, code, resp.String())), code
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, fmt.Sprintf("%s解析返回信息出错", msg)), code
	}
	return d, nil, code
}

// ProductPreOrder 预下单
func (c *Sso) ProductPreOrder(data PreOrder) (PreOrderResp, error) {
	var d PreOrderResp
	url := c.UrlGen(c.Prefix, "/pre_order")
	resp, err := c.getReq().Post(url, c.getParam(), req.BodyJSON(data))
	if err != nil {
		return d, errors.Wrap(err, "预下单出错")
	}
	code := resp.Response().StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("预下单相应失败 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析预下单返回失败")
	}
	return d, nil
}

// UidGetUserInfo 通过uid获取用户信息
func (c *Sso) UidGetUserInfo(uid string) (UidGetUserResp, error) {
	var d UidGetUserResp
	url := c.UrlGen(c.Prefix, "/get_user")
	body := req.BodyJSON(map[string]interface{}{"uid": uid})
	resp, err := c.getReq().Post(url, c.getParam(), body)
	if err != nil {
		return d, errors.Wrap(err, "获取用户信息请求出错")
	}
	code := resp.Response().StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("获取用户信息请求错误 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析用户信息出错")
	}
	return d, nil
}

// GetUploadKey 获取上传凭据
func (c *Sso) GetUploadKey() (UploadKeyResp, error) {
	var d UploadKeyResp
	url := c.UrlGen(c.Prefix, "/upload_key")
	resp, err := c.getReq().Get(url)
	if err != nil {
		return d, errors.Wrap(err, "获取上传凭据请求出错")
	}
	code := resp.Response().StatusCode
	if code != http.StatusOK {
		return d, errors.New(fmt.Sprintf("获取上传凭据请求出错 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析上传凭据请求出错")
	}
	return d, nil
}
