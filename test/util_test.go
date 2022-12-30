package test

import (
	"fmt"
	"github.com/chenluzhong150394/gotools/pkg/util"
	"github.com/chenluzhong150394/gotools/pkg/wechat"
	"net/url"
	"testing"
)

func TestXml(t *testing.T) {

	type Iterm struct {
		LogisticsCode string `xml:"logisticsCode"`
		LogisticsName string `xml:"logisticsName"`
	}

	iterm := Iterm{
		LogisticsCode: "123",
		LogisticsName: "123",
	}

	b := make([]interface{}, 0)
	b = append(b, iterm)

	output, err := util.ParseXml(b)
	if err != nil {
		return
	}
	fmt.Println(string(output))

}

func TestRegx(t *testing.T) {

	result, num := util.MathPattern(`(?U)\b.+\b`, "Hello 世界！123 Go.")

	fmt.Println(result, num)
}

func TestGetIP(t *testing.T) {

	ip := util.GetHost()

	fmt.Println(ip)

}

func TestGetNowTimeUnix(t *testing.T) {

	fmt.Println(util.GetNowTimeUnix("ms"))

}

func TestGenXml(t *testing.T) {

	w := wechat.WxTextStruct{
		Content: "你好asd",
		MsgID:   "asd",
	}
	w.Base.ToUserName = "123"
	xml, err := util.ParseXml(w)
	if err != nil {
		return
	}

	fmt.Println("test .... ", string(xml))

}

func TestParseXml(t *testing.T) {
	w := wechat.WxTextStruct{
		Content: "你好",
		MsgID:   "000",
	}

	w.Base.CreateTime = "CreateTime"
	w.Base.ToUserName = "ToUserName"
	w.Base.MsgType = "MsgType"
	w.Base.FromUserName = "FromUserName"

	wechat.WxCoverData(&w)

	xml, err := util.ParseXml(w)
	if err != nil {
		return
	}
	fmt.Println(string(xml))
}

func TestTuWen(t *testing.T) {

	tuwen := wechat.TuWen{
		Base: wechat.Base{
			ToUserName:   "ToUserName",
			FromUserName: "FromUserName",
			MsgType:      "text",
		},
		ArticleCount: "1",
		Articles: wechat.Articles{
			Item: wechat.Item{
				Title:       "123",
				Description: "123",
				PicURL:      "123",
				URL:         "123",
			},
		},
	}

	wechat.WxCoverData(&tuwen)

	xml, err := util.ParseXml(tuwen)
	if err != nil {
		return
	}

	fmt.Println(string(xml))
}

func TestTuWen2(t *testing.T) {

	tuwen := wechat.TuWen{
		Base: wechat.Base{
			ToUserName:   "ToUserName",
			FromUserName: "FromUserName",
			MsgType:      "text",
		},
		ArticleCount: "1",
		Articles: wechat.Articles{
			Item: wechat.Item{
				Title:       "123",
				Description: "123",
				PicURL:      "123",
				URL:         "123",
			},
		},
	}

	xmlStr := wechat.TransferWxXmlString(&tuwen)

	fmt.Println(xmlStr)

}

func TestUrlParse(t *testing.T) {

	r, err := url.Parse("http://www.baidu.com/search?a=1")
	if err != nil {
		fmt.Println(r.Path)
	}

}
