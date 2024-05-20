package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 阿里云OSS配置
const (
	endpoint        = "your_endpoint" // 例如 "https://oss-cn-hangzhou.aliyuncs.com"
	accessKeyID     = "your_access_key_id"
	accessKeySecret = "your_access_key_secret"
	bucketName      = "your_bucket_name"
)

// 上传单个文件到阿里云OSS
func uploadFile(client *oss.Client, bucketName, localFilePath, remoteFilePath string) error {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("获取bucket失败: %v", err)
	}

	err = bucket.PutObjectFromFile(remoteFilePath, localFilePath)
	if err != nil {
		return fmt.Errorf("上传文件失败: %v", err)
	}

	fmt.Printf("上传成功: %s -> %s\n", localFilePath, remoteFilePath)
	return nil
}

// 递归遍历文件夹并上传图片
func uploadDirectory(client *oss.Client, bucketName, localDir, remoteDir string) error {
	err := filepath.Walk(localDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".png") || strings.HasSuffix(info.Name(), ".jpg")) {
			remotePath := filepath.Join(remoteDir, strings.TrimPrefix(path, localDir))
			remotePath = strings.ReplaceAll(remotePath, "\\", "/")
			return uploadFile(client, bucketName, path, remotePath)
		}
		return nil
	})
	return err
}

func main() {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Fatalf("创建OSS客户端失败: %v", err)
	}

	localDir := "./"                     // 本地图片文件夹的根路径
	remoteDir := "your_remote_directory" // OSS上对应的远程目录

	err = uploadDirectory(client, bucketName, localDir, remoteDir)
	if err != nil {
		log.Fatalf("上传文件夹失败: %v", err)
	}
}
