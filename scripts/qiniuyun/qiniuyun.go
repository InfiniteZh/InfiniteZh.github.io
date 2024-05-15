/*
将图片上传到七牛云的oss对象存储服务
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func main() {
	// 配置七牛云的 AccessKey 和 SecretKey
	accessKey := "CQwDB_qxHOG4oRWM98WtnhINjsInqUDo4qrPdsLB"
	secretKey := "3rqpIliuvCUNnBXoqIpLi1kQp8dlvXF09890_bfn"
	bucket := "infinite-blog"
	localDir := "C:\\Users\\11930\\Desktop\\博客\\JWT"

	// 获取最后一个文件夹的名称
	_, lastDir := filepath.Split(localDir)

	// 初始化七牛云配置
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 根据你的存储区域选择Zone
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	uploader := storage.NewFormUploader(&cfg)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	// 遍历本地目录并上传文件
	err := filepath.Walk(localDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// 计算相对路径
		relPath, err := filepath.Rel(localDir, path)
		if err != nil {
			return err
		}
		//////////////////////////////////////////////////
		/* 针对全部数据的一次性上传 */
		// key := strings.ReplaceAll(relPath, "\\", "/")
		/* 针对单个文件夹的上传 */
		key := filepath.Join(lastDir, relPath)
		key = strings.ReplaceAll(key, "\\", "/")
		///////////////////////////////////////////////////
		fmt.Printf("Uploading file: %s as key: %s\n", path, key) // 调试输出
		err = uploadFile(uploader, upToken, key, path)
		if err != nil {
			log.Printf("Failed to upload %s: %v", path, err)
		} else {
			log.Printf("Successfully uploaded %s to %s", path, key)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to traverse directory: %v", err)
	}
}

func uploadFile(uploader *storage.FormUploader, upToken, key, filePath string) error {
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := uploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return err
	}
	fmt.Printf("File uploaded: %s => %s\n", filePath, ret.Key)
	return nil
}
