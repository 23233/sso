package sso_sdk

import (
	"os"
	"testing"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TestNew(t *testing.T) {

	publicKey := GetEnv("public_key", "")
	secretKey := GetEnv("secret_key", "")
	if len(publicKey) < 1 || len(secretKey) < 1 {
		t.Error("未获取到参数")
		return
	}
	s := New(publicKey, secretKey)
	s.SetHost("http://127.0.0.1:7778")
	//
	//t.Run("上传key", func(t *testing.T) {
	//	getUpload(t,s)
	//})
	//
	//// 预下单
	//order, err := s.ProductPreOrder(PreOrder{
	//	Name:      "测试预下单",
	//	Price:     1,
	//	Desc:      "描述",
	//	Extra:     "aaaa",
	//	Substance: "测试sub",
	//	Uid:       "uuuiiiddd",
	//	Count:     10,
	//})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(order.PreOrderId)
	//
	//// 获取用户信息
	//info, err := s.UidGetUserInfo("3ITM5gDN3MDMzA")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(info.Info.NickName)
	//
	//l, err := s.PreOrderIdGetSuccessList(order.PreOrderId, 1, 10)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//println(len(l.Data))
	//
	//getInfo, err := s.OrderIdGetInfo("4213bcd65bf64312a27191f6ca46bacc")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(getInfo.PreOrder.ProductInfo.Desc)

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

func getUpload(t *testing.T, s Sso) {
	resp, err := s.GetUploadKey()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Prefix)
}
