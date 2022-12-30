package wechat

import (
	"fmt"
	"github.com/chenluzhong150394/gotools/pkg/util"
	"reflect"
	"strings"
)

// Base 微信公众号的消息都会有这4个字段。所以封装为一个结构体。
type Base struct {
	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	MsgType      string `xml:"MsgType" json:"MsgType"`
	CreateTime   string `xml:"CreateTime" json:"CreateTime"`
}

// WxCoverData 递归搜索 value 值 并处理为 cdata模式
func WxCoverData(body interface{}) {

	var bodyType reflect.Type
	var bodyValue reflect.Value

	if reflect.TypeOf(body).String() == "reflect.Value" {
		bodyType = reflect.TypeOf(body.(reflect.Value).Interface())
		bodyValue = body.(reflect.Value)
	} else {
		bodyType = reflect.TypeOf(body)
		bodyValue = reflect.ValueOf(body)
	}

	if bodyType.Kind() == reflect.Ptr {
		bodyType = bodyType.Elem()
	}

	if bodyValue.Kind() == reflect.Ptr {
		bodyValue = bodyValue.Elem()
	}

	for i := 0; i < bodyType.NumField(); i++ {

		field := bodyType.Field(i)

		value := bodyValue.Field(i)

		//fmt.Println(reflect.TypeOf(value).String())

		if value.IsValid() == false {
			fmt.Println("为空跳过")
			continue
		}

		switch value.Kind().String() {

		case "string":
			if ok := strings.HasPrefix(value.String(), "![CDATA"); ok {
				// 已经是 cdata 开头了，就过滤
				continue
			}

			value.SetString(fmt.Sprintf("![CDATA[%s]]", value))
		case "struct":
			if field.Name == "XMLName" {
				//fmt.Println("发现 XMLName 内置struct 跳过")
				continue
			}
			//fmt.Println("发现struct 进行递归调用")
			WxCoverData(value) // 递归调用
		default:
			//fmt.Println("即不是struct 也不是 string 跳过")
		}

		//fmt.Printf("%s %s: %v = %v\n", value.Kind(), field.Name, field.Type, value)
	}

}

func WxFilterData(body interface{}) {

	var bodyType reflect.Type
	var bodyValue reflect.Value

	if reflect.TypeOf(body).String() == "reflect.Value" {
		bodyType = reflect.TypeOf(body.(reflect.Value).Interface())
		bodyValue = body.(reflect.Value)
	} else {
		bodyType = reflect.TypeOf(body)
		bodyValue = reflect.ValueOf(body)
	}

	if bodyType.Kind() == reflect.Ptr {
		bodyType = bodyType.Elem()
	}

	if bodyValue.Kind() == reflect.Ptr {
		bodyValue = bodyValue.Elem()
	}

	for i := 0; i < bodyType.NumField(); i++ {

		field := bodyType.Field(i)

		value := bodyValue.Field(i)

		//fmt.Println(reflect.TypeOf(value).String())

		if value.IsValid() == false {
			fmt.Println("为空跳过")
			continue
		}

		switch value.Kind().String() {

		case "string":
			if ok := strings.HasPrefix(value.String(), "![CDATA"); ok {
				// 已经是 cdata 开头了，就过滤
				valueString := value.String()
				valueString = strings.Replace(valueString, "![CDATA[", "", -1)
				valueString = strings.Replace(valueString, "]]", "", -1)

				value.SetString(valueString)
			}

			//value.SetString(fmt.Sprintf("![CDATA[%s]]", value))
		case "struct":
			if field.Name == "XMLName" {
				//fmt.Println("发现 XMLName 内置struct 跳过")
				continue
			}
			//fmt.Println("发现struct 进行递归调用")
			WxCoverData(value) // 递归调用
		default:
			//fmt.Println("即不是struct 也不是 string 跳过")
		}

		//fmt.Printf("%s %s: %v = %v\n", value.Kind(), field.Name, field.Type, value)
	}

}

func TransferWxXmlString(body interface{}) (xmlStr string) {

	if reflect.ValueOf(body).Kind() != reflect.Ptr {
		panic("不是内存地址类型的数据")
	}

	WxCoverData(body)

	xmlByte, err := util.ParseXml(body)

	if err != nil {
		return
	}

	xmlStr = string(xmlByte)

	return
}
