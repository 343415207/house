package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRouter(gin *gin.Engine) {
	gin.POST("/house/index", IndexForApp)
	gin.POST("/house/user/loginOrCreate", LoginOrCreate)
	gin.POST("/house/user/upPwd", UpPwd)
	gin.POST("/house/project/list", GetPList)
	gin.POST("/house/customer/create", SaveOrUpdateOrder)
	gin.POST("/house/customer/list", GetOrders)
}

//HandleResp 序列化
func HandleResp(c *gin.Context, data interface{}, err error) {
	logrus.Infof("http response: %+v, %+v, err: %v", c.Request.URL, data, err)
	res := &Resp{}
	res.Data = data
	if err != nil {
		res.Code = -1
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

//Resp HTTP返回值
type Resp struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

//JParse json序列化
func JParse(source interface{}) string {
	if target, err := json.Marshal(source); err == nil {
		return string(target)
	}
	return ""
}
