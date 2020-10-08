package sso

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"time"
)

type Sso struct {
	Host         string
	AppKey       string
	r            *req.Req
	CdnSecret    string
	CdnExpired   string
	CdnPrefixUrl string
	SecretId     string
	SecretKey    string
}

func New(appKey, secret string) *Sso {
	var s = new(Sso)
	s.Host = "https://Sso.rycsg.com"
	s.AppKey = appKey
	s.r = req.New()
	s.CdnSecret = secret
	s.CdnExpired = "10m"
	s.CdnPrefixUrl = "https://static.rycsg.com"
	s.r.SetTimeout(10 * time.Second)
	return s
}

func (c *Sso) setHost(host string) {
	c.Host = host
}
func (c *Sso) setCdnUrl(host string) {
	c.CdnPrefixUrl = host
}
func (c *Sso) setSecret(secretId, secretKey string) {
	c.SecretId = secretId
	c.SecretKey = secretKey
}

// 校验jwt是否过期
func (c *Sso) CheckJwtExpired(jwt string) (int, error) {
	url := c.Host + "/user/check"
	header := req.Header{
		"Authorization": fmt.Sprintf("bearer %s", jwt),
	}
	resp, err := c.r.Get(url, header)
	if err != nil {
		return 0, errors.Wrap(err, "发送验证jwt是否过期的请求错误")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		if code == iris.StatusUnauthorized {
			return code, errors.New("用户登录信息已过期,请重新登录")
		}
		return code, errors.New(fmt.Sprintf("向sso验证用户jwt发生意外错误 %d %s", code, resp.String()))
	}
	return code, nil
}

// 获取当前用户
func (c *Sso) GetCurrentUser(jwt string) (UserInfo, int, error) {
	var d UserInfo
	url := c.Host + "/user/get_user_info"
	header := req.Header{
		"Authorization": fmt.Sprintf("bearer %s", jwt),
	}
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Get(url, header, param)
	if err != nil {
		return d, 0, errors.Wrap(err, "发送验证jwt是否过期的请求错误")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		if code == iris.StatusUnauthorized {
			return d, code, errors.New("用户登录信息已过期,请重新登录")
		}
		return d, code, errors.New(fmt.Sprintf("向sso获取当前用户失败 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, code, errors.Wrap(err, "解析用户结构出错")
	}
	return d, code, nil

}

// 批量获取用户信息
func (c *Sso) BulkGetUserInfo(body BulkGetUserInfoReq) ([]UserInfo, int, error) {
	d := make([]UserInfo, 0)

	url := c.Host + "/bulk_user_info"
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, param, req.BodyJSON(body))
	if err != nil {
		return d, 0, errors.Wrap(err, "批量获取用户失败请求发送失败")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		if code == iris.StatusUnauthorized {
			return d, code, errors.New("用户登录信息已过期,请重新登录")
		}
		return d, code, errors.New(fmt.Sprintf("批量获取用户失败 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, code, errors.Wrap(err, "批量解析用户数据结构出错")
	}
	return d, code, nil
}

// 发起支付
func (c *Sso) RunPay(jwt string, data ProductPay) (ProductResp, error) {
	var d ProductResp
	url := c.Host + "/user/product_pay"
	header := req.Header{
		"Authorization": fmt.Sprintf("bearer %s", jwt),
	}
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, header, param, req.BodyJSON(data))
	if err != nil {
		return d, errors.Wrap(err, "发起sso支付出错")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return d, errors.New(fmt.Sprintf("发起sso支付出错 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析商品支付数据出错")
	}
	return d, nil
}

// 发起支付退费
func (c *Sso) PayUndo(jwt string, data ProductUndoReq) (int, error) {
	url := c.Host + "/user/pay_undo"
	header := req.Header{
		"Authorization": fmt.Sprintf("bearer %s", jwt),
	}
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, header, param, req.BodyJSON(data))
	if err != nil {
		return 0, errors.Wrap(err, "支付退费请求出错")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return code, errors.New(fmt.Sprintf("支付退费请求出错 %d %s", code, resp.String()))
	}
	return code, nil
}

// 发起用户交易
func (c *Sso) PayReward(jwt string, data ProductRewardReq) (int, error) {
	url := c.Host + "/user/reward"
	header := req.Header{
		"Authorization": fmt.Sprintf("bearer %s", jwt),
	}
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, header, param, req.BodyJSON(data))
	if err != nil {
		return 0, errors.Wrap(err, "支付请求出错")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return code, errors.New(fmt.Sprintf("支付请求出错 %d %s", code, resp.String()))
	}
	return code, nil
}

// 发起app后端退费
func (c *Sso) PayUndoUserUid(data ProductUndoUserReq) (int, error) {
	url := c.Host + "/pay_undo"
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, param, req.BodyJSON(data))
	if err != nil {
		return 0, errors.Wrap(err, "[u]支付退费请求出错")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return code, errors.New(fmt.Sprintf("[u]支付退费请求出错 %d %s", code, resp.String()))
	}
	return code, nil
}

// 通过ticket获取用户
func (c *Sso) TicketGetUser(data ValidTicketReq) (ValidTicketResp, error) {
	var d ValidTicketResp
	url := c.Host + "/valid_ticket"
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Post(url, param, req.BodyJSON(data))
	if err != nil {
		return d, errors.Wrap(err, "发送ticket验证请求出错")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return d, errors.New(fmt.Sprintf("发送ticket验证错误 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析valid出错")
	}
	if len(d.Token) < 1 {
		return d, errors.New("获取token失败")
	}
	return d, nil
}

// 获取所有头像
func (c *Sso) GetAllAvatar() ([]UserAvatar, error) {
	var d []UserAvatar
	url := c.Host + "/get_avatar"
	param := req.Param{
		"app_key": c.AppKey,
	}
	resp, err := c.r.Get(url, param)
	if err != nil {
		return d, errors.Wrap(err, "获取所有头像失败")
	}
	code := resp.Response().StatusCode
	if code != iris.StatusOK {
		return d, errors.New(fmt.Sprintf("获取所有头像失败 %d %s", code, resp.String()))
	}
	err = resp.ToJSON(&d)
	if err != nil {
		return d, errors.Wrap(err, "解析valid出错")
	}
	if len(d) < 1 {
		return d, errors.New("获取所有头像失败")
	}
	return d, nil
}

// 生成资源url访问地址
func (c *Sso) GenCosFileUrl(fileName string) string {
	// 使用typea 办法 详情看 https://cloud.tencent.com/document/product/228/33115#typea
	CdnExpired, _ := time.ParseDuration(c.CdnExpired) // 10分钟链接过期
	nowTime := time.Now().Add(CdnExpired)
	nowTimeStr := fmt.Sprintf("%d", nowTime.Unix())
	randStr := RandString(12)
	uid := fmt.Sprintf("%d", 0)
	md5HashStr := fmt.Sprintf("%s-%s-%s-%s-%s", fileName, nowTimeStr, randStr, uid, c.CdnSecret)
	md5Str := c.GetMD5Encode(md5HashStr)
	url := c.CdnPrefixUrl + fileName + "?sign=" + nowTimeStr + "-" + randStr + "-" + uid + "-" + md5Str
	return url
}

// 获取字符串md5
func (c *Sso) GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
