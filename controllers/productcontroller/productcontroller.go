package productcontroller

import (
	"net/http"

	"github.com/fawzy17/rest-api-gin-gorm/config"
	"github.com/fawzy17/rest-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, models.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Status:        "success",
		Data:          products,
	})
}
func Show(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode:    http.StatusNotFound,
				StatusMessage: http.StatusText(http.StatusNotFound),
				Status:        "failed",
				Data:          "Not Found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode:    http.StatusInternalServerError,
				StatusMessage: http.StatusText(http.StatusInternalServerError),
				Status:        "failed",
				Data:          "Internal Server Error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Status:        "success",
		Data:          product,
	})
}

func Create(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
