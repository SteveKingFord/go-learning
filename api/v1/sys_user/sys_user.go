package sys_user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Tags System
// @Summary 用户登录
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "login"})
}


// @Tags System
// @Summary 用户注册
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/register [post]
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "register"})
}

// @Tags System
// @Summary 获取用户列表
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user [get]
func GetList(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"msg": "get list"})
}


// @Tags System
// @Summary 获取用户列表
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/:id [get]
func Get(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"msg": "get one data"})
}

// @Tags System
// @Summary  修改用户信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/:id [put]
func Put(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"msg": "put one data"})
}


// @Tags System
// @Summary  删除用户信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/:id [delete]
func Delete(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"msg": "delete one data"})
}