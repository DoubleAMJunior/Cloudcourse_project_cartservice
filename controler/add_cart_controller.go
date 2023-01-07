package controler

import (
	"cart_service/models"
	"cart_service/servicespkg"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddCartController struct {
	Helper servicespkg.ExternalServices
}

func (controler *AddCartController) AddCart(c *gin.Context) {
	fmt.Println("Proccessing cart add")
	val, exist := c.Get("id")
	if exist {
		if userID, ok := val.(string); ok {
			fmt.Println("CartController + id:" + userID)
			var body models.AddCartJson
			err := json.NewDecoder(c.Request.Body).Decode(&body)
			if err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(400, gin.H{"message": "bad data"})
				return
			}
			item, err := controler.Helper.GetProduct(int32(body.Id))
			if err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(500, gin.H{"message": "Internal error"})
			}
			if item.Count == 0 {
				c.AbortWithStatusJSON(400, gin.H{"message": "Item not available"})
				return
			}
			models.DataBase.AddItem(userID, item.Title)
			c.JSON(http.StatusOK, gin.H{"message": "Added to cart"})
			return
		}
	}
	c.AbortWithStatusJSON(500, gin.H{"message": "Internal error"})
}
