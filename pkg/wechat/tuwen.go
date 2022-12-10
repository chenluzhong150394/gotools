package wechat

import "encoding/xml"

type Item struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	PicURL      string `json:"PicUrl"`
	URL         string `json:"Url"`
}
type Articles struct {
	Item Item `json:"item"`
}
type TuWen struct {
	XMLName xml.Name `xml:"xml"`
	Base
	ArticleCount string   `xml:"articleCount" json:"ArticleCount" default:"1"`
	Articles     Articles `xml:"articles" json:"Articles"`
}
