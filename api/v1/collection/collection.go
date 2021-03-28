package collection

import (
	"github.com/gin-gonic/gin"
	"net/http"

	service "kingford-backend/service/collection"
)

// @Tags Collection
// @Summary 创建收藏
// @Produce  application/json
// @Param menu body service.CreateService true "创建收藏"
// @success 200 {object} dto.Response
// @Failure 400 {object}  dto.Response "失败"
// @Router /collection [post]
func Create(c *gin.Context) {
	var createService service.CreateService
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}