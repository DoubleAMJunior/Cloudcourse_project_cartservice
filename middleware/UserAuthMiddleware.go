package middleWare

import (
	"cart_service/servicespkg"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserAuthMiddleware struct {
	Helper servicespkg.ExternalServices
}

func (userAuth *UserAuthMiddleware) Action(c *gin.Context) {
	fmt.Println(c.Request.Method + c.Request.URL.Path)
	token := c.Request.Header["Authorization"][0]
	jwtToken := strings.ReplaceAll(token, "Bearer", "")
	jwtToken = strings.ReplaceAll(jwtToken, " ", "")
	jwtToken = strings.ReplaceAll(jwtToken, "\t", "")
	fmt.Println(jwtToken)
	access, err := userAuth.Helper.HasAccess(c.Request.Method, c.Request.URL.Path, jwtToken)
	if err != nil || !access {
		fmt.Println(err)
		fmt.Println(access)
		c.AbortWithStatusJSON(401, gin.H{"message": "Authorization issue"})
		return
	}
	user, err := userAuth.Helper.GetUser(jwtToken)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"message": "Problem with jwt tken"})
		return
	}
	c.Set("id", user.Id)
}
