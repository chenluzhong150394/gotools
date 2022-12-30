package wechat

import "github.com/chenluzhong150394/gotools/pkg/wechat"

// 聊天机器人

func WxText(w *wechat.WxTextStruct) {

	w.FromUserName, w.ToUserName = w.ToUserName, w.FromUserName
	w.Content = "你好呀"

	return
}
