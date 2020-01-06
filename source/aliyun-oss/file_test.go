package aliyun_oss_test

import "testing"

func TestClient_PutObjectFromFile(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	if err := client.PutObjectFromFile("file.go", "file.go"); err != nil {
		t.Error(err)
		return
	}
}

func TestClient_GetObjectToFile(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	if err := client.GetObjectToFile("my-profile", "p.jpg"); err != nil {
		t.Error(err)
		return
	}
}
