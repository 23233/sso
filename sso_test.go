package sso_sdk

import (
	"github.com/pkg/errors"
	"os"
	"testing"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getSdk() (*Sso, error) {
	publicKey := GetEnv("public_key", "")
	secretKey := GetEnv("secret_key", "")
	if len(publicKey) < 1 || len(secretKey) < 1 {
		return nil, errors.New("未找到参数")
	}
	s := New(publicKey, secretKey)
	s.SetHost("http://127.0.0.1:7778")
	return &s, nil
}

func TestNew(t *testing.T) {

	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s.Host)

}

func TestUpload(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := s.GetUploadKey()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Prefix)
}

func TestPreOrder(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}
	// 预下单
	order, err := s.ProductPreOrder(PreOrder{
		Name:      "测试预下单",
		Price:     1,
		Desc:      "描述",
		Extra:     "aaaa",
		Substance: "测试sub",
		Uid:       "uuuiiiddd",
		Count:     10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(order.PreOrderId)
}
func TestGetUserInfo(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}

	// 获取用户信息
	info, err := s.UidGetUserInfo("3ITM5gDN3MDMzA")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info.Info.NickName)

}

func TestPreOrderGetSuccessList(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}
	// 预下单
	order, err := s.ProductPreOrder(PreOrder{
		Name:      "测试预下单",
		Price:     1,
		Desc:      "描述",
		Extra:     "aaaa",
		Substance: "测试sub",
		Uid:       "uuuiiiddd",
		Count:     10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	l, err := s.PreOrderIdGetSuccessList(order.PreOrderId, 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	println(len(l.Data))
}
func TestOrderGetInfo(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}

	getInfo, err := s.OrderIdGetInfo("4213bcd65bf64312a27191f6ca46bacc")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(getInfo.PreOrder.ProductInfo.Desc)

}
func TestImgUpload(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}

	fd, err := os.Open("./t.png")
	if err != nil {
		t.Fatal(err)
		return
	}
	r, err := s.UploadImage(fd, "t.png", 1920)
	if err != nil {
		t.Fatal(err)
		return
	}
	println(r.Origin)
	println(r.Thumbnail)
}
func TestChangeUserPower(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}

	var d PowerChangeReq
	d.Uid = "test_power"
	d.Open = true
	d.Reason = "通过"
	success, err := s.ChangeUserPower(d)
	if err != nil {
		t.Error(err)
	}
	println(success)
}
func TestUidGetUserPowerSetting(t *testing.T) {
	s, err := getSdk()
	if err != nil {
		t.Error(err)
		return
	}
	uid := "3ITM5gDN3MDMzA"
	eng := "test_power"
	resp, err := s.UidGetUserPowerSetting(uid, eng)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.UpdateAt)
	t.Log(resp.Data)

}
func TestAll(t *testing.T) {
	t.Run("获取用户信息", TestGetUserInfo)
	t.Run("获取上传key", TestUpload)
	t.Run("预下单", TestPreOrder)
	t.Run("获取预下单Id获取支付成功列表", TestPreOrderGetSuccessList)
	t.Run("通过orderId获取成交记录", TestOrderGetInfo)
	t.Run("图片上传", TestImgUpload)
	t.Run("变更用户能力", TestChangeUserPower)
	t.Run("获取用户能力设置", TestUidGetUserPowerSetting)
}
