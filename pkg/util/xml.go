package util

import (
	"encoding/xml"
)

//// ParseXml 将数据解析为xml
//func ParseXml(datas []interface{}) (output []byte, err error) {
//	body := RequestXmlBody{Items: Items{ItemData: datas}}
//	output, err = xml.MarshalIndent(body, "", "")
//	if err != nil {
//		return
//	}
//	return
//}

func ParseXml(body any) (output []byte, err error) {
	output, err = xml.MarshalIndent(body, "", "")
	if err != nil {
		return
	}
	return
}
