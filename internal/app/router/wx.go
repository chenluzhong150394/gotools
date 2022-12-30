package router

import (
	"github.com/chenluzhong150394/gotools/internal/app/controller"
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	r := gin.Default()

	wxV1 := r.Group("/v1")

	{
		wxV1.POST("/wechat/call_back", controller.WxCallBack) // 获取物料的版本号详情 支持条件过滤状态
	}
}
