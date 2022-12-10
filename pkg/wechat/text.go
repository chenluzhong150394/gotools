package wechat

import (
	"encoding/xml"
)

// WxTextStruct 文本消息内容
type WxTextStruct struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content string `xml:"Content" json:"Content"`
	MsgID   string `xml:"MsgID" json:"MsgId,omitempty"`
}
