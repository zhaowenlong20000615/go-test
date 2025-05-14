package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Crypt(str string) string {
	// 创建MD5对象
	hash := md5.New()
	// 将字符串转换为字节数组并写入MD5对象
	hash.Write([]byte(str))
	// 计算MD5值
	bytes := hash.Sum(nil)
	// 将字节数组转换为十六进制字符串
	md5Str := hex.EncodeToString(bytes)
	return md5Str
}
