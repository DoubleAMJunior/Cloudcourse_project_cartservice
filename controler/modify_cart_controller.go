package controler

import (
	"cart_service/models"
	"cart_service/servicespkg"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModifyCartController struct {
	Helper servicespkg.ExternalServices
}

func (controler *ModifyCartController) ModifyCart(c *gin.Context) {
	fmt.Println("Proccessing Modify cart")
	val, exist := c.Get("id")
	if exist {
		if userID, ok := val.(string); ok {
			fmt.Println("CartController + id:" + userID)
			var body models.ModifyCartJson
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
			if body.Op == "rem" {
				body.Amount *= -1
			}
			err = models.DataBase.MofidyCart(userID, item.Title, body.Amount, int(item.Count))
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{"message": err})
				fmt.Println(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Cart Modified"})
			return
		}
	}
	c.AbortWithStatusJSON(500, gin.H{"message": "Internal error"})
}
