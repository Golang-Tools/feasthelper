package registryfileloader

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"

	"github.com/Golang-Tools/feasthelper/feast/core"
	"google.golang.org/protobuf/proto"
)

// Loads 转化字节流为`core.Registry`对象
// @params s []byte 待转化字节流
func Loads(s []byte) (*core.Registry, error) {
	registry := new(core.Registry)
	err := proto.Unmarshal(s, registry)
	if err != nil {
		return nil, err
	}
	return registry, nil
}

// Load 加载路径下的pb格式文件到`core.Registry`对象
// @params path string 待转化文件
func Load(path string) (*core.Registry, error) {
	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	res, err := Loads(fb)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MD5s 计算字节流的md5的16进制字符串
// @params s []byte 待计算字节流
func MD5s(s []byte) string {
	hasher := md5.New()
	hasher.Write(s)
	md5str := hex.EncodeToString(hasher.Sum(nil))
	return md5str
}

// MD5 计算文件的md5的16进制字符串
// @params path string 待计算文件
func MD5(path string) (string, error) {
	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	res := MD5s(fb)
	return res, nil
}
