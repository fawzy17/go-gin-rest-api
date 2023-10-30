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
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
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
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode:    http.StatusInternalServerError,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Status:        "failed",
			Data:          err.Error(),
		})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, models.Response{
		StatusCode:    http.StatusCreated,
		StatusMessage: http.StatusText(http.StatusCreated),
		Status:        "success",
		Data:          product,
	})
}
func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Status:        "failed",
			Data:          err.Error(),
		})
		return
	}

	if config.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Status:        "failed",
			Data:          "Update data gagal",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Status:        "success",
		Data:          &product,
	})
}
func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if config.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
			StatusCode:    http.StatusNotFound,
			StatusMessage: http.StatusText(http.StatusNotFound),
			Status:        "failed",
			Data: map[string]interface{}{
				"message": "id tidak ditemukan",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Status:        "success",
		Data:          product,
	})
}
