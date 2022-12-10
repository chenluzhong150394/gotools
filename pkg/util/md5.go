package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
)

// Md5Sign32 生成32位md5
func Md5Sign32(s []byte, options ...string) (signStr []byte) {
	m := md5.New()
	m.Write(s)

	if len(options) > 0 {
		switch options[0] {
		case "upper":
			signStr = []byte(strings.ToUpper(hex.EncodeToString(m.Sum(nil))))
		default:
			signStr = []byte(hex.EncodeToString(m.Sum(nil)))
		}
	} else {
		signStr = []byte(hex.EncodeToString(m.Sum(nil)))
	}
	return
}

// SignatureStruct 通过struct构造签名
func SignatureStruct(body interface{}) (sign []byte) {
	bodyStr, err := json.Marshal(body)
	//fmt.Println(bodyJsonStr)
	if err != nil {
		log.Println("json Marshal is error: ", err)
		return
	}
	sign = Md5Sign32(bodyStr)
	return
}
