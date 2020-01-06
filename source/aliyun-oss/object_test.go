package aliyun_oss_test

import (
	"testing"
)

func TestClient_GetObject(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	fileBytes, err := client.GetObject("test/test.json")
	if err != nil {
		t.Errorf("get file bytes error : %v", err)
		return
	}
	t.Log(string(fileBytes))
}

func TestClient_PutObject(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	if err := client.PutObject("test/test.txt", "wsqunasfdsd"); err != nil {
		t.Error(err)
		return
	}
}
