package aliyun_oss

func (f *Client) PutObjectFromFile(remotePath, localPath string) error {
	return f.bucket.PutObjectFromFile(remotePath, localPath)
}

func (f *Client) GetObjectToFile(remotePath, localPath string) error {
	return f.bucket.GetObjectToFile(remotePath, localPath)
}
