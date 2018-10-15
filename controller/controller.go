package controller

import (
	"house/controller/form"
	"house/service"

	"github.com/gin-gonic/gin"
)

//IndexForApp 首页信息
func IndexForApp(c *gin.Context) {

}

//LoginOrCreate 登录或者注册
func LoginOrCreate(c *gin.Context) {
	formLogin := &form.Login{}
	if err := c.Bind(formLogin); err != nil {
		HandleResp(c, nil, err)
		return
	}
	u, err := service.LoginOrCreateService(formLogin)
	HandleResp(c, u, err)
}

//UpPwd 更新密码
func UpPwd(c *gin.Context) {
	formLogin := &form.Login{}
	if err := c.Bind(formLogin); err != nil {
		HandleResp(c, nil, err)
		return
	}
	u, err := service.UpPwdService(formLogin)
	HandleResp(c, u, err)
}

//GetPList 获取项目列表
func GetPList(c *gin.Context) {
	fm := &form.PListForm{}
	if err := c.Bind(fm); err != nil {
		HandleResp(c, nil, err)
		return
	}
	u, err := service.GetPListService(fm)
	HandleResp(c, u, err)
}

//SaveOrUpdateOrder 报备客户
func SaveOrUpdateOrder(c *gin.Context) {

}

//GetOrders 查看报备订单
func GetOrders(c *gin.Context) {

}
