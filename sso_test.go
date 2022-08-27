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
	//s.SetHost("http://127.0.0.1:7778")
	return &s, nil
}

var sdk *Sso

func TestMain(m *testing.M) {
	s, err := getSdk()
	if err != nil {
		println("未找到参数")
		return
	}
	sdk = s
	code := m.Run()
	os.Exit(code)
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

	resp, err := sdk.GetUploadKey()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Prefix)
}

func TestPreOrder(t *testing.T) {
	// 预下单
	order, err := sdk.ProductPreOrder(PreOrder{
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

	// 获取用户信息
	info, err := sdk.UidGetUserInfo("3ITM5gDN3MDMzA")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info.Info.NickName)

}

func TestPreOrderGetSuccessList(t *testing.T) {

	// 预下单
	order, err := sdk.ProductPreOrder(PreOrder{
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
	l, err := sdk.PreOrderIdGetSuccessList(order.PreOrderId, 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	println(len(l.Data))
}
func TestOrderGetInfo(t *testing.T) {
	// 预下单
	order, err := sdk.ProductPreOrder(PreOrder{
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
	getInfo, err := sdk.OrderIdGetInfo(order.PreOrderId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(getInfo.PreOrder.ProductInfo.Desc)

}
func TestImgUpload(t *testing.T) {
	r, err := sdk.UploadImage("./t.png", 1920)
	if err != nil {
		t.Fatal(err)
		return
	}
	println(r.Origin)
	println(r.Thumbnail)
}
func TestChangeUserPower(t *testing.T) {

	var d PowerChangeReq
	d.Uid = "test_power"
	d.Open = true
	d.Reason = "通过"
	success, err := sdk.ChangeUserPower(d)
	if err != nil {
		t.Error(err)
	}
	println(success)
}
func TestUidGetUserPowerSetting(t *testing.T) {
	uid := "3ITM5gDN3MDMzA"
	eng := "test_power"
	resp, err := sdk.UidGetUserPowerSetting(uid, eng)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.UpdateAt)
	t.Log(resp.Data)
}

func TestHitText(t *testing.T) {
	resp, err := sdk.HitText("赵日天我得乖乖得亲亲蛋")
	if err != nil {
		t.Error(err)
	}
	t.Logf("测试结果 %v 文本 %s", resp.Success, resp.Msg)
	resp, _ = sdk.HitText("淫妹阴毛小穴")
	t.Logf("测试结果 %v 文本 %s", resp.Success, resp.Msg)
}

func TestHitImage(t *testing.T) {
	resp, err := sdk.HitImage("https://cdn.golangdocs.com/wp-content/uploads/2020/09/Download-Files-for-Golang.png")
	if err != nil {
		t.Error(err)
	}
	t.Logf("测试结果 %v ", resp.Success)
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
	t.Run("文本内容安全判断", TestHitText)
	t.Run("图像内容安全判断", TestHitImage)
}
