package api

import (
	"github.com/gin-gonic/gin"
	"kingford-backend/modules/collection/service/category"
	"net/http"
)

// @Description 获取分页列表
// @Summary 分页列表
// @Tags Collection-Category
// @Accept application/json
// @Param pageIndex query int false "当前页"
// @Param pageSize query int false "分页条数"
// @success 200 {object} global.Response
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection-category [get]
func GetItemList(c *gin.Context) {
	var listService category.ListService
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.GetList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}


// @Description  通过id获取信息
// @Summary  通过Id获取信息
// @Tags Collection-Category
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} global.Response "成功"
// @Failure 400 object  global.Response "失败"
// @Router /api/collection-category/{id} [get]
func GetItem(c *gin.Context) {
	deleteService := category.GetService{}
	res := deleteService.Get(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

//@Description 创建信息
// @Tags Collection-Category
// @Summary 创建信息
// @Produce  application/json
// @Param collection body category.CreateService true "创建信息"
// @success 200 {object} global.Response "成功"
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection-category [post]
func CreateItem(c *gin.Context) {
	var createService category.CreateService
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description 通过id修改信息
// @Summary 通过id修改信息
// @Tags Collection-Category
// @Accept application/json
// @Param id path string true "主键id"
// @Param collection body category.UpdateService true "修改信息"
// @success 200 {object} global.Response "成功"
// @Failure 400 {object}  global.Response "失败"
// @Router /api/collection-category/{id} [put]
func UpdateItem(c *gin.Context) {
	var updateService category.UpdateService
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
// @Tags Collection-Category
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} global.Response "成功"
// @Failure 400 object  global.Response "失败"
// @Router /api/collection-category/{id} [delete]
func DeleteItem(c *gin.Context) {
	deleteService := category.DeleteService{}
	res := deleteService.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
