package main

import (
	"fmt"
	"github.com/chenluzhong150394/gotools/pkg/wechat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/simple/server/get", GetHandler)
	router.POST("wechat/call_back", WxHandler)
	//监听端口
	http.ListenAndServe(":8005", router)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func WxHandler(c *gin.Context) {
	var wx wechat.WxTextStruct
	_ = c.ShouldBindXML(&wx)
	wechat.WxFilterData(&wx)
	fmt.Println(wx.CreateTime)
	return
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
