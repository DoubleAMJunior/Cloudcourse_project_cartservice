package controler

import (
	"cart_service/models"
	"cart_service/servicespkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCartController struct {
	Helper servicespkg.ExternalServices
}

func (controller *GetCartController) GetCart(c *gin.Context) {
	val, exist := c.Get("id")
	if exist {
		if userID, ok := val.(string); ok {
			cart, err := models.DataBase.GetCart(userID)
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{"message": "No cart"})
				return
			}
			c.JSON(http.StatusOK, cart)
			return
		}
	}
	c.AbortWithStatusJSON(500, gin.H{"message": "Internal error"})
}
