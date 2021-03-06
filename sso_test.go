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
	resp, err := s.GetUploadKey()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Prefix)

}
