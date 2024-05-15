/*
获取七牛云的所有图片的链接
*/
package main

import (
	"fmt"
	"log"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func main() {
	// 配置七牛云的 AccessKey 和 SecretKey
	accessKey := "CQwDB_qxHOG4oRWM98WtnhINjsInqUDo4qrPdsLB"
	secretKey := "3rqpIliuvCUNnBXoqIpLi1kQp8dlvXF09890_bfn"
	bucket := "infinite-blog"
	domain := "http://sdifmsux4.hd-bkt.clouddn.com" // 替换为你自己的七牛云域名，例如 http://yourbucket.qiniudn.com

	// 初始化七牛云配置
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 根据你的存储区域选择Zone
		UseHTTPS:      true,                 // 如果你的域名支持HTTPS，设置为true
		UseCdnDomains: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	// 列举存储桶中的所有文件
	limit := 1000         // 单次列举数量限制
	prefix := ""          // 列举文件的前缀
	delimiter := ""       // 列举文件的分隔符
	marker := ""          // 标记位置，继续列举时使用
	var fileKeys []string // 保存所有文件的key

	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			log.Fatalf("Failed to list files: %v", err)
		}
		for _, entry := range entries {
			fileKeys = append(fileKeys, entry.Key)
		}
		if !hasNext {
			break
		}
		marker = nextMarker
	}

	// 生成文件链接
	for _, key := range fileKeys {
		link := storage.MakePublicURL(domain, key)
		fmt.Println(link)
	}
}
