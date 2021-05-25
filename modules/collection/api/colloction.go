package api

import (
	"github.com/gin-gonic/gin"
	"kingford-backend/modules/collection/service/category"
	"net/http"

	service "kingford-backend/modules/collection/service/collection"
)

// @Description 获取分页列表
// @Summary 分页列表
// @Tags Collection
// @Accept application/json
// @Param pageIndex query int false "当前页"
// @Param pageSize query int false "分页条数"
// @success 200 {object} global.Response
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection [get]
func GetList(c *gin.Context) {
	var listService service.ListService
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.GetList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description  通过id获取信息
// @Summary  通过Id获取信息
// @Tags Collection
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} global.Response "成功"
// @Failure 400 object  global.Response "失败"
// @Router /api/collection/{id} [get]
func Get(c *gin.Context) {
	deleteService := service.GetService{}
	res := deleteService.Get(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

//@Description 创建信息
// @Tags Collection
// @Summary 创建信息
// @Produce  application/json
// @Param collection body service.CreateService true "创建信息"
// @success 200 {object} global.Response "成功"
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection [post]
func Create(c *gin.Context) {
	var createService service.CreateService
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description 通过id修改信息
// @Summary 通过id修改信息
// @Tags Collection
// @Accept application/json
// @Param id path string true "主键id"
// @Param collection body service.UpdateService true "修改信息"
// @success 200 {object} global.Response "成功"
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection/{id} [put]
func Update(c *gin.Context) {
	var updateService category.UpdateItemService
	if err := c.ShouldBind(&updateService); err == nil {
		id, _ := c.Params.Get("id")
		res := updateService.Update(id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description  通过id删除信息
// @Summary  通过id删除信息
// @Tags Collection
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} global.Response "成功"
// @Failure 400 object  global.Response "失败"
// @Router /api/collection/{id} [delete]
func Delete(c *gin.Context) {
	deleteService := service.DeleteService{}
	res := deleteService.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
